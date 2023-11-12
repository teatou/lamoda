package main

import (
	"os"

	"github.com/teatou/lamoda/internal/app"
	"github.com/teatou/lamoda/internal/config"
	"github.com/teatou/lamoda/pkg/logger"
)

const configEnv = "CONFIG"

func main() {
	val, ok := os.LookupEnv(configEnv)
	if !ok {
		// val = "configs/dev.yaml"
		panic("config env not found")
	}

	cfg, err := config.LoadConfig(val)
	if err != nil {
		panic("uploading config error")
	}

	logger, err := logger.NewZapLogger(cfg.Logger.Level)
	if err != nil {
		panic("init logger error")
	}
	defer logger.Sync()

	app := app.NewApp(cfg, logger)
	if err = app.Run(); err != nil {
		panic(err)
	}
}
