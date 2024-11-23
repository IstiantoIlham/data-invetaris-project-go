package controllers

import (
	. "data-invetaris/database"
	"data-invetaris/models"
	"data-invetaris/utils"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type Order struct{}

func (o Order) CreateOrder(c *fiber.Ctx) error {
	type OrderInput struct {
		ProductID uint `json:"product_id" validate:"required"`
		Quantity  int  `json:"quantity" validate:"required,gte=1"`
	}

	var input OrderInput

	if err := c.BodyParser(&input); err != nil {
		return utils.ResponseError(c, fiber.StatusBadRequest, "Invalid input", nil)
	}

	if errors, err := utils.ValidateInput(input); err != nil {
		return utils.ResponseError(c, fiber.StatusBadRequest, "Validation failed", errors)
	}

	var products models.Product

	if err := DB.First(&products, input.ProductID).Error; err != nil {
		return utils.ResponseError(c, fiber.StatusNotFound, "Product not found", nil)
	}
	var inventory models.Inventory
	if err := DB.Where("product_id = ?", input.ProductID).First(&inventory).Error; err != nil {
		return utils.ResponseError(c, fiber.StatusNotFound, "Inventory not found for the product", nil)
	}
	if inventory.Stock < input.Quantity {
		return utils.ResponseError(c, fiber.StatusBadRequest, "Insufficient stock", nil)
	}
	order := models.Order{
		ProductID: input.ProductID,
		Quantity:  input.Quantity,
		Status:    "Pending",
	}
	if err := DB.Create(&order).Error; err != nil {
		return utils.ResponseError(c, fiber.StatusInternalServerError, "Failed to create order", nil)
	}
	inventory.Stock -= input.Quantity
	DB.Save(&inventory)
	return utils.ResponseSuccess(c, "Order created successfully", order)

}

func (o Order) GetOrderByID(c *fiber.Ctx) error {
	orderID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.ResponseError(c, fiber.StatusBadRequest, "Invalid order ID", nil)
	}

	var order models.Order
	if err := DB.Preload("Product").First(&order, orderID).Error; err != nil {
		return utils.ResponseError(c, fiber.StatusNotFound, "Order not found", nil)
	}

	return utils.ResponseSuccess(c, "Order fetched successfully", order)
}
