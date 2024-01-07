package handlers

import (
	"CRUD_API_FIBER/common"
	"CRUD_API_FIBER/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AssignTask(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "id is required",
		})
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	collTasks := common.GetDbCollection("tasks")
	collProjects := common.GetDbCollection("projects")

	existingProject := collProjects.FindOne(c.Context(), bson.M{"_id": objectID})
	if existingProject.Err() == mongo.ErrNoDocuments {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "document doesn't exist"})
	}

	task := new(models.Task)
	if err := c.BodyParser(task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}
	task.CreationDate = time.Now()

	result, err := collTasks.InsertOne(c.Context(), task)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to assign a task",
			"message": err.Error(),
		})
	}
	insertedID := result.InsertedID.(primitive.ObjectID).Hex()

	filter := bson.M{"_id": objectID}
	task.ID = insertedID
	update := bson.M{"$push": bson.M{"tasks": task}}

	_, err = collProjects.UpdateOne(c.Context(), filter, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to assign a task to the project",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Task assigned successfully",
	})
}
