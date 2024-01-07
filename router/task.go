package router

import (
	handlers "CRUD_API_FIBER/handlers/tasks"

	"github.com/gofiber/fiber/v2"
)

func AddTaskGroup(app *fiber.App) {
	tasks := app.Group("/tasks")
	tasks.Get("/", handlers.GetTasks)
	tasks.Get("/:id", handlers.GetTask)
	tasks.Post("/:id", handlers.CreateTask)
	tasks.Delete("/:id", handlers.DeleteTask)
}
