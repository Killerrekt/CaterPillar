package controller

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ScoobieNoobie/Caterpillar/internal/model"
	"github.com/ScoobieNoobie/Caterpillar/internal/services"
	"github.com/ScoobieNoobie/Caterpillar/internal/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateEngine(c *fiber.Ctx) error {
	var req model.GetEngine
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{Message: "Failed to parse the json", Status: false})
	}

	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{Message: "Didn't pass all the fields", Status: false})
	}

	ctx := context.Background()
	err = services.CreateEngine(ctx, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{Message: "Something went wrong", Status: false})
	}

	return c.Status(fiber.StatusAccepted).JSON(model.Response{Message: "Added successfully", Status: true})
}

func AddEngineImage(c *fiber.Ctx) error {
	id := c.Params("inspection_id", "")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{Message: "Bad request alert", Status: false})
	}
	uint_id, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{Message: "The id failed to parse to uint", Status: false})
	}
	files, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{Message: "No file passed under key Engine", Status: false})
	}

	tires := files.File["Engine"]
	var images string
	for _, file := range tires {
		id := fmt.Sprint(uuid.New())
		utils.UploadImage(file, id)
		images = id
	}

	ctx := context.Background()
	err = services.InsertEngineImage(ctx, images, uint32(uint_id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{Message: "Failed to save the images id", Status: false})
	}

	return c.Status(fiber.StatusAccepted).JSON(model.Response{Message: "Saved successfully", Status: true})
}

func AddEngineImages(c *fiber.Ctx) error {
	id := c.Params("inspection_id", "")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{Message: "Bad request alert", Status: false})
	}
	uint_id, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{Message: "The id failed to parse to uint", Status: false})
	}
	files, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{Message: "No file passed under key Engine", Status: false})
	}

	tires := files.File["Engine"]
	var images []string
	for _, file := range tires {
		id := fmt.Sprint(uuid.New())
		utils.UploadImage(file, id)
		images = append(images, id)
	}

	ctx := context.Background()
	err = services.InsertEngineImages(ctx, images, uint32(uint_id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{Message: "Failed to save the images id", Status: false})
	}

	return c.Status(fiber.StatusAccepted).JSON(model.Response{Message: "Saved successfully", Status: true})
}

func GetEngine(c *fiber.Ctx) error {
	id := c.Params("inspection_id", "")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{Message: "Bad request alert", Status: false})
	}
	uint_id, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{Message: "The id failed to parse to uint", Status: false})
	}

	ctx := context.Background()
	tire, err := services.GetEngine(ctx, uint32(uint_id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{Message: "Something went wrong", Status: false})
	}

	return c.Status(fiber.StatusAccepted).JSON(model.Response{Message: "Successfully fetch the detail", Status: true, Data: tire})
}
