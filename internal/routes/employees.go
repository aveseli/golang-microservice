package routes

import (
	"fmt"

	"github.com/aveseli/golang-microservice/internal/repository"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterEmployeeRoutes(app *fiber.App) {
	app.Get("/employees", GetAllEmployees)
	app.Get("/employees/:id", GetEmployee)
	app.Post("/employees", PostEmployee)
}

func GetAllEmployees(c *fiber.Ctx) error {
	e, err := repository.GetAllEmployees()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.SendStatus(fiber.StatusNotFound)
		}
		fmt.Println(fmt.Errorf("GetAllEmployees: Error %v", err))
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(e)
}

func GetEmployee(c *fiber.Ctx) error {
	e, err := repository.GetEmployee(c.Params("id"))
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.SendStatus(fiber.StatusNotFound)
		}
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(e)
}

func PostEmployee(c *fiber.Ctx) error {
	e := new(repository.Employee)
	if err := c.BodyParser(e); err != nil {
		return err
	}

	repository.InsertEmployee(*e)

	return nil
}
