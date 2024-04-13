package main

import (
	"log"

	"github.com/carlosveg/go_test/services"
	"github.com/gofiber/fiber/v2"
	_ "github.com/swaggo/fiber-swagger/example/docs" // docs is generated by Swag CLI, you have to import it.
)

// @title API
// @version 1.0
// @description Prueba GO
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	app := fiber.New()

	// swagger:route GET /admin/company/ admin listCompany
	// Get companies list
	//
	// security:
	// - apiKey: []
	// responses:
	//  401: CommonError
	//  200: GetCompanies
	app.Get("/", services.ListUsers)
	app.Post("/", services.AddUser)
	app.Post("/login", services.Login)

	err := app.Listen(":3001")

	if err != nil {
		log.Fatalf("fiber.Listen failed %s", err)
	}
}
