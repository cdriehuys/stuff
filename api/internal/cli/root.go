package cli

import (
	"context"
	"errors"
	"io"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/cdriehuys/stuff/api/internal/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewRootCmd(logStream io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stuff-api",
		Short: "Run the stuff API",
		RunE:  apiRunner(logStream),
	}

	cmd.PersistentFlags().Bool("debug", false, "Enable debug logging")
	viper.BindPFlag("debug", cmd.PersistentFlags().Lookup("debug"))

	cmd.Flags().String("addr", ":8080", "Address to listen for requests on")
	viper.BindPFlag("addr", cmd.Flags().Lookup("addr"))

	return cmd
}

func apiRunner(logStream io.Writer) func(*cobra.Command, []string) error {
	return func(cli *cobra.Command, _ []string) error {
		logger := createLogger(logStream)

		apiServer := api.NewStrictHandler(&api.Server{}, nil)
		r := http.NewServeMux()
		h := api.HandlerFromMux(apiServer, r)

		s := http.Server{
			Addr:         viper.GetString("addr"),
			Handler:      h,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  120 * time.Second,
		}

		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt)

		go func() {
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

		logger.Info("Shutdown complete.")

		return nil
	}
}
