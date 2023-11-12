package app

import (
	"fmt"

	"github.com/teatou/lamoda/internal/config"
	"github.com/teatou/lamoda/pkg/logger"

	gojson "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type App struct {
	fiber     *fiber.App
	cfg       *config.Config
	apiLogger logger.Logger
}

func NewApp(cfg *config.Config, apiLogger logger.Logger) *App {
	return &App{
		fiber: fiber.New(fiber.Config{
			JSONEncoder:           gojson.Marshal,
			JSONDecoder:           gojson.Unmarshal,
			DisableStartupMessage: true,
		}),
		cfg:       cfg,
		apiLogger: apiLogger,
	}
}

func (a *App) Run() error {
	if err := a.MapHandlers(a.fiber, &a.apiLogger); err != nil {
		a.apiLogger.Fatal("Cannot map handlers")
	}
	a.apiLogger.Infof("Start server on port: %s", a.cfg.Server.Port)
	if err := a.fiber.Listen(fmt.Sprintf(":%d", a.cfg.Server.Port)); err != nil {
		a.apiLogger.Fatalf("Error starting Server: ", err)
	}

	return nil
}
