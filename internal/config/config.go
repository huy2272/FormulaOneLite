package config

import (
	"log"
	"os"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Configuration struct {
	Uri string `env:"MONGODB_URI"`
}

func NewConfig() (*Configuration, error) {
	err := godotenv.Load()
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file")
		return nil, err
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Printf("You must set your 'MONGODB_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
		return nil, err
	}

	cfg := Configuration{}
	err = env.Parse(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
