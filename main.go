package main

import (
	"math/rand"
	"os"
	"time"

	_ "github.com/lib/pq"
	"go.uber.org/zap"

	"github.com/semka95/payment-service/cmd"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		os.Exit(1)
	}
	defer logger.Sync()
	zap.ReplaceGlobals(logger)

	rand.Seed(time.Now().UnixNano())

	config, err := cmd.NewConfig()
	if err != nil {
		logger.Error("can't decode config", zap.Error(err))
		return
	}

	srv := cmd.NewServer(logger, config)
	srv.RunServer()
}
