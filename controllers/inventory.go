package controllers

import (
	. "data-invetaris/database"
	"data-invetaris/models"
	"data-invetaris/utils"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type Inventory struct{}

func (i Inventory) GetStockProductID(c *fiber.Ctx) error {
	productID, err := strconv.Atoi(c.Params("product_id"))
	if err != nil {
		return utils.ResponseError(c, fiber.StatusBadRequest, "Invalid product ID", nil)
	}

	var inventory models.Inventory

	if err := DB.Where("product_id = ?", productID).First(&inventory).Error; err != nil {
		return utils.ResponseError(c, fiber.StatusInternalServerError, "Inventory not found", nil)
	}

	return utils.ResponseSuccess(c, "Inventory fetched successfully", inventory)
}

func (i Inventory) UpdateOrCreateStockByProductID(c *fiber.Ctx) error {
	productID, err := strconv.Atoi(c.Params("product_id"))
	if err != nil {
		return utils.ResponseError(c, fiber.StatusBadRequest, "Invalid product ID", nil)
	}

	type StockInput struct {
		Stock int `json:"stock" validate:"required"`
	}

	var input StockInput
	if err := c.BodyParser(&input); err != nil {
		return utils.ResponseError(c, fiber.StatusBadRequest, "Invalid input", nil)
	}

	var inventory models.Inventory

	if err := DB.Where("product_id = ?", productID).First(&inventory).Error; err != nil {
		newInventory := models.Inventory{
			ProductID: uint(productID),
			Stock:     input.Stock,
		}

		if errors, err := utils.ValidateInput(newInventory); err != nil {
			return utils.ResponseError(c, fiber.StatusBadRequest, "Invalid input", errors)
		}

		if err := DB.Create(&newInventory).Error; err != nil {
			return utils.ResponseError(c, fiber.StatusInternalServerError, "Inventory creation failed", nil)
		}
		return utils.ResponseSuccess(c, "Inventory created successfully", inventory)
	}

	inventory.Stock = input.Stock

	if inventory.Stock < 0 {
		return utils.ResponseError(c, fiber.StatusBadRequest, "Stock cannot be negative", nil)
	}

	if err := DB.Save(&inventory).Error; err != nil {
		return utils.ResponseError(c, fiber.StatusInternalServerError, "Failed to update inventory", nil)
	}
	return utils.ResponseSuccess(c, "Inventory updated successfully", nil)
}
