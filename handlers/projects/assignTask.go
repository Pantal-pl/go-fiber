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
	projectId := c.Params("id")

	if projectId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "project id is required",
		})
	}

	objectID, err := primitive.ObjectIDFromHex(projectId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid project id",
		})
	}

	collTasks := common.GetDbCollection("tasks")
	collProjects := common.GetDbCollection("projects")

	projectFilter := bson.M{"_id": objectID}

	existingProject := collProjects.FindOne(c.Context(), projectFilter)
	if existingProject.Err() == mongo.ErrNoDocuments {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "document doesn't exist"})
	}

	var decodedExistingProject models.Project
	err = existingProject.Decode(&decodedExistingProject)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to decode existing project",
			"message": err.Error(),
		})
	}

	task := new(models.Task)
	if err := c.BodyParser(task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}
	task.CreationDate = time.Now()
	task.AssignedProject = decodedExistingProject.ID

	result, err := collTasks.InsertOne(c.Context(), task)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to assign a task",
			"message": err.Error(),
		})
	}

	taskInsertedID := result.InsertedID.(primitive.ObjectID).Hex()

	task.ID = taskInsertedID
	update := bson.M{"$push": bson.M{"tasks": task}}

	_, err = collProjects.UpdateOne(c.Context(), projectFilter, update)
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
