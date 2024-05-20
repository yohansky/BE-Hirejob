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

// func CreateProject(c *fiber.Ctx) error {
// 	var project models.Project

// 	if err := c.BodyParser(&project); err != nil {
// 		return err
// 	}

// 	config.DB.Create(&project)

// 	return c.JSON(fiber.Map{
// 		"message": "Project created successfully",
// 		"data":    project,
// 	})
// }

// bisa cuma masalah di worker id
// func CreateProject(c *fiber.Ctx) error {
// 	var project models.Project
// 	// Parsing form-data untuk mendapatkan nilai teks dan file-file yang diunggah
// 	form, err := c.MultipartForm()
// 	if err != nil {
// 		return err
// 	}

// 	if err != nil {
// 		return err
// 	}

// 	// Mendapatkan file gambar dari form-data
// 	files := form.File["Gambar"]
// 	if len(files) == 0 {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": "No image file uploaded",
// 		})
// 	}

// 	// Mengambil file gambar pertama dari form-data
// 	file := files[0]

// 	// Menyimpan file gambar ke dalam direktori uploads
// 	if err := c.SaveFile(file, "src/uploads/"+file.Filename); err != nil {
// 		return err
// 	}

// 	// Membuat objek Project dengan data dari form-data
// 	item := models.Project{
// 		Nama:     project.Nama,
// 		Link:     project.Link,
// 		Tipe:     project.Tipe,
// 		Gambar:   file.Filename,
// 		WorkerId: uint(project.WorkerId),
// 	}

// 	// Menambahkan item baru ke database
// 	config.DB.Create(&item)

// 	// Membuat respons JSON dengan pesan berhasil dan data project
// 	return c.JSON(fiber.Map{
// 		"message": "Project created successfully",
// 		"data":    project,
// 	})
// }

func CreateProject(c *fiber.Ctx) error {
	// Parsing form-data untuk mendapatkan nilai teks dan file-file yang diunggah
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	// Mendapatkan nilai teks dari form-data
	nama := form.Value["Nama"][0]
	link := form.Value["Link"][0]
	tipe := form.Value["Tipe"][0]

	// Mendapatkan nilai WorkerID dari form-data dan mengonversinya ke uint
	workerIDStr := form.Value["WorkerId"][0]
	workerID, err := strconv.ParseUint(workerIDStr, 10, 64)
	if err != nil {
		return err
	}
	workerIDUint := uint(workerID) // Konversi ke uint

	// Mengambil file gambar dari form-data
	files := form.File["Gambar"]
	if len(files) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No image file uploaded",
		})
	}

	// Mengambil file gambar pertama dari form-data
	file := files[0]

	// Menyimpan file gambar ke dalam direktori uploads
	if err := c.SaveFile(file, "src/uploads/"+file.Filename); err != nil {
		return err
	}

	// Membuat objek Project dengan data dari form-data
	project := models.Project{
		Nama:     nama,
		Link:     link,
		Tipe:     tipe,
		Gambar:   "http://localhost:8080/uploads/" + file.Filename,
		WorkerId: workerIDUint, // Menetapkan WorkerID sebagai uint yang sudah dikonversi
	}

	// Menambahkan project baru ke database
	config.DB.Create(&project)

	// Membuat respons JSON dengan pesan berhasil dan data project
	return c.JSON(fiber.Map{
		"message": "Project created successfully",
		"data":    project,
	})
}

func GetProject(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	// karena menggunakan Gorm.Model jadi tidak ada ID, maknaya harus di deklarasikan dulu jika ingin dipanggil IDnya
	var project models.Project

	project.Id = uint(id)

	config.DB.Preload("Worker").Find(&project)

	return c.JSON(project)
}

func GetWorkerByWorkerIDProject(c *fiber.Ctx) error {

	id := c.Params("id")

	var project []models.Project
	if err := config.DB.Where("worker_id = ?", id).First(&project).Error; err != nil {

		return c.JSON(fiber.Map{"error": "Worker not found"})
	}

	return c.JSON(project)
}

func GetProjectsByWorkerID(c *fiber.Ctx) error {

	id := c.Params("id")

	var projects []models.Project
	if err := config.DB.Where("worker_id = ?", id).Find(&projects).Error; err != nil {

		return c.JSON(fiber.Map{"error": "Projects not found"})
	}

	return c.JSON(projects)
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
