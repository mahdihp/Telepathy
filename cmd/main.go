package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/mahdihp/telepathy/configs"
	"github.com/mahdihp/telepathy/internal/database"
)

func main() {
	config := configs.LoadConfig()
	scylla, _ := database.NewScylla(config)
	fmt.Println(scylla.AllKeyspaceMetadata())
	app := fiber.New()

	app.Get("/health", healthcheck)
	log.Fatal(app.Listen(":" + config.HostPort))

}
func healthcheck(ctx fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "ok",
	})
}
