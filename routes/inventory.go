package routes

import (
	"data-invetaris/controllers"
	"github.com/gofiber/fiber/v2"
)

func InventoryRoute(router fiber.Router) {
	inventoryController := controllers.Inventory{}

	inventory := router.Group("/inventory")
	inventory.Get("/:product_id", inventoryController.GetStockProductID)
	inventory.Put("/:product_id", inventoryController.UpdateOrCreateStockByProductID)
}
