package route

import (
	"github.com/gofiber/fiber/v2"

	"n1h41/marketplace/internal/delivery/http/handler"
	"n1h41/marketplace/internal/infrastructure/database"
	"n1h41/marketplace/internal/repository/adminrepo"
	"n1h41/marketplace/internal/repository/productrepo"
	"n1h41/marketplace/internal/usecase/adminusc"
)

func RegisterRoutes(app *fiber.App) {
	// INFO: ADMIN
	adminRepo := adminrepo.NewAdminRepo(database.Db)
	productRepo := productrepo.NewProductRepo(database.Db)
	adminUsc := adminusc.NewAdminUsc(adminRepo, productRepo)
	adminHandler := handler.NewAdminHandler(adminUsc)

	adminGroup := app.Group("/admin")
	adminGroup.Get("/", adminHandler.GetAdminView)
	adminGroup.Get("/signin", adminHandler.GetAdminSignInView)
	adminGroup.Post("/signin", adminHandler.HandleAdminLogin)
	adminGroup.Get("/products", adminHandler.GetProductSection)
	adminGroup.Get("/products/add", adminHandler.GetAddProductForm)
	adminGroup.Post("/products/add", adminHandler.HandleAddProductFormSubmition)
	adminGroup.Get("/categories", adminHandler.GetCategoryList)
	adminGroup.Get("/categories/add", adminHandler.GetCreateCategoryForm)
	adminGroup.Post("/categories/add", adminHandler.HandleCreateCategoryForm)
}
