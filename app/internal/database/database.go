package database

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/suraboy/go-miniproject/app/internal"
	"github.com/suraboy/go-miniproject/app/internal/models"
)

// DB holds the database connection
var DB *gorm.DB

// InitDatabase initializes the database connection
func InitDatabase(config *internal.Config) error {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Bangkok",
		config.Database.Host,
		config.Database.User,
		config.Database.Password,
		config.Database.DBName,
		config.Database.Port,
	)

	// Configure GORM logger
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	// Connect to database
	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	// Configure connection pool
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %w", err)
	}

	// Set connection pool settings
	sqlDB.SetMaxOpenConns(config.Database.MaxOpenConns)
	sqlDB.SetMaxIdleConns(config.Database.MaxIdleConns)

	if lifetime, err := time.ParseDuration(config.Database.ConnMaxLifetime); err == nil {
		sqlDB.SetConnMaxLifetime(lifetime)
	}

	// Test connection
	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	DB = db
	log.Println("✅ Database connected successfully")

	return nil
}

// AutoMigrate runs database migrations
func AutoMigrate() error {
	if DB == nil {
		return fmt.Errorf("database not initialized")
	}

	// Auto migrate models
	err := DB.AutoMigrate(
		&models.Loan{},
	)

	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	log.Println("✅ Database migration completed")
	return nil
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return DB
}
