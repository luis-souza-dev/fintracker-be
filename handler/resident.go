package handler

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/luis-souza-dev/fintracker-be/database"
	"github.com/luis-souza-dev/fintracker-be/models"
)

func GetResident(c *fiber.Ctx) error {
	result := &[]models.Resident{}
	query := &models.Resident{}
	if err := c.QueryParser(query); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(fmt.Sprintf("query err%v", err))
	}
	
	database.DB.Where(query).Find(result)

	jsonval, err := json.Marshal(result);
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(fmt.Sprintf("json err %v", err))
	}
	return c.Status(200).SendString(string(jsonval))
}

func CreateResident(c *fiber.Ctx) error {
	resident := new(models.Resident)
	
	if err := c.BodyParser(resident); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(fmt.Sprintf("body err%v", err))
	}

	if resident.Name == "" {
		return c.Status(fiber.ErrBadRequest.Code).SendString("missing name")
	}

	if resident.Birthday == "" {
		return c.Status(fiber.ErrBadRequest.Code).SendString("missing date")
	}

	res := database.DB.Create(resident)
	if res.Error != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString("smthg went wrong")
	}

	return c.Status(200).SendString(fmt.Sprint(resident.ID))
}

func EditResident(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.ErrBadRequest.Code).SendString("Missing id")
	}

	resident := new(models.Resident)
	reqBody := new(models.Resident)

	if err := c.BodyParser(reqBody); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(fmt.Sprintf("body err%v", err))
	}

	intId, err := strconv.Atoi(id)

	if err == nil {
		resident.ID = uint(intId)
		res := database.DB.Model(resident).Updates(reqBody)

		if res.Error == nil {
			return c.Status(200).SendString(fmt.Sprintf("body err %v, \n and queryBody %v", resident, reqBody))
		} else {
			return c.Status(fiber.ErrBadRequest.Code).SendString("Error while updating")
		}
	} else {
		return c.Status(fiber.ErrBadRequest.Code).SendString(fmt.Sprintf("incorrect id value: %v", err))
	}
}

func DeleteResident(c *fiber.Ctx) error {
	id := c.Params("id");
	if id == "" {
		return c.Status(fiber.ErrBadRequest.Code).SendString("Missing id")
	}
	res := database.DB.Delete(&models.Resident{}, id)

	if res.Error != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString("Deletion err")
	}
	return c.Status(200).SendString("done")
}
