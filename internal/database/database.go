package database

import (
	"fmt"
	"log"
	"os"

	"github.com/1shoukr/swiftplay-backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	conn *gorm.DB
}

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func LoadConfig() *Config {
	return &Config{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", ""),
		DBName:   getEnv("DB_NAME", "swiftplay_db"),
		SSLMode:  getEnv("DB_SSLMODE", "disable"),
	}
}

// NewDatabase creates a new GORM database connection
func NewDatabase(config *Config) (*Database, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
	)

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Get underlying sql.DB to test connection
	sqlDB, err := conn.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Connected to PostgreSQL database successfully with GORM")

	// Auto-migrate the schema
	if err := conn.AutoMigrate(&models.User{}, &models.Profile{}, &models.Match{}, &models.Message{}); err != nil {
		return nil, fmt.Errorf("failed to auto-migrate schema: %w", err)
	}

	log.Println("Database schema auto-migrated successfully")

	return &Database{conn: conn}, nil
}

// Close closes the database connection
func (db *Database) Close() error {
	if db.conn != nil {
		sqlDB, err := db.conn.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}
	return nil
}

// GetDB returns the underlying gorm.DB connection
func (db *Database) GetDB() *gorm.DB {
	return db.conn
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
