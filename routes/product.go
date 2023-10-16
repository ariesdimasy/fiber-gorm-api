package routes

import (
	"fmt"

	"github.com/ariesdimasy/fiber-gorm-api/database"
	"github.com/gofiber/fiber/v2"
)

// serializer
type Product struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

func CreateResponseProduct(productModel Product) Product {
	return Product{
		ID:           productModel.ID,
		Name:         productModel.Name,
		SerialNumber: productModel.SerialNumber,
	}
}

func ProductList(c *fiber.Ctx) error {
	var products []Product

	query := database.Database.Db

	query.Select("id", "name", "serial_number").Find(&products)
	fmt.Println(query.Statement.Vars...)
	if query.Error != nil {
		return c.Status(500).JSON(query.Error)
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "fetch products data successfully",
		"data":    products,
	})

}

func ProductCreate(c *fiber.Ctx) error {
	var product Product

	query := database.Database.Db

	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	query.Create(&product)
	if query.Error != nil {
		return c.Status(500).JSON(query.Error)
	}
	responseProduct := CreateResponseProduct(product)

	return c.Status(200).JSON(fiber.Map{
		"message": "product Create Success",
		"data":    responseProduct,
	})
}

func ProductDetail(c *fiber.Ctx) error {

	productId := c.Params("id")

	fmt.Println(productId)

	return c.Status(200).JSON(fiber.Map{
		"message": "product Create Success",
		"data":    productId,
	})
}
