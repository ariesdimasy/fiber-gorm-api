package main

import (
	"log"

	"github.com/ariesdimasy/fiber-gorm-api/config"
	"github.com/ariesdimasy/fiber-gorm-api/database"
	"github.com/ariesdimasy/fiber-gorm-api/routes"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	type User struct {
		ID        uint   `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	var users []User

	db := config.ConnecDb()

	db.Find(&users)

	if db.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": db.Error,
			"data":    "",
		})
	} else {
		return c.Status(200).JSON(fiber.Map{
			"message": "fetch users data successfully",
			"data":    users,
		})
	}
}

func setupRoutes(app *fiber.App) {

	app.Get("/api", welcome)
	app.Get("/api/users", routes.UserList)
	app.Get("/api/users/:id", routes.GetUserDetail)
	app.Post("/api/users", routes.CreateUser)

	app.Get("/api/products", routes.ProductList)
	app.Get("/api/products/:id<int>", routes.ProductDetail)
	app.Post("/api/products", routes.ProductCreate)
	app.Put("/api/products/:id<int>", routes.ProductUpdate)
}

func main() {
	database.ConnectDb()

	app := fiber.New()

	setupRoutes(app)
	log.Fatal(app.Listen(":5700"))
}
