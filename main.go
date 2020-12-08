package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm-fiberv2-go/routes/contactFormRoutes"
	"gorm-fiberv2-go/routes/personRoutes"
	"gorm-fiberv2-go/routes/todoRoutes"
)

func main() {

	//Setting Fiber v2 to start server
	app := fiber.New()
	app.Use(logger.New())

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Welcome\n")
	})
	//Setting to routes Models functionalities
	todoRoutes.SetupTodoRoutes(app)
	personRoutes.SetupPersonRoutes(app)
	contactFormRoutes.SetupContactFormRoutes(app)

	//Setting to listen server port
	app.Listen(":3000")
}
