package handlers

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"

	"n1h41/marketplace/models"
	"n1h41/marketplace/utils"
	adminviews "n1h41/marketplace/views/admin_views"
	"n1h41/marketplace/views/partials"
)

func GetAdminView(c *fiber.Ctx) error {
	token := c.Cookies("token")
	if token == "" {
		return c.Redirect("/admin/signin")
	}
	adminView := adminviews.AdminView()
	handler := adaptor.HTTPHandler(templ.Handler(adminView))
	return handler(c)
}

func GetAdminSignInView(c *fiber.Ctx) error {
	signInView := adminviews.AdminSignInView()
	handler := adaptor.HTTPHandler(templ.Handler(signInView))
	return handler(c)
}

func HandleSignIn(c *fiber.Ctx) error {
	var errors map[string]string
	errors = make(map[string]string)
	var params models.AdminLoginModel
	if err := c.BodyParser(&params); err != nil {
		return err
	}
	if params.Email == "" {
		errors["email"] = "Email is required*"
	} else if !utils.ValidateEmail(params.Email) {
		errors["email"] = "Enter a valid email"
	}
	if params.Password == "" {
		errors["password"] = "Password is required*"
	} else if len(params.Password) < 8 {
		errors["password"] = "Password should be atleast 8 char long"
	}
	if len(errors) > 0 {
		signInComp := partials.SignIn(params, errors)
		handler := adaptor.HTTPHandler(templ.Handler(signInComp))
		return handler(c)
	}

	// TODO: call admin service to login user

	return c.Redirect("/admin")
}
