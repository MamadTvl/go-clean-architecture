package db

import (
	"clean-architecture/domain/model"
	"clean-architecture/infrastructure/config"
	"clean-architecture/infrastructure/logger"
	"context"
	"fmt"
	"time"

	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func NewDatabase(logger logger.Logger, config *config.Config) *Database {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		config.Database.Host,
		config.Database.User,
		config.Database.Password,
		config.Database.DbName,
		config.Database.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.GetGormLogger()})
	if err != nil {
		logger.Panic(err)
	}

	conn, err := db.DB()
	if err != nil {
		logger.Info("couldn't get db connection")
		logger.Panic(err)
	}

	conn.SetConnMaxLifetime(time.Minute * 5)
	conn.SetMaxOpenConns(5)
	conn.SetMaxIdleConns(1)

	return &Database{DB: db}
}

func Migrate(lc fx.Lifecycle, logger logger.Logger, db *Database) {
	lc.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				err := db.DB.Migrator().AutoMigrate(&model.User{})
				if err == nil {
					logger.Info("Connected To Database")
					return nil
				}
				logger.Error(err)
				return err
			},
		},
	)
}
