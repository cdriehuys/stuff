package cli

import (
	"fmt"
	"io"
	"io/fs"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/tern/v2/migrate"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func newMigrateCmd(logStream io.Writer, migrationFS fs.FS) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "Migrate the database forwards",
		RunE:  migrateRunner(logStream, migrationFS),
	}

	return cmd
}

func migrateRunner(logStream io.Writer, migrationFS fs.FS) func(*cobra.Command, []string) error {
	return func(cli *cobra.Command, s []string) error {
		logger := createLogger(logStream)

		dsn := viper.GetString("dsn")
		pool, err := pgxpool.New(cli.Context(), dsn)
		if err != nil {
			return fmt.Errorf("failed to connect to database: %v", err)
		}

		defer pool.Close()

		err = pool.AcquireFunc(cli.Context(), func(c *pgxpool.Conn) error {
			migrator, err := migrate.NewMigrator(cli.Context(), c.Conn(), "public.schema_version")
			if err != nil {
				return fmt.Errorf("failed to build migrator: %v", err)
			}

			if err := migrator.LoadMigrations(migrationFS); err != nil {
				return fmt.Errorf("failed to load migrations: %v", err)
			}

			migrator.OnStart = func(i int32, name string, direction string, sql string) {
				logger.Info("Executing migration", "name", name, "direction", direction)
				logger.Debug("Migration contents", "sql", sql)
			}

			if err := migrator.Migrate(cli.Context()); err != nil {
				return fmt.Errorf("failed to run migrations: %v", err)
			}

			logger.Info("Database migrations completed successfully")

			return nil
		})

		return err
	}
}
