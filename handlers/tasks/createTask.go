package handlers

import (
	"CRUD_API_FIBER/common"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type createDTO struct {
	Title        string    `json:"title" bson:"title"`
	Author       string    `json:"author" bson:"author"`
	CreationDate time.Time `json:"creationDate,omitempty" bson:"creationDate,omitempty"`
	Description  string    `json:"description" bson:"description"`
}

func CreateTask(c *fiber.Ctx) error {

	id := c.Params("id")

	if id == "" {
		return c.Status(400).JSON(fiber.Map{"error": "project id is required"})
	}

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid project id"})
	}

	task := new(createDTO)
	err = c.BodyParser(task)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}

	task.CreationDate = time.Now()
	coll := common.GetDbCollection("tasks")

	result, err := coll.InsertOne(c.Context(), task)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Falied to create a task",
			"message": err.Error(),
		})
	}

	coll = common.GetDbCollection("projects")

	filter := bson.M{"_id": objectId}

	update := bson.M{
		"$push": bson.M{"tasks": task},
	}

	_, err = coll.UpdateOne(c.Context(), filter, update)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to create a task",
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"result": fiber.Map{
			"insertedID": result.InsertedID,
			"task":       task,
		},
	})
}
