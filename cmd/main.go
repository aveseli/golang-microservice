package main

import (
	"fmt"
	"os"

	"github.com/aveseli/golang-microservice/internal/configuration"
	"github.com/aveseli/golang-microservice/internal/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	err := configuration.Connect()
	if err != nil {
		fmt.Print("Could not connect to mongodb. Error: ", err)
		os.Exit(1)
	}

	app := fiber.New()

	routes.RegisterEmployeeRoutes(app)

	app.Listen(":3000")
}
