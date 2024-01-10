package handlers

import (
	"CRUD_API_FIBER/common"
	"CRUD_API_FIBER/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteProject(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "id is required",
		})
	}

	projectObjectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	collProjects := common.GetDbCollection("projects")

	project := models.Project{}

	err = collProjects.FindOneAndDelete(c.Context(), bson.M{"_id": projectObjectId}).Decode(&project)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "failed to delete a project",
			"message": err.Error(),
		})
	}

	collTasks := common.GetDbCollection("tasks")
	var taskDeletionErrors []error

	for _, v := range project.Tasks {
		taskObjectId, err := primitive.ObjectIDFromHex(v.ID)
		if err != nil {
			taskDeletionErrors = append(taskDeletionErrors, err)
			continue
		}

		_, err = collTasks.DeleteOne(c.Context(), bson.M{"_id": taskObjectId})

		if err != nil {
			taskDeletionErrors = append(taskDeletionErrors, err)
		}
	}

	if len(taskDeletionErrors) > 0 {
		return c.Status(500).JSON(fiber.Map{"error": "failed to delete tasks", "messages": taskDeletionErrors})
	}

	return c.Status(200).JSON(fiber.Map{
		"result": project,
	})
}
