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

func InsertTire(c *fiber.Ctx) error {
	var req model.InsertTire
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{Message: "Failed to parse the json", Status: false})
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{Message: "Didn't pass some value", Status: false})
	}

	ctx := context.Background()
	err := services.InsertTire(ctx, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{Message: "Something went wrong", Status: false})
	}

	return c.Status(fiber.StatusAccepted).JSON(model.Response{Message: "Inserted successfully", Status: true})
}

func GetTire(c *fiber.Ctx) error {
	id := c.Params("inspection_id", "")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{Message: "Bad request alert", Status: false})
	}
	uint_id, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{Message: "The id failed to parse to uint", Status: false})
	}

	ctx := context.Background()
	tire, err := services.GetTire(ctx, uint32(uint_id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{Message: "Something went wrong", Status: false})
	}

	return c.Status(fiber.StatusAccepted).JSON(model.Response{Message: "Successfully fetch the detail", Status: true, Data: tire})
}

func UploadTireImage(c *fiber.Ctx) error {
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

	tires := files.File["tire"]
	var images []string
	for _, file := range tires {
		id := fmt.Sprint(uuid.New())
		utils.UploadImage(file, id)
		images = append(images, id)
	}

	ctx := context.Background()
	err = services.InsertTireImages(ctx, images, uint32(uint_id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{Message: "Failed to save the images id", Status: false})
	}

	return c.Status(fiber.StatusAccepted).JSON(model.Response{Message: "Saved successfully", Status: true})
}

func GetTireImages(c *fiber.Ctx) error {
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
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{Message: "No tire exists for this id", Status: false})
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

	if len(tire) != 4 {
		return c.Status(fiber.StatusAccepted).JSON(model.Response{Message: "Images retrived successfully testing purpose", Status: true, Data: urls})
	}

	res := struct {
		FrontLeftTire  string
		FrontRightTire string
		BackLeftTire   string
		BackRightTire  string
	}{urls[0], urls[1], urls[2], urls[3]}

	return c.Status(fiber.StatusAccepted).JSON(model.Response{Message: "Images retrived successfully", Status: true, Data: res})
}
