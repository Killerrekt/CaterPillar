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

func InsertBattery(c *fiber.Ctx) error {
	var req model.InsertBattery
	if err := c.BodyParser(&req); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{Message: "Failed to parse the json", Status: false})
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{Message: "Didn't pass some value", Status: false})
	}

	ctx := context.Background()
	err := services.InsertBattery(ctx, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{Message: "Something went wrong", Status: false})
	}

	return c.Status(fiber.StatusAccepted).JSON(model.Response{Message: "Inserted successfully", Status: true})
}

func GetBattery(c *fiber.Ctx) error {
	id := c.Params("inspection_id", "")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{Message: "Bad request alert", Status: false})
	}
	uint_id, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{Message: "The id failed to parse to uint", Status: false})
	}

	ctx := context.Background()
	tire, err := services.GetBattery(ctx, uint32(uint_id))
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{Message: "Something went wrong", Status: false})
	}

	return c.Status(fiber.StatusAccepted).JSON(model.Response{Message: "Successfully fetch the detail", Status: true, Data: tire})
}

func UploadBatteryImage(c *fiber.Ctx) error {
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
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{Message: "No file passed under key tire", Status: false})
	}

	tires := files.File["battery"]
	var images []string
	for _, file := range tires {
		id := fmt.Sprint(uuid.New())
		utils.UploadImage(file, id)
		images = append(images, id)
	}

	ctx := context.Background()
	err = services.InsertBatteryImages(ctx, images, uint32(uint_id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{Message: "Failed to save the images id", Status: false})
	}

	return c.Status(fiber.StatusAccepted).JSON(model.Response{Message: "Saved successfully", Status: true})
}

func GetBatteryImages(c *fiber.Ctx) error {
	id := c.Params("inspection_id", "")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{Message: "Bad request alert", Status: false})
	}
	uint_id, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{Message: "The id failed to parse to uint", Status: false})
	}

	ctx := context.Background()
	tire, err := services.GetTireImage(ctx, uint32(uint_id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{Message: "No battery exists for this id", Status: false})
	}
	var urls []string
	fmt.Println(tire)
	for _, v := range tire {
		url, err := utils.GetImage(v)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(model.Response{Message: "Failed to get the images", Status: false})
		}
		urls = append(urls, url)
	}
	return c.Status(fiber.StatusAccepted).JSON(model.Response{Message: "Images retrived successfully testing purpose", Status: true, Data: urls})
}
