package handler

import (
	"github.com/gofiber/fiber/v2"

	"n1h41/marketplace/platform/database"
	"n1h41/marketplace/repositories"
	"n1h41/marketplace/services"
)

func Setup(app *fiber.App) {
	// INFO: ADMIN
	adminRepo := repositories.NewAdminRepo(database.Db)
	adminServ := services.AdminServiceConstructor(adminRepo)
	adminHandler := AdminControllerConstructor(adminServ)

	adminGroup := app.Group("/admin")
	adminGroup.Get("/", adminHandler.GetAdminView)
	adminGroup.Get("/signin", adminHandler.GetAdminSignInView)
	adminGroup.Post("/signin", adminHandler.HandleAdminLogin)
	adminGroup.Get("/products", adminHandler.GetProductSection)
	adminGroup.Get("/products/add", adminHandler.GetAddProductForm)
	adminGroup.Post("/products/add", adminHandler.HandleAddProductFormSubmition)
	adminGroup.Get("/categories", adminHandler.GetCategoryList)
	adminGroup.Get("/categories/add", adminHandler.GetCreateCategoryForm)
}
