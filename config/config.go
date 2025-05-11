package config

import (
	"log"
	"sync"
)

type Config struct {
	Port            string
	AuthServiceURL  string // HTTP URL (если есть REST-интерфейс)
	UserServiceGRPC string // gRPC адрес
}

var (
	once     sync.Once
	instance *Config
)

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{
			Port:           ":8080",
			UserServiceGRPC: "localhost:50051", // gRPC-адрес сервиса пользователей
		}
		log.Println("Configuration loaded")
	})
	return instance
}