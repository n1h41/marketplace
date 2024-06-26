package controllers

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"

	"n1h41/marketplace/views"
)

func GetStoreView(c *fiber.Ctx) error {
	baseView := views.Base(false)
	handler := adaptor.HTTPHandler(templ.Handler(baseView))
	return handler(c)
}
