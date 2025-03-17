package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	RedisPort   string
	RedisHost   string
	ServerPort  string
	ServerHost  string
	MaterialAPI string
	SupplierAPI string
}

var Config *AppConfig

func init() {
	_ = godotenv.Load(".env")
	// Initialize Configuration
	Config = &AppConfig{
		RedisHost:   getEnv("REDIS_HOST", "localhost"),
		RedisPort:   getEnv("REDIS_PORT", "6379"),
		ServerPort:  getEnv("SERVER_PORT_USER", "8081"),
		ServerHost:  getEnv("SERVER_HOST_USER", "0.0.0.0"),
		MaterialAPI: getEnv("MATERIAL_API", "https://jgpjqcuk9e.execute-api.us-east-2.amazonaws.com/materials"),
		SupplierAPI: getEnv("SUPPLIER_API", "https://jgpjqcuk9e.execute-api.us-east-2.amazonaws.com/suppliers"),
	}
}

// to explicity define values for variables if value not set in .env

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		if value == "" {
			log.Printf("Environment variable %s is empty; using default value: %s", key, defaultValue)
			return defaultValue
		}
		log.Printf("Found environment variable %s with value: %s", key, value)
		return value
	}
	log.Printf("Environment variable %s not found; using default value: %s", key, defaultValue)
	return defaultValue
}
