package routes

import (
	"errors"
	"fmt"
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

func createUserSerializer(userModel models.User) User {
	return User{ID: userModel.ID, FirstName: userModel.FirstName, LastName: userModel.LastName}
}

func CreateUser(ctx *fiber.Ctx) error {
	var userModel models.User
	if err := ctx.BodyParser(&userModel); err != nil {
		log.Fatal("Could not parse request")
		return ctx.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&userModel)
	return ctx.Status(200).JSON(createUserSerializer(userModel))

}

func GetAllUsers(ctx *fiber.Ctx) error {
	var userModels = []models.User{}

	database.Database.Db.Find(&userModels)

	var users = []User{}

	for _, userModel := range userModels {
		var user = createUserSerializer(userModel)
		users = append(users, user)
	}

	return ctx.Status(200).JSON(users)
}

func findUserById(id int, userModel *models.User) error {
	database.Database.Db.Find(&userModel, "id = ?", id)
	if userModel.ID == 0 {
		return errors.New(fmt.Sprintf("No user with id %v exists", id))
	}
	return nil
}

func GetUserById(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")

	if err != nil {
		return ctx.Status(400).JSON("ensure id is correct")
	}

	var userModel models.User
	if err := findUserById(id, &userModel); err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	return ctx.Status(200).JSON(createUserSerializer(userModel))
}

func UpdateUserById(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")

	if err != nil {
		return ctx.Status(400).JSON("ensure id is correct")
	}

	var userModel models.User
	if err := findUserById(id, &userModel); err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	type UserUpdate struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	var userUpdate UserUpdate
	if err := ctx.BodyParser(&userUpdate); err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	userModel.FirstName = userUpdate.FirstName
	userModel.LastName = userUpdate.LastName

	database.Database.Db.Save(&userModel)
	return ctx.Status(200).JSON(createUserSerializer(userModel))
}

func DeleteUserById(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")

	if err != nil {
		return ctx.Status(400).JSON("ensure id is correct")
	}

	var userModel models.User
	if err := findUserById(id, &userModel); err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&userModel).Error; err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	return ctx.Status(200).JSON("user deleted")
}
