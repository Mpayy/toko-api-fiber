package config

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewGorm(config *viper.Viper, log *logrus.Logger) *gorm.DB {
	host := config.GetString("DATABASE_HOST")
	port := config.GetInt("DATABASE_PORT")
	username := config.GetString("DATABASE_USERNAME")
	password := config.GetString("DATABASE_PASSWORD")
	database := config.GetString("DATABASE_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		database,
	)

	gormLogger := logger.New(
		&logrusWriter{Logger: log},
		logger.Config{
			SlowThreshold:             time.Second,
			Colorful:                  false,
			LogLevel:                  logger.LogLevel(config.GetInt("LOG_LEVEL")),
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})

	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	connection, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database connection", err)
	}

	connection.SetMaxOpenConns(100)
	connection.SetMaxIdleConns(10)
	connection.SetConnMaxLifetime(time.Second * time.Duration(300))

	return db
}

type logrusWriter struct {
	Logger *logrus.Logger
}

func (l *logrusWriter) Printf(message string, args ...interface{}) {
	l.Logger.Tracef(message, args...)
}
