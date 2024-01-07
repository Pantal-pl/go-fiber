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

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	coll := common.GetDbCollection("projects")

	project := models.Project{}

	err = coll.FindOneAndDelete(c.Context(), bson.M{"_id": objectId}).Decode(&project)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "falied to delete a project",
			"message": err.Error(),
		})
	}

	for _, v := range project.Tasks {
		coll = common.GetDbCollection("tasks")
		objectId, err = primitive.ObjectIDFromHex(v.ID)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "falied to delete project", "message": err.Error()})
		}
		_, err = coll.DeleteOne(c.Context(), bson.M{"_id": objectId})

		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "falied to delete project", "message": err.Error()})
		}
	}

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "falied to delete a project",
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"result": project,
	})
}
