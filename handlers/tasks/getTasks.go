package handlers

import (
	"CRUD_API_FIBER/common"
	"CRUD_API_FIBER/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetTasks(c *fiber.Ctx) error {
	coll := common.GetDbCollection("tasks")
	tasks := make([]models.Task, 0)

	cursor, err := coll.Find(c.Context(), bson.M{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	defer cursor.Close(c.Context()) // Remember to close the cursor when done

	// Iterate over the cursor and decode results
	for cursor.Next(c.Context()) {
		task := models.Task{}
		if err := cursor.Decode(&task); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		tasks = append(tasks, task)
	}

	if err := cursor.Err(); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"tasks": tasks,
	})
}
