package routes

import (
	"log"

	"github.com/alphanumericentity/fiber-api/database"
	"github.com/alphanumericentity/fiber-api/models"
	"github.com/gofiber/fiber/v2"
)

// this is serializer, not model
type User struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateUserSerializer(userModel models.User) User {
	return User{ID: userModel.ID, FirstName: userModel.FirstName, LastName: userModel.LastName}
}

func CreateUser(ctx *fiber.Ctx) error {
	var userModel models.User //should be user model? todo
	if err := ctx.BodyParser(&userModel); err != nil {
		log.Fatal("Could not parse request")
		return ctx.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&userModel)
	return ctx.Status(200).JSON(CreateUserSerializer(userModel))

}

func GetAllUsers(ctx *fiber.Ctx) error {
	var userModels = []models.User{}

	database.Database.Db.Find(&userModels)

	var users = []User{}

	for _, userModel := range userModels {
		var user = CreateUserSerializer(userModel)
		users = append(users, user)
	}

	return ctx.Status(200).JSON(users)

}
