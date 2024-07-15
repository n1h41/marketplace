package main

import (
	"github.com/gofiber/fiber/v2"

	"n1h41/marketplace/controllers"
	"n1h41/marketplace/platform/database"
	"n1h41/marketplace/repositories"
	"n1h41/marketplace/services"
)

func Setup(app *fiber.App) {
	/* storeGroup := app.Group("/store")
	storeGroup.Get("/", GetStoreView) */

	// INFO: ADMIN
	repository := repositories.AdminRepoConstructor(database.Db)
	service := services.AdminServiceConstructor(repository)
	controller := controllers.AdminControllerConstructor(service)

	adminGroup := app.Group("/admin")
	adminGroup.Get("/", controller.GetAdminView)
	adminGroup.Get("/signin", controller.GetAdminSignInView)
	adminGroup.Post("/signin", controller.HandleAdminLogin)
	adminGroup.Get("/products", controller.GetProductSection)
	adminGroup.Get("/products/add", controller.GetAddProductForm)
	adminGroup.Post("/products/add", controller.HandleAddProductFormSubmition)
}
