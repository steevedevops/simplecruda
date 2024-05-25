package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var cfg *config

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	User   string
	Pass   string
	Host   string
	Port   string
	DBname string
}

// funcoa de ciclo de vida de golang
// para definir valores default
// viper
// func init() {
// 	viper.SetDefault("API_PORT", "9000")
// 	viper.SetDefault("DB_HOST", "localhost")
// 	viper.SetDefault("DB_PORT", "5433")
// }

func Load() error {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalln("Error loading .env file")
	}
	cfg = &config{}

	cfg.API = APIConfig{
		Port: getEnv("API_PORT", "9000"),
	}

	cfg.DB = DBConfig{
		Host:   getEnv("DB_HOST", "localhost"),
		Port:   getEnv("DB_PORT", "5433"),
		User:   getEnv("DB_USER", ""),
		Pass:   getEnv("DB_PASS", ""),
		DBname: getEnv("DB_NAME", ""),
	}
	return nil
}

// initializados
func GetDBConfig() DBConfig {
	return cfg.DB
}

func GetServerPort() string {
	return cfg.API.Port
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
