package router

import (
	handlers "CRUD_API_FIBER/handlers/projects"

	"github.com/gofiber/fiber/v2"
)

func AddProjectGroup(app *fiber.App) {
	projects := app.Group("/projects")
	projects.Get("/", handlers.GetProjects)
	projects.Get("/:id", handlers.GetProject)
	projects.Post("/", handlers.CreateProject)
	projects.Post("/:id", handlers.AssignTask)
	projects.Delete("/:id", handlers.DeleteProject)
}
