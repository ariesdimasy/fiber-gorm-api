package routes

import (
	"github.com/ariesdimasy/fiber-gorm-api/config"
	"github.com/ariesdimasy/fiber-gorm-api/models"
	"github.com/gofiber/fiber/v2"
)

// serializer
type User struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// create serializer based on model
func CreateResponseUser(userModel models.User) User {
	return User{
		ID:        userModel.ID,
		FirstName: userModel.FirstName,
		LastName:  userModel.LastName,
	}
}

func UserList(c *fiber.Ctx) error {
	var users []User

	query := config.DB

	query.Select("ID", "FirstName", "LastName").Find(&users)
	if query.Error != nil {
		return c.Status(500).JSON(query.Error)
	}

	// responseUsers := []User{}
	// for _, user := range users {
	// 	responseUser := CreateResponseUser(user)
	// 	responseUsers = append(responseUsers, responseUser)
	// }

	return c.Status(200).JSON(users)
}

func GetUserDetail(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	query := config.DB

	query.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return c.Status(400).JSON("User not found")
	}

	return c.Status(200).JSON(&user)

}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	query := config.DB

	query.Create(&user)
	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}
