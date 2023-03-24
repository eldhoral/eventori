package cmd

import (
	"fmt"
	"os"
	"time"

	"eldho/eventori/config"
	"eldho/eventori/internal/app/appcontext"
	"eldho/eventori/internal/app/commons"
	"eldho/eventori/internal/app/repository"
	"eldho/eventori/internal/app/repository/cataloguesRepository"
	"eldho/eventori/internal/app/repository/schedulesRepository"
	"eldho/eventori/internal/app/server"
	"eldho/eventori/internal/app/service"
	"eldho/eventori/internal/app/service/modelService"

	"github.com/joho/godotenv"
	gologger "github.com/mo-taufiq/go-logger"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "nobu-loan-api",
	Short: "A brief description of your application",
	Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your application.`,
	Run: func(cmd *cobra.Command, args []string) {
		loadEnv("")
		start()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()
}

var levelMapper = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func initLogger(cfg config.Provider) *zap.Logger {
	var level zapcore.Level
	if lvl, ok := levelMapper[cfg.GetString("logger.level")]; ok {
		level = lvl
	} else {
		level = zapcore.InfoLevel
	}

	loggerCfg := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(level),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    "time",
			EncodeTime: zapcore.RFC3339NanoTimeEncoder,

			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	logger, _ := loggerCfg.Build()
	return logger
}

func loadEnv(envName string) {
	gologger.LogConf.NestedLocationLevel = 2
	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:     os.Stderr,
			NoColor: false,
		},
	)

	dotenvPath := "params/.env"

	if envName == "test" {
		dotenvPath = "params/.env.test"
	}

	err := godotenv.Load(dotenvPath)
	if err != nil {
		log.Error().Msg("Error loading .env file")
	}
}

func start() {
	var err error
	cfg := config.Config()
	logger := initLogger(cfg)

	if cast.ToBool(os.Getenv("FORCE_TO_UTC")) {
		loc, err := time.LoadLocation("UTC")
		if err != nil {
			logger.Fatal(err.Error(),
				zap.String("context", "Set timezone to UTC"),
			)
			return
		}
		time.Local = loc
	}

	app := appcontext.NewAppContext(cfg)

	opt := commons.InitCommonOptions(
		commons.WithDB(app),
		commons.WithLogger(logger),
		commons.WithValidator(app),
	)

	if opt.Errors != nil {
		logger.Fatal(err.Error(),
			zap.String("context", "Init Common Options"),
		)
		return
	}

	repo := wiringRepository(repository.Option{
		Options: *opt,
	})

	service := wiringService(service.Option{
		Options:      *opt,
		Repositories: repo,
	})

	server := server.NewServer(*opt, service)

	// run app
	server.StartApp()

}

func wiringRepository(repoOption repository.Option) *repository.Repositories {
	repo := repository.Repositories{
		CataloguesRepository: cataloguesRepository.NewCataloguesRepository(repoOption.Db),
		SchedulesRepository:  schedulesRepository.NewSchedulesRepository(repoOption.Db),
	}

	return &repo
}

func wiringService(serviceOption service.Option) *service.Services {
	// wiring up all services
	svc := service.Services{

		ModelService: modelService.NewModelServiceImpl(
			serviceOption.CataloguesRepository,
			serviceOption.SchedulesRepository,
		),
	}
	return &svc
}
