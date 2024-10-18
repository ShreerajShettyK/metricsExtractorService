package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	SplunkURL    string
	SplunkToken  string
	SplunkIndex  string
	SplunkHost   string
	SplunkSource string
	SplunkSType  string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load(".env") // Load .env file

	return Config{
		SplunkURL:   getEnv("SPLUNKURL", ""),
		SplunkToken: getEnv("SPLUNKTOKEN", ""),
		SplunkIndex: getEnv("SPLUNKINDEX", ""),
		SplunkHost:  getEnv("SPLUNK_HOST", "localhost"),
		SplunkSType: getEnv("SPLUNK_SOURCETYPE", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// func getEnvAsInt(key string, fallback int64) int64 {
// 	if value, ok := os.LookupEnv(key); ok {
// 		i, err := strconv.ParseInt(value, 10, 64)
// 		if err != nil {
// 			return fallback
// 		}

// 		return i
// 	}

// 	return fallback
// }
