package main

import (
	"CRUD_API_FIBER/common"
	"CRUD_API_FIBER/router"
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	client := common.ConnectDb()

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	app := fiber.New()
	app.Use(logger.New())

	router.AddTaskGroup(app)
	router.AddProjectGroup(app)

	log.Fatal(app.Listen(":7070"))
}
