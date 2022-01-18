package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func RegisterEmployeeRoutes(app *fiber.App) {
	app.Get("/employees", GetAllEmployees)
	app.Get("/employees/:id", GetEmployee)
}

func GetAllEmployees(c *fiber.Ctx) error {

	return c.SendString("some value")
}

func GetEmployee(c *fiber.Ctx) error {

	return c.SendString(fmt.Sprintf("requested employee with id: %v ", c.Params("id")))
}
