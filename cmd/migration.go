package cmd

import (
	"eldho/eventori/config"
	"eldho/eventori/internal/app/appcontext"
	"eldho/eventori/internal/app/commons"
	"eldho/eventori/internal/app/commons/applicationConstants"
	"eldho/eventori/internal/app/migration"
	"errors"
	"flag"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var migrateUpCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate Up DB Loan API",
	Long:  `Please you know what are you doing by using this command`,
	Run: func(cmd *cobra.Command, args []string) {
		envName, _ := cmd.Flags().GetString("env")
		loadEnv(envName)

		var err error
		cfg := config.Config()
		logger := initLogger(cfg)
		app := appcontext.NewAppContext(cfg)
		opt := commons.InitCommonOptions(
			commons.WithDB(app),
		)
		if opt.Errors != nil {
			logger.Fatal(err.Error(),
				zap.String("context", "Init Common Options"),
			)
			return
		}

		runMigration(opt, applicationConstants.MigrateUp)
	},
}

var migrateDownCmd = &cobra.Command{
	Use:   "migratedown",
	Short: "Migrate Up DB Loan API",
	Long:  `Please you know what are you doing by using this command`,
	Run: func(cmd *cobra.Command, args []string) {
		envName, _ := cmd.Flags().GetString("env")
		loadEnv(envName)

		var err error
		cfg := config.Config()
		logger := initLogger(cfg)
		app := appcontext.NewAppContext(cfg)
		opt := commons.InitCommonOptions(
			commons.WithDB(app),
		)
		if opt.Errors != nil {
			logger.Fatal(err.Error(),
				zap.String("context", "Init Common Options"),
			)
			return
		}

		runMigration(opt, applicationConstants.MigrateDown)
	},
}

var seederUpCmd = &cobra.Command{
	Use:   "seeder",
	Short: "Seeder Up DB Loan API",
	Long:  `Please you know what are you doing by using this command`,
	Run: func(cmd *cobra.Command, args []string) {
		envName, _ := cmd.Flags().GetString("env")
		loadEnv(envName)

		var err error
		cfg := config.Config()
		logger := initLogger(cfg)
		app := appcontext.NewAppContext(cfg)
		opt := commons.InitCommonOptions(
			commons.WithDB(app),
		)
		if opt.Errors != nil {
			logger.Fatal(err.Error(),
				zap.String("context", "Init Common Options"),
			)
			return
		}

		runSeeder(opt, applicationConstants.MigrateUp)
	},
}

func init() {
	rootCmd.AddCommand(migrateUpCmd)
	rootCmd.AddCommand(migrateDownCmd)
	rootCmd.AddCommand(seederUpCmd)

	migrateUpCmd.PersistentFlags().StringP("env", "e", "prod", "environment type")
	migrateDownCmd.PersistentFlags().StringP("env", "e", "prod", "environment type")
	seederUpCmd.PersistentFlags().StringP("env", "e", "prod", "environment type")
}

func runMigration(opt *commons.Options, direction int) {
	pathMigration := os.Getenv("APP_MIGRATION_PATH")
	migrationDir := flag.String("migration-dir", pathMigration, "migration directory")
	log.Info().Msg("path migration : " + pathMigration)

	migrationConf, errMigrationConf := migration.NewMigrationConfig(*migrationDir,
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE_NAME"),
		"mysql")
	if errMigrationConf != nil {
		log.Error().Msg(errMigrationConf.Error())
		return
	}

	var errMigration error
	switch direction {
	case applicationConstants.MigrateUp:
		errMigration = migration.MigrateUp(migrationConf)
		break
	case applicationConstants.MigrateDown:
		errMigration = migration.MigrateDown(migrationConf)
		break
	default:
		errMigration = errors.New("Unknown migration direction")
	}
	if errMigration != nil {
		if errMigration.Error() != "no change" {
			log.Error().Msg(errMigration.Error())
			return
		}
		log.Info().Msg("Migration success : no change table . . .")
	}
}

func runSeeder(opt *commons.Options, direction int) {
	pathMigration := os.Getenv("APP_SEEDER_PATH")
	migrationDir := flag.String("migration-dir", pathMigration, "migration directory")
	log.Info().Msg("path migration : " + pathMigration)

	migrationConf, errMigrationConf := migration.NewMigrationConfig(*migrationDir,
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE_NAME"),
		"mysql")
	if errMigrationConf != nil {
		log.Error().Msg(errMigrationConf.Error())
		return
	}

	var errMigration error
	switch direction {
	case applicationConstants.MigrateUp:
		errMigration = migration.MigrateUp(migrationConf)
		break
	case applicationConstants.MigrateDown:
		errMigration = migration.MigrateDown(migrationConf)
		break
	default:
		errMigration = errors.New("Unknown migration direction")
	}
	if errMigration != nil {
		if errMigration.Error() != "no change" {
			log.Error().Msg(errMigration.Error())
			return
		}
		log.Info().Msg("Migration success : no change table . . .")
	}
}
