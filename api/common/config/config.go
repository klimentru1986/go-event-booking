package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DbDriver string
	DbSource string
}

func New(filenames ...string) *Config {

	if err := godotenv.Load(filenames...); err != nil {
		log.Print("No .env file found")
	}

	return &Config{
		DbDriver: getEnv("DB_DRIVER", "sqlite3"),
		DbSource: getEnv("DB_SOURCE", "common/db/api.db"),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		log.Print(key, ": ", value)
		return value
	}

	return defaultVal
}
