package routes

import (
	"backend-fiber/src/controller"
	"backend-fiber/src/middleware"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	app.Post("/register", controller.Register)
	app.Post("/login", controller.Login)

	app.Use(middleware.IsAuth)

	app.Put("/user/info", controller.UpdateInfor)
	app.Put("/user/password", controller.UpdatePassword)

	app.Get("/user", controller.User)
	app.Post("/logout", controller.Logout)

	app.Get("/users", controller.AllUsers)
	app.Post("/users", controller.CreateUser)
	app.Get("/user/:id", controller.GetUser)
	app.Put("/user/:id", controller.UpdateUser)
	app.Delete("/user/:id", controller.DeleteUser)

	app.Get("/roles", controller.AllRoles)
	app.Post("/roles", controller.CreateRole)
	app.Get("/role/:id", controller.GetRole)
	app.Delete("/role/:id", controller.DeleteRole)

	app.Get("/workers", controller.AllWorkers)
	app.Post("/workers", controller.CreateWorker)
	app.Get("/worker/:id", controller.GetWorker)
	//buat kondisi ketika sudah ada userid akan menemukan workerid
	app.Get("/user/:id/worker", controller.GetWorkerByUserID)
	app.Put("/worker/:id", controller.UpdateWorker)
	app.Delete("/worker/:id", controller.DeleteWorker)

	app.Get("/recruiters", controller.AllRecruiters)
	app.Post("/recruiters", controller.CreateRecruiter)
	app.Get("/recruiter/:id", controller.GetRecruiter)
	app.Put("/recruiter/:id", controller.UpdateRecruiter)
	app.Delete("/recruiter/:id", controller.DeleteRecruiter)

	app.Get("/skills", controller.AllSkills)
	app.Post("/skills", controller.CreateSkill)
	app.Get("/skill/:id", controller.GetSkill)
	app.Put("/skill/:id", controller.UpdateSkill)
	app.Delete("/skill/:id", controller.DeleteSkill)

	app.Get("/projects", controller.AllProjects)
	app.Post("/projects", controller.CreateProject)
	app.Get("/project/:id", controller.GetProject)
	app.Get("/worker/:id/project", controller.GetWorkerByWorkerIDProject)
	app.Get("/worker/:id/projects", controller.GetProjectsByWorkerID)
	app.Put("/project/:id", controller.UpdateProject)
	app.Delete("/project/:id", controller.DeleteProject)

	app.Get("/experiences", controller.AllExperiences)
	app.Post("/experiences", controller.CreateExperience)
	app.Get("/experience/:id", controller.GetExperience)
	app.Get("/worker/:id/experience", controller.GetWorkerByWorkerIDExperience)
	app.Get("/worker/:id/experiences", controller.GetExperiencesByWorkerID)
	app.Put("/experience/:id", controller.UpdateExperience)
	app.Delete("/experience/:id", controller.DeleteExperience)

	app.Static("/uploads", "src/uploads")
}
