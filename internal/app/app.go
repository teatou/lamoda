package app

import (
	"fmt"

	"github.com/teatou/lamoda/internal/config"
	"github.com/teatou/lamoda/pkg/logger"

	gojson "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type App struct {
	fiber  *fiber.App
	cfg    *config.Config
	logger logger.Logger
}

func NewApp(cfg *config.Config, logger logger.Logger) *App {
	return &App{
		fiber: fiber.New(fiber.Config{
			JSONEncoder:           gojson.Marshal,
			JSONDecoder:           gojson.Unmarshal,
			DisableStartupMessage: true,
		}),
		cfg:    cfg,
		logger: logger,
	}
}

func (a *App) Run() error {
	if err := a.MapHandlers(a.fiber, a.logger); err != nil {
		a.logger.Fatal("Cannot map handlers")
	}
	a.logger.Infof("Start server on port: %s", a.cfg.Server.Port)
	if err := a.fiber.Listen(fmt.Sprintf(":%d", a.cfg.Server.Port)); err != nil {
		a.logger.Fatalf("Error starting Server: ", err)
	}

	return nil
}
