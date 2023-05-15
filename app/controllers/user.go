package controllers

import (
	"fmt"

	api_model "github.com/FurkanSamaraz/Dependency-Injection/app/pkg/model"
	api_service "github.com/FurkanSamaraz/Dependency-Injection/app/pkg/services"
	api_structure "github.com/FurkanSamaraz/Dependency-Injection/app/structures"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	service   *api_service.UserService
	userModel *api_model.UserModel
}

func NewUserController(service *api_service.UserService, model *api_model.UserModel) *UserController {
	return &UserController{
		service:   service,
		userModel: model,
	}
}

// RegisterHandler godoc
// @Summary       Register User
// @Description   Registers a new user
// @Tags          User
// @Accept        json
// @Produce       json
// @Param         body body api_structure.User true "Request body"
// @Success       200 {object} api_structure.Response
// @Failure       500 {object} api_structure.ErrorResponse
// @Router        /register [post]
func (controller *UserController) RegisterHandler(c *fiber.Ctx) error {
	var user api_structure.User
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(api_structure.ErrorResponse{
			Type:    "Fetch Data",
			Message: "Invalid request",
		})
	}

	result := controller.userModel.Register(&user)
	if user.Username == "" || user.Password == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(api_structure.ErrorResponse{
			Type:    "Fetch Data",
			Message: "Username and password are required",
		})

	}

	results, err := controller.service.Register(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(api_structure.ErrorResponse{
			Type:    "Fetch Data",
			Message: err.Error(),
		})
	}

	fmt.Println(results)
	result.Message = result.Jwt

	return c.Status(fiber.StatusOK).JSON(result)
}

// LoginHandler godoc
// @Summary       Login User
// @Description   Logs in a user
// @Tags          User
// @Accept        json
// @Produce       json
// @Param         body body api_structure.User true "Request body"
// @Success       200 {object} api_structure.Response
// @Failure       500 {object} api_structure.ErrorResponse
// @Router        /login [post]
func (controller *UserController) LoginHandler(c *fiber.Ctx) error {
	var user api_structure.User
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(api_structure.ErrorResponse{
			Type:    "Fetch Data",
			Message: "Invalid request",
		})
	}
	result, rerr := controller.service.Login(user)
	if user.Username == "" || user.Password == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(api_structure.ErrorResponse{
			Type:    "Fetch Data",
			Message: "Username and password are required",
		})
	}

	if rerr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(api_structure.ErrorResponse{
			Type:    "Create Data",
			Message: rerr.Error(),
		})
	}
	c.Locals("Authorization", result)

	res := controller.userModel.Login(&user)
	res.Jwt = result

	return c.Status(fiber.StatusOK).JSON(res)
}
