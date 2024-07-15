package controllers

import (
	"errors"
	"log"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"

	"n1h41/marketplace/dto"
	"n1h41/marketplace/services"
	"n1h41/marketplace/utils"
	adminviews "n1h41/marketplace/views/admin_views"
	"n1h41/marketplace/views/partials"
)

type AdminController interface {
	GetAdminView(ctx *fiber.Ctx) error
	GetAdminSignInView(ctx *fiber.Ctx) error
	GetAddProductForm(ctx *fiber.Ctx) error
	GetProductSection(ctx *fiber.Ctx) error
	HandleAdminLogin(ctx *fiber.Ctx) error
	HandleAddProductFormSubmition(c *fiber.Ctx) error
}

type adminController struct {
	service services.AdminService
}

func AdminControllerConstructor(service services.AdminService) AdminController {
	return &adminController{
		service: service,
	}
}

func (c adminController) GetAdminView(ctx *fiber.Ctx) error {
	/* token := c.Cookies("token")
	if token == "" {
		return c.Redirect("/admin/signin")
	} */
	adminView := adminviews.AdminView()
	handler := adaptor.HTTPHandler(templ.Handler(adminView))
	return handler(ctx)
}

func (c adminController) GetAdminSignInView(ctx *fiber.Ctx) error {
	signInView := adminviews.AdminSignInView()
	handler := adaptor.HTTPHandler(templ.Handler(signInView))
	return handler(ctx)
}

func (c adminController) HandleAdminLogin(ctx *fiber.Ctx) error {
	var errors map[string]string
	errors = make(map[string]string)
	var params dto.AdminLoginModel
	if err := ctx.BodyParser(&params); err != nil {
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
		return handler(ctx)
	}

	if err := c.service.Login(params); err != nil {
		log.Println(err)
		errors["loginError"] = err.Error()
		signInComp := partials.SignIn(params, errors)
		handler := adaptor.HTTPHandler(templ.Handler(signInComp))
		return handler(ctx)
	}

	return ctx.Redirect("/admin")
}

func (c adminController) GetAddProductForm(ctx *fiber.Ctx) error {
	addProductForm := partials.AddProductForm()
	handler := adaptor.HTTPHandler(templ.Handler(addProductForm))
	return handler(ctx)
}

func (c adminController) GetProductSection(ctx *fiber.Ctx) error {
	productSection := partials.AdminProductSection()
	handler := adaptor.HTTPHandler(templ.Handler(productSection))
	return handler(ctx)
}

func (c adminController) HandleAddProductFormSubmition(ctx *fiber.Ctx) error {
	var params dto.AddProductModel
	if err := ctx.BodyParser(&params); err != nil {
		return err
	}
	form, err := ctx.MultipartForm()
	if err != nil {
		return err
	}
	productImageFiles := form.File["productImageFiles"]
	if len(productImageFiles) == 0 {
		return errors.New("No product images uploaded")
	}
	// INFO: This is the part where we are setting the productImageFiles to the params
	for _, file := range productImageFiles {
		params.ProductImageFiles = append(params.ProductImageFiles, file)
	}
	return nil
}
