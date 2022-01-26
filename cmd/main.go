package main

import (
	"fmt"
	"os"

	"github.com/aveseli/golang-microservice/internal/cfg"
	"github.com/aveseli/golang-microservice/internal/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	err := cfg.Connect()
	if err != nil {
		fmt.Print("Could not connect to mongodb. Error: ", err)
		os.Exit(1)
	}
	defer func() {
		if err := cfg.Disconnect(); err != nil {
			panic(err)
		}
	}()

	app := fiber.New()

	routes.RegisterEmployeeRoutes(app)

	app.Listen(":3000")
}
