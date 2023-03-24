package driver

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	_ "github.com/go-sql-driver/mysql" // defines mysql driver used
)

// DBMysqlOption options for mysql connection
type DBMysqlOption struct {
	IsEnable             bool
	Host                 string
	Port                 int
	Username             string
	Password             string
	DBName               string
	AdditionalParameters string
	MaxOpenConns         int
	MaxIdleConns         int
	ConnMaxLifetime      time.Duration
}

// NewMysqlDatabase return gorp dbmap object with MySQL options param
func NewMysqlDatabase(option DBMysqlOption) (*gorm.DB, error) {
	dbDsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&%s", option.Username, option.Password, option.Host, option.Port, option.DBName, option.AdditionalParameters)
	timeZone := "&loc=" + strings.Replace(os.Getenv("LOG_TIME_ZONE"), "/", "%2F", -1)
	dbDsn += timeZone

	db, err := sql.Open("mysql", dbDsn)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(option.ConnMaxLifetime)
	db.SetMaxIdleConns(option.MaxIdleConns)
	db.SetMaxOpenConns(option.MaxOpenConns)

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
		NowFunc: func() time.Time {
			ti, _ := time.LoadLocation("Asia/Jakarta")
			return time.Now().In(ti)
		},
	})
	if err != nil {
		return nil, err
	}

	return gormDB, nil
}
