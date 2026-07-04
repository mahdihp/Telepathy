package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/mahdihp/telepathy/configs"
	"github.com/mahdihp/telepathy/internal/database"
	mainlogger "github.com/mahdihp/telepathy/internal/logger"
	"go.uber.org/zap"
)

var loggerZap *zap.Logger
var err error

func init() {
	loggerZap, err = mainlogger.NewDevelopment()
	if err != nil {
		panic(err)
	}

}
func main() {
	defer func(loggerZap *zap.Logger) {
		err := loggerZap.Sync()
		if err != nil {

		}
	}(loggerZap)

	config := configs.LoadConfig()

	scylla, err := database.NewScyllaDb(config)
	if err != nil {
		log.Fatal(err)
	}
	err = database.RunMigrations(scylla)
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New(fiber.Config{})
	app.Hooks().OnPreStartupMessage(func(sm *fiber.PreStartupMessageData) error {
		sm.BannerHeader = "FOOBER " + sm.Version + "\n-------"

		// Optional: you can also remove old entries
		// sm.ResetEntries()

		sm.AddInfo("prefork", "Prefork", fmt.Sprintf("%v", sm.Prefork), 15)
		return nil
	})

	app.Get("/health", healthcheck)

	go func() {
		if err := app.Listen(":" + config.HostPort); err != nil {
			log.Fatal(err)
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	loggerZap.Info("Shutting down...")

	scylla.Close()
	//redis.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := app.ShutdownWithContext(ctx); err != nil {
		loggerZap.Fatal(err.Error())
	}

	loggerZap.Info("Application stopped")
}
func healthcheck(ctx fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"Status":      "ok",
		"Server Time": time.Now().Format("2006-1-02 15:04:05"),
	})
}
