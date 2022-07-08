Example: load config from a local file
```go
package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/leaderseek/service/pkg/config"
	"github.com/leaderseek/service/pkg/worker"
)

func main() {
	if err := godotenv.Load("./env/worker.env"); err != nil {
		log.Fatalf("failed to load config to env, error = %v", err)
	}

	cfg, err := config.NewWorkerConfigFromEnv()
	if err != nil {
		log.Fatalf("failed to create config from env, error = %v", err)
	}

	w, err := worker.NewWorker(cfg)
	if err != nil {
		log.Fatalf("failed to create worker from config, error = %v", err)
	}

	w.Run()
}

```