package main

import (
	"github.com/carlosveg/go_test/services"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", services.ListUsers)
	app.Post("/", services.AddUser)

	app.Listen(":3000")
}
