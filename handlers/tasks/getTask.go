package handlers

import (
	"CRUD_API_FIBER/common"
	"CRUD_API_FIBER/models"
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTask(c *fiber.Ctx) error {
	coll := common.GetDbCollection("tasks")

	id := c.Params("id")

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}

	filter := bson.M{"_id": objectId}

	task := models.Task{}

	err = coll.FindOne(context.TODO(), filter).Decode(&task)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"task": task})
}
