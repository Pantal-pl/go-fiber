package handlers

import (
	"CRUD_API_FIBER/common"
	"CRUD_API_FIBER/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetProjects(c *fiber.Ctx) error {
	coll := common.GetDbCollection("projects")
	projects := make([]models.Project, 0)

	cursor, err := coll.Find(c.Context(), bson.M{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	defer cursor.Close(c.Context()) // Remember to close the cursor when done

	// Iterate over the cursor and decode results
	for cursor.Next(c.Context()) {
		project := models.Project{}
		if err := cursor.Decode(&project); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		projects = append(projects, project)
	}

	if err := cursor.Err(); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"projects": projects,
	})
}
