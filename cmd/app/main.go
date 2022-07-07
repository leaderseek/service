package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/leaderseek/service/pkg/app"
	"github.com/leaderseek/service/pkg/config"
)

func main() {
	if err := godotenv.Load("./env/app.env"); err != nil {
		log.Fatalf("failed to load config to env, error = %v", err)
	}

	cfg, err := config.NewAppConfigFromEnv()
	if err != nil {
		log.Fatalf("failed to create config from env, error = %v", err)
	}

	app, err := app.NewApp(cfg)
	if err != nil {
		log.Fatalf("failed to create app from config, error = %v", err)
	}

	app.Run()
}
