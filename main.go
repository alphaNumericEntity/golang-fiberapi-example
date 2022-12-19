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
	// Product endpoints
	app.Post("/api/products", routes.CreateProduct)
	app.Get("/api/products", routes.GetProducts)
	app.Get("/api/products/:id", routes.GetProduct)
	app.Put("/api/products/:id", routes.UpdateProduct)
	// Order endpoints
	app.Post("/api/orders", routes.CreateOrder)
	app.Get("/api/orders", routes.GetOrders)
	app.Get("/api/orders/:id", routes.GetOrder)

}

func main() {

	database.Connect()
	app := fiber.New()
	createRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
