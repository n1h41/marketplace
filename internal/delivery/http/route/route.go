package route

import (
	"github.com/gofiber/fiber/v2"

	"n1h41/marketplace/internal/delivery/http/handler"
	"n1h41/marketplace/internal/infrastructure/database"
	"n1h41/marketplace/internal/repository"
	"n1h41/marketplace/internal/usecase"
)

func RegisterRoutes(app *fiber.App) {
	adminRepo := repository.NewAdminRepo(database.Db)
	productRepo := repository.NewProductRepo(database.Db)
	adminUsc := usecase.NewAdminUsc(adminRepo, productRepo)
	adminHandler := handler.NewAdminHandler(adminUsc)

	app.Group("/admin").Get("/", adminHandler.GetAdminView).
		Get("/signin", adminHandler.GetAdminSignInView).Post("/signin", adminHandler.HandleAdminLogin).
		Get("/products", adminHandler.GetProductSection).Get("/products/add", adminHandler.GetAddProductForm).
		Post("/products/add", adminHandler.HandleAddProductFormSubmition).Get("/categories", adminHandler.GetCategoryList).
		Get("/categories/add", adminHandler.GetCreateCategoryForm).Post("/categories/add", adminHandler.HandleCreateCategoryForm)
}

