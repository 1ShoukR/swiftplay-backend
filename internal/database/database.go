package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

type Database struct {
	conn *pgx.Conn
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

// NewDatabase creates a new database connection
func NewDatabase(config *Config) (*Database, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
	)

	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := conn.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Connected to PostgreSQL database successfully")

	return &Database{conn: conn}, nil
}

//  closes the database connection
func (db *Database) Close() error {
	if db.conn != nil {
		return db.conn.Close(context.Background())
	}
	return nil
}

//  returns the underlying pgx connection
func (db *Database) GetConn() *pgx.Conn {
	return db.conn
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
