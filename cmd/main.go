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
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	defer logger.Sync()

	//logger.Info("Application started")
	//logger.Warn("Memory usage is high")
	//logger.Error("Database connection failed")

	config := configs.LoadConfig()
	scylla, _ := database.NewScylla(config)
	fmt.Println(scylla.AllKeyspaceMetadata())
	app := fiber.New()

	app.Get("/health", healthcheck)

	go func() {
		if err := app.Listen(":" + config.HostPort); err != nil {
			log.Fatal(err)
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c

	logger.Info("Shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := app.ShutdownWithContext(ctx); err != nil {
		logger.Fatal(err.Error())
	}

	scylla.Close()
	//redis.Close()

	logger.Info("Application stopped")
}
func healthcheck(ctx fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "ok",
	})
}
