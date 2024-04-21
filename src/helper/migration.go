package helper

import (
	"backend-fiber/src/config"
	"backend-fiber/src/models"
)

func Migrate() {
	config.DB.AutoMigrate(&models.User{}, &models.Role{}, &models.Experience{}, &models.Project{}, &models.Recruiter{}, &models.Skill{}, &models.Worker{})
}
