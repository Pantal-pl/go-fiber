package handlers

import (
	"CRUD_API_FIBER/common"
	"CRUD_API_FIBER/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTask(c *fiber.Ctx) error {

	projectId := c.Params("id")

	if projectId == "" {
		return c.Status(400).JSON(fiber.Map{"error": "project id is required"})
	}

	convertedProjectId, err := primitive.ObjectIDFromHex(projectId)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid project id"})
	}

	task := new(models.Task)
	err = c.BodyParser(task)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}

	task.AssignedProject = projectId
	task.CreationDate = time.Now()
	collTasks := common.GetDbCollection("tasks")

	insertedTaskResult, err := collTasks.InsertOne(c.Context(), task)

	insertedID := insertedTaskResult.InsertedID.(primitive.ObjectID).Hex()
	task.ID = insertedID

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Falied to create a task",
			"message": err.Error(),
		})
	}

	collProjects := common.GetDbCollection("projects")

	filter := bson.M{"_id": convertedProjectId}
	update := bson.M{
		"$push": bson.M{"tasks": task},
	}

	_, err = collProjects.UpdateOne(c.Context(), filter, update)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to create a task",
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"result": fiber.Map{
			"insertedID": insertedTaskResult.InsertedID,
			"task":       task,
		},
	})
}
