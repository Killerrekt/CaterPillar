package controller

import (
	"context"
	"time"

	"github.com/ScoobieNoobie/Caterpillar/config"
	"github.com/ScoobieNoobie/Caterpillar/internal/model"
	"github.com/ScoobieNoobie/Caterpillar/internal/services"
	"github.com/ScoobieNoobie/Caterpillar/internal/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *fiber.Ctx) error {
	var req model.User
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{Message: "Failed to parse the json", Status: false})
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{Message: "Didn't pass all the value", Status: false})
	}

	hashedPassword, errr := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if errr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{Message: "Failed to create a password hash", Status: false})
	}
	req.Password = string(hashedPassword)

	ctx := context.Background()
	err := services.InsertUser(ctx, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{Message: "Failed to add the user", Status: false})
	}

	return c.Status(fiber.StatusAccepted).JSON(model.Response{Message: "User added successfully", Status: true})
}

func Login(c *fiber.Ctx) error {
	var req struct {
		EmployeeID string `json:"employee_id" validate:"required"`
		Password   string `json:"password" validate:"required"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{Message: "Failed to parse the json", Status: false})
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{Message: "Didn't pass all the value", Status: false})
	}

	ctx := context.Background()
	user, err := services.GetUser(ctx, req.EmployeeID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{Message: "Employee id doesn't exists", Status: false})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(model.Response{Message: "The password doesn't match", Status: false})
	}

	access_key, err := utils.GenerateToken(req.EmployeeID, config.ConfigData.AccessTokenSecret, time.Hour)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{Message: "Failed to generate a jwt token", Status: false})
	}

	return c.Status(fiber.StatusAccepted).JSON(model.Response{Message: "The user has been logged in successfully", Status: true, Data: access_key})
}
