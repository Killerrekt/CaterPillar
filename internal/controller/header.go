package controller

import (
	"context"
	"fmt"

	"github.com/ScoobieNoobie/Caterpillar/internal/model"
	"github.com/ScoobieNoobie/Caterpillar/internal/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetHeader(c *fiber.Ctx) error {
	emp_id := fmt.Sprint(c.Locals("EmployeeID"))
	ctx := context.Background()
	header, err := services.GetHeader(ctx, emp_id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{Message: "There is no header for this inspector", Status: false})
	}
	return c.Status(fiber.StatusAccepted).JSON(model.Response{Message: "Successfully fetch the header for the inspector", Status: true, Data: header})
}

func UploadHeader(c *fiber.Ctx) error {
	var req model.CreateHeader
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{Message: "Failed to parse the json", Status: false})
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{Message: "Didn't pass all the value", Status: false})
	}

	req.InspectorEmployeeID = fmt.Sprint(c.Locals("EmployeeID"))

	ctx := context.Background()
	err := services.InsertHeader(ctx, req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{Message: "Failed to insert into header", Status: false})
	}
	return c.Status(fiber.StatusAccepted).JSON(model.Response{Message: "Successfully added the header", Status: true})
}
