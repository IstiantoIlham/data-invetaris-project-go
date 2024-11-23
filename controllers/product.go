package controllers

import (
	. "data-invetaris/database"
	"data-invetaris/models"
	"data-invetaris/utils"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"os"
	"path/filepath"
)

var validate = validator.New()

type Product struct{}

func (p Product) Get(c *fiber.Ctx) error {
	var products []models.Product
	DB.Find(&products)
	return utils.ResponseSuccess(c, "Success get all products", products)
}

func (p Product) GetById(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product
	if err := DB.First(&product, id).Error; err != nil {
		return utils.ResponseError(c, fiber.StatusNotFound, "Product not found", nil)
	}
	return utils.ResponseSuccess(c, "Success get product", product)
}

func (p Product) Create(c *fiber.Ctx) error {
	product := new(models.Product)
	if err := c.BodyParser(product); err != nil {
		return utils.ResponseError(c, fiber.StatusBadRequest, "Invalid input", nil)
	}

	if errors, err := utils.ValidateInput(product); err != nil {
		return utils.ResponseError(c, fiber.StatusBadRequest, "Invalid input", errors)
	}

	DB.Create(product)
	return utils.ResponseSuccess(c, "Success create product", product)
}

func (p Product) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product
	if err := DB.First(&product, id).Error; err != nil {
		return utils.ResponseError(c, fiber.StatusNotFound, "Product not found", nil)
	}

	if err := c.BodyParser(&product); err != nil {
		return utils.ResponseError(c, fiber.StatusBadRequest, "Invalid input", nil)
	}

	if errors, err := utils.ValidateInput(product); err != nil {
		return utils.ResponseError(c, fiber.StatusBadRequest, "Invalid input", errors)
	}

	DB.Save(&product)
	return utils.ResponseSuccess(c, "Success update product", product)
}

func (p Product) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product
	if err := DB.First(&product, id).Error; err != nil {
		return utils.ResponseError(c, fiber.StatusNotFound, "Product not found", nil)
	}
	DB.Delete(&product)
	return utils.ResponseSuccess(c, "Success delete product", product)
}

func (p Product) UploadImageProduct(c *fiber.Ctx) error {
	productID := c.Params("id")
	file, err := c.FormFile("image")
	if err != nil {
		return utils.ResponseError(c, fiber.StatusBadRequest, "No file uploaded", nil)
	}
	uploadDir := "./uploads"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		err := os.MkdirAll(uploadDir, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating uploads directory")
		}
	}
	filename := fmt.Sprintf("%s_%s", productID, file.Filename)
	filePath := filepath.Join(uploadDir, filename)
	if err := c.SaveFile(file, filePath); err != nil {
		return utils.ResponseError(c, fiber.StatusInternalServerError, "Failed to save file", nil)
	}
	var product models.Product
	if err := DB.First(&product, productID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Product not found",
		})
	}
	product.ImagePath = filePath
	if err := DB.Save(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to update product image",
		})
	}
	return utils.ResponseSuccess(c, "File upload successfully", fiber.Map{
		"product_id": productID,
		"file_path":  filePath,
	})
}
