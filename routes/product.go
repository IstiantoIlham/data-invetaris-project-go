package routes

import (
	"data-invetaris/controllers"
	"github.com/gofiber/fiber/v2"
)

func ProductRoute(router fiber.Router) {
	productController := controllers.Product{}

	products := router.Group("/products")
	products.Get("/", productController.Get)
	products.Get("/:id", productController.GetById)
	products.Post("/", productController.Create)
	products.Put("/:id", productController.Update)
	products.Delete("/:id", productController.Delete)

	products.Post("/upload-image/:id", productController.UploadImageProduct)
}
