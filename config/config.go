package config

import (
	"os"
)

type AppConfig struct {
	KafkaConfig KafkaConfig
	DBConfig    DBConfig
	Port        string
	Env         string
}

type KafkaConfig struct {
	Broker string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Name     string
	Password string
}

func LoadConfig() *AppConfig {
	return &AppConfig{
		DBConfig: DBConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Name:     getEnv("DB_NAME", "userdb"),
			Password: getEnv("DB_PASSWORD", "password"),
		},
		KafkaConfig: KafkaConfig{
			Broker: getEnv("KAFKA_BROKER", "localhost:9092"),
		},
		Port: getEnv("APP_PORT", "8080"),
		Env:  getEnv("APP_ENV", "local"),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
