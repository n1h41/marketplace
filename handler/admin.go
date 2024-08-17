package handler

import (
	"errors"
	"log"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"

	"n1h41/marketplace/dto"
	"n1h41/marketplace/services"
	"n1h41/marketplace/utils"
	"n1h41/marketplace/views/admin"
	"n1h41/marketplace/views/partials"
)

type AdminHandler interface {
	GetAdminView(ctx *fiber.Ctx) error
	GetAdminSignInView(ctx *fiber.Ctx) error
	GetAddProductForm(ctx *fiber.Ctx) error
	GetProductSection(ctx *fiber.Ctx) error
	HandleAdminLogin(ctx *fiber.Ctx) error
	HandleAddProductFormSubmition(c *fiber.Ctx) error
	GetCategoryList(ctx *fiber.Ctx) error
	GetCreateCategoryForm(ctx *fiber.Ctx) error
}

type adminHandler struct {
	service services.AdminService
}

func (a *adminHandler) GetCreateCategoryForm(ctx *fiber.Ctx) error {
	view := partials.CreateCategory()
	handler := adaptor.HTTPHandler(templ.Handler(view))
	return handler(ctx)
}

func AdminControllerConstructor(service services.AdminService) AdminHandler {
	return &adminHandler{
		service: service,
	}
}

func (a adminHandler) GetAdminView(ctx *fiber.Ctx) error {
	adminView := admin.AdminView()
	handler := adaptor.HTTPHandler(templ.Handler(adminView))
	return handler(ctx)
}

func (a adminHandler) GetAdminSignInView(ctx *fiber.Ctx) error {
	signInView := admin.AdminSignInView()
	handler := adaptor.HTTPHandler(templ.Handler(signInView))
	return handler(ctx)
}

func (a adminHandler) HandleAdminLogin(ctx *fiber.Ctx) error {
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

	if err := a.service.Login(params); err != nil {
		log.Println(err)
		errors["loginError"] = err.Error()
		signInComp := partials.SignIn(params, errors)
		handler := adaptor.HTTPHandler(templ.Handler(signInComp))
		return handler(ctx)
	}

	return ctx.Redirect("/admin")
}

func (a adminHandler) GetAddProductForm(ctx *fiber.Ctx) error {
	addProductForm := partials.AddProductForm()
	handler := adaptor.HTTPHandler(templ.Handler(addProductForm))
	return handler(ctx)
}

func (a adminHandler) GetProductSection(ctx *fiber.Ctx) error {
	productSection := partials.AdminProductSection()
	handler := adaptor.HTTPHandler(templ.Handler(productSection))
	return handler(ctx)
}

func (a adminHandler) HandleAddProductFormSubmition(ctx *fiber.Ctx) error {
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

func (a *adminHandler) GetCategoryList(ctx *fiber.Ctx) error {
	view := partials.ListCategories()
	handler := adaptor.HTTPHandler(templ.Handler(view))
	return handler(ctx)
}
