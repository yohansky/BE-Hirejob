package controller

import (
	"backend-fiber/src/config"
	"backend-fiber/src/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func AllProjects(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(config.DB, &models.Project{}, page))
}

func CreateProject(c *fiber.Ctx) error {
	var project models.Project

	if err := c.BodyParser(&project); err != nil {
		return err
	}

	config.DB.Create(&project)

	return c.JSON(project)
}

func GetProject(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	// karena menggunakan Gorm.Model jadi tidak ada ID, maknaya harus di deklarasikan dulu jika ingin dipanggil IDnya
	var project models.Project

	project.Id = uint(id)

	config.DB.Preload("User").Find(&project)

	return c.JSON(project)
}

func UpdateProject(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	// karena menggunakan Gorm.Model jadi tidak ada kolom ID, maknaya harus di deklarasikan dulu jika ingin dipanggil IDnya
	var project models.Project

	project.Id = uint(id)

	if err := c.BodyParser(&project); err != nil {
		return err
	}

	config.DB.Model(&project).Updates(project)

	return c.JSON(project)
}

func DeleteProject(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var project models.Project

	project.Id = uint(id)

	config.DB.Delete(&project)

	return c.JSON(fiber.Map{
		"Message": "Delete Complete",
	})
}
