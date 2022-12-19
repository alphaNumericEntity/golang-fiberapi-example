package main

import (
	"log"

	"github.com/alphanumericentity/fiber-api/database"
	"github.com/alphanumericentity/fiber-api/routes"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welecome")
}

func createRoutes(app *fiber.App) {
	app.Get("/api/", welcome)
	app.Post("/api/user", routes.CreateUser)
	app.Get("/api/users", routes.GetAllUsers)
	app.Get("/api/user/:id", routes.GetUserById)
	app.Put("/api/user/:id", routes.UpdateUserById)
	app.Delete("/api/user/:id", routes.DeleteUserById)

}

func main() {

	database.Connect()
	app := fiber.New()
	createRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
