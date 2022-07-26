package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct{}

func (c *Config) Load() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func (c *Config) GetPort() string {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}
	return ":" + port
}
