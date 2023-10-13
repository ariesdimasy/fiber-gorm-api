package main

import (
	"log"

	"github.com/ariesdimasy/fiber-gorm-api/database"
	"github.com/ariesdimasy/fiber-gorm-api/routes"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Hello")
}

func setupRoutes(app *fiber.App) {

	app.Get("/api", welcome)
	app.Get("/api/users", routes.UserList)
	app.Get("/api/users/:id", routes.GetUserDetail)
	app.Post("/api/users", routes.CreateUser)
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)
	log.Fatal(app.Listen(":5700"))
}
