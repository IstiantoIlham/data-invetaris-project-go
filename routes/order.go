package routes

import (
	"data-invetaris/controllers"
	"github.com/gofiber/fiber/v2"
)

func OrderRoute(router fiber.Router) {
	orderController := controllers.Order{}
	orders := router.Group("/orders")
	orders.Post("/", orderController.CreateOrder)
	orders.Get("/:id", orderController.GetOrderByID)
}
