package database

import (
	"context"
	"fmt"
	"rest-api/bin/config"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgre(ctx context.Context) *gorm.DB {
	connection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		config.GlobalEnv.PostgreSQL.Host,
		config.GlobalEnv.PostgreSQL.User,
		config.GlobalEnv.PostgreSQL.Password,
		config.GlobalEnv.PostgreSQL.DBName,
		config.GlobalEnv.PostgreSQL.Port,
		config.GlobalEnv.PostgreSQL.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database postgre")
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("Failed to create pool connection database postgre")
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}
