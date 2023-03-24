package commons

import (
	"eldho/eventori/internal/app/appcontext"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Options common option for all object that needed
type Options struct {
	AppCtx    *appcontext.AppContext
	Db        *gorm.DB
	Dwh       *gorm.DB
	Logger    *zap.Logger
	Validator *validator.Validate
	Errors    error
}

func InitCommonOptions(options ...func(*Options)) *Options {
	logLevel := zerolog.InfoLevel
	logLevelP, err := zerolog.ParseLevel(os.Getenv("APP_LOG_LEVEL"))
	if err == nil {
		logLevel = logLevelP
	}
	zerolog.SetGlobalLevel(logLevel)

	opt := &Options{}
	for _, o := range options {
		o(opt)
		if opt.Errors != nil {
			return opt
		}
	}
	return opt
}

func WithDB(appCtx *appcontext.AppContext) func(*Options) {
	return func(opt *Options) {
		db, err := appCtx.GetDBInstance(appcontext.DBDialectMysql)
		if err != nil {
			opt.Errors = err
			return
		}
		opt.Db = db
	}
}

func WithLogger(logger *zap.Logger) func(*Options) {
	return func(opt *Options) {
		opt.Logger = logger
	}
}

func WithValidator(appCtx *appcontext.AppContext) func(*Options) {
	return func(opt *Options) {
		validator := validator.New()
		opt.Validator = validator
	}
}
