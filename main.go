package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm-fiberv2-go/routes/todoRoutes"
)

func main() {

	//Setting Fiber v2 to start server
	app := fiber.New()
	app.Use(logger.New())

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Welcome\n")
	})
	//Setting to routes TodoModel functionalities
	todoRoutes.SetupTodoRoutes(app)

	//Setting to listen server port
	app.Listen(":3000")
}
