package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	storeGroup := app.Group("/store")
	storeGroup.Get("/", GetStoreView)

	adminGroup := app.Group("/admin")
	adminGroup.Get("/", GetAdminView)
}
