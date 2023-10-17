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

type ProductRequest struct {
	Name         string `json:"name" form:"name"`
	SerialNumber string `json:"serial_number" form:"serial_number"`
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

	var product Product
	productId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
			"data":    "",
		})
	}

	query := database.Database.Db

	query.First(&product, " id = ? ", productId)
	if query.Error != nil {
		return c.Status(500).JSON(query.Error)
	}

	if product.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "product not found",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "product fetch detail Success",
		"data":    product,
	})
}

func ProductUpdate(c *fiber.Ctx) error {

	productId, err := c.ParamsInt("id")

	productRequest := new(ProductRequest)
	errReq := c.BodyParser(productRequest)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
			"data":    "",
		})
	}

	if errReq != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": errReq.Error(),
			"data":    "",
		})
	}

	query := database.Database.Db

	query.Where("id = ?", productId).Updates(&Product{
		Name:         productRequest.Name,
		SerialNumber: productRequest.SerialNumber,
	})

	if query.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": query.Error,
			"data":    "",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "product updated success",
		"data":    "",
	})
}
