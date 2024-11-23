package routes

import "github.com/gofiber/fiber/v2"

func Routes(app *fiber.App) {
	api := app.Group("/api")

	ProductRoute(api)
	InventoryRoute(api)
	OrderRoute(api)
}
