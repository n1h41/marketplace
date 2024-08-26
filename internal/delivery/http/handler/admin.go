package handler

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"

	"n1h41/marketplace/internal/domain/productdmn"
	"n1h41/marketplace/internal/model"
	"n1h41/marketplace/internal/usecase"
	"n1h41/marketplace/internal/utils"
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
	HandleCreateCategoryForm(ctx *fiber.Ctx) error
}

type adminHandler struct {
	usecase usecase.AdminUsecase
}

func NewAdminHandler(usecase usecase.AdminUsecase) AdminHandler {
	return &adminHandler{
		usecase: usecase,
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
	var params model.LoginAdminUserRequest
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

	if err := a.usecase.Login(params); err != nil {
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
	var params model.AddProductRequest
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
	// TODO: Need to add logic to store files into an object storage and then save the
	//       file urls tp the database
	for _, file := range productImageFiles {
		params.ProductImageFiles = append(params.ProductImageFiles, file)
	}
	return nil
}

func (a *adminHandler) GetCategoryList(ctx *fiber.Ctx) (err error) {
	var categoryList []productdmn.Category
	categoryList, err = a.usecase.ListCategories()
	view := partials.ListCategories(categoryList)
	handler := adaptor.HTTPHandler(templ.Handler(view))
	return handler(ctx)
}

func (a *adminHandler) HandleCreateCategoryForm(ctx *fiber.Ctx) error {
	var params model.AddCategoryReqeust
	if err := ctx.BodyParser(&params); err != nil {
		return err
	}
	log.Println(params)
	newCategory := productdmn.Category{
		Name:          params.CategoryName,
		IsSubCategory: params.IsSubCategory,
	}
	if params.IsSubCategory {
		parentId, err := strconv.ParseInt(params.ParentId, 10, 0)
		if err != nil {
			return err
		}
		newCategory.Parent = int(parentId)
	}
	if err := a.usecase.CreateCategory(newCategory); err != nil {
		return err
	}
	return ctx.SendStatus(http.StatusCreated) /* .Redirect("/admin/categories") */
}

func (a *adminHandler) GetCreateCategoryForm(ctx *fiber.Ctx) (err error) {
	var categoryList []productdmn.Category
	categoryList, err = a.usecase.ListCategories()
	view := partials.CreateCategory(categoryList)
	handler := adaptor.HTTPHandler(templ.Handler(view))
	return handler(ctx)
}
