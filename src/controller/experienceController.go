package controller

import (
	"backend-fiber/src/config"
	"backend-fiber/src/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func AllExperiences(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(config.DB, &models.Experience{}, page))
}

func CreateExperience(c *fiber.Ctx) error {
	var experience models.Experience

	if err := c.BodyParser(&experience); err != nil {
		return err
	}

	config.DB.Create(&experience)

	return c.JSON(experience)
}

func GetExperience(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	// karena menggunakan Gorm.Model jadi tidak ada ID, maknaya harus di deklarasikan dulu jika ingin dipanggil IDnya
	var experience models.Experience

	experience.Id = uint(id)

	config.DB.Preload("User").Find(&experience)

	return c.JSON(experience)
}

func UpdateExperience(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	// karena menggunakan Gorm.Model jadi tidak ada kolom ID, maknaya harus di deklarasikan dulu jika ingin dipanggil IDnya
	var experience models.Experience

	experience.Id = uint(id)

	if err := c.BodyParser(&experience); err != nil {
		return err
	}

	config.DB.Model(&experience).Updates(experience)

	return c.JSON(experience)
}

func DeleteExperience(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var experience models.Experience

	experience.Id = uint(id)

	config.DB.Delete(&experience)

	return c.JSON(fiber.Map{
		"Message": "Delete Complete",
	})
}
