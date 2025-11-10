package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string
}

func Load() *Config {
	// Загружаем .env файл
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  Файл .env не найден, используем переменные окружения")
	}

	return &Config{
		ServerPort: getEnv("SERVER_PORT", "8080"), // 8080 — значение по умолчанию
	}
}

// Вспомогательная функция: получить переменную или вернуть значение по умолчанию
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists { // exists - true если value  по ключу есть иначе false -> return defaultValue
		return value
	}
	return defaultValue
}
