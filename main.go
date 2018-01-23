package main

import (
	"os"
	"time"

	"go.uber.org/zap"
)

func run() int {
	logger, _ := zap.NewProduction()
	for {
		err := notifyNewCoin()
		if err != nil {
			logger.Fatal("Fatal error", zap.Error(err))
			return 1
		}

		time.Sleep(5 * time.Minute)
	}
}

func main() {
	os.Exit(run())
}
