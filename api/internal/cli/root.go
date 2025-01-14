package cli

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"os/signal"
	"reflect"
	"strings"
	"time"

	"github.com/cdriehuys/stuff/api/internal/api"
	"github.com/cdriehuys/stuff/api/internal/models"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/text/language"
)

func NewRootCmd(logStream io.Writer, localeFS fs.FS, migrationFS fs.FS) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stuff-api",
		Short: "Run the stuff API",
		RunE:  apiRunner(logStream, localeFS),
	}

	cmd.PersistentFlags().Bool("debug", false, "Enable debug logging")
	viper.BindPFlag("debug", cmd.PersistentFlags().Lookup("debug"))

	cmd.Flags().String("addr", ":8080", "Address to listen for requests on")
	viper.BindPFlag("addr", cmd.Flags().Lookup("addr"))

	cmd.Flags().String("dsn", "postgres://localhost", "DSN to connect to the application database")
	viper.BindEnv("dsn", "STUFF_DSN")
	viper.BindPFlag("dsn", cmd.Flags().Lookup("dsn"))

	cmd.AddCommand(newMigrateCmd(logStream, migrationFS))

	return cmd
}

func apiRunner(logStream io.Writer, localeFS fs.FS) func(*cobra.Command, []string) error {
	return func(cli *cobra.Command, _ []string) error {
		logger := createLogger(logStream)

		bundle := i18n.NewBundle(language.English)
		bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
		bundle.LoadMessageFileFS(localeFS, "locale.*.toml")

		pool, err := pgxpool.New(cli.Context(), viper.GetString("dsn"))
		if err != nil {
			return fmt.Errorf("failed to create database pool: %v", err)
		}

		defer pool.Close()

		validate := validator.New(validator.WithRequiredStructEnabled())
		validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			// skip if tag key says it should be ignored
			if name == "-" {
				return ""
			}
			return name
		})

		assetModel := models.NewAssetModel(logger, pool, validate)
		modelModel := models.NewModelModel(logger, pool, validate)
		vendorModel := models.NewVendorModel(logger, pool, validate)

		server := api.NewServer(logger, bundle, validate, assetModel, modelModel, vendorModel)

		strictHandler := api.NewStrictHandler(server, []api.StrictMiddlewareFunc{
			// Middleware are executed in order, so the panic recovery and error handling need to
			// happen last.
			server.LocalizationMiddleware(),
			server.PanicRecoveryMiddleware(),
			server.ErrorMiddleware(),
		})
		mux := http.NewServeMux()
		handler := api.HandlerFromMux(strictHandler, mux)

		s := http.Server{
			Addr:         viper.GetString("addr"),
			Handler:      handler,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  120 * time.Second,
		}

		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt)

		go func() {
			logger.Info("Starting API server.", "addr", s.Addr)

			if err := s.ListenAndServe(); err != nil {
				if !errors.Is(err, http.ErrServerClosed) {
					logger.Error("Unexpected server error.", "error", err)
				}
			}
		}()

		<-interrupt
		signal.Stop(interrupt)

		logger.Info("Received interrupt; Shutting down.")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		if err := s.Shutdown(shutdownCtx); err != nil {
			logger.Error("Server did not shut down gracefully.", "error", err)
		}

		pool.Close()
		logger.Info("Closed database pool.")

		logger.Info("Shutdown complete.")

		return nil
	}
}
