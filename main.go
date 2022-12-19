package main

import (
	"log"

	"github.com/alphanumericentity/fiber-api/database"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welecome")

}

func main() {

	database.Connect()
	app := fiber.New()

	app.Get("/api", welcome)

	log.Fatal(app.Listen(":3000"))
}
