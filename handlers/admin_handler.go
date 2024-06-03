package handlers

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"

	adminviews "n1h41/marketplace/views/admin_views"
)

func GetAdminView(c *fiber.Ctx) error {
	adminView := adminviews.AdminView()
	handler := adaptor.HTTPHandler(templ.Handler(adminView))
	return handler(c)
}
