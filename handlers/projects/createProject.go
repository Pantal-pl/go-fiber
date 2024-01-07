package handlers

import (
	"CRUD_API_FIBER/common"
	"CRUD_API_FIBER/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

type createProjectDTO struct {
	Title        string        `json:"title" bson:"title"`
	Author       string        `json:"author" bson:"author"`
	CreationDate time.Time     `json:"creationDate,omitempty" bson:"creationDate,omitempty"`
	Tasks        []models.Task `json:"tasks,omitempty" bson:"tasks,omitempty"`
}

func CreateProject(c *fiber.Ctx) error {
	project := new(createProjectDTO)
	err := c.BodyParser(project)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}
	project.CreationDate = time.Now()
	coll := common.GetDbCollection("projects")
	// if project.Tasks == nil {

	// }
	result, err := coll.InsertOne(c.Context(), project)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Falied to create a project",
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"result": fiber.Map{
			"insertedID": result.InsertedID,
			"project":    project,
		},
	})
}
