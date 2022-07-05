package main

import (
	"log"

	"github.com/leaderseek/service/pkg/worker"
)

func main() {
	cfg, err := worker.NewConfigFromEnv()
	if err != nil {
		log.Fatalf("failed to create config from env, error = %v", err)
	}

	w, err := worker.NewWorker(cfg)
	if err != nil {
		log.Fatalf("failed to create worker from config, error = %v", err)
	}

	w.Run()
}
