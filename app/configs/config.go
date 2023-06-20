package configs

import "os"

// Add your Env Keys here
type Config struct {
	GIN_MODE       string
	APP_PORT       string
	DB_HOST        string
	DB_USER        string
	DB_PASS        string
	DB_NAME        string
	DB_PORT        string
	JWT_SECRET_KEY string
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// Load all your env keys here
func LoadConfig() Config {
	return Config{
		GIN_MODE:       getEnv("APP_MODE", "development"),
		APP_PORT:       getEnv("APP_PORT", ""),
		DB_HOST:        getEnv("DB_HOST", ""),
		DB_USER:        getEnv("DB_USER", ""),
		DB_PASS:        getEnv("DB_PASS", ""),
		DB_NAME:        getEnv("DB_NAME", ""),
		DB_PORT:        getEnv("DB_PORT", ""),
		JWT_SECRET_KEY: getEnv("JWT_SECRET_KEY", ""),
	}
}
