package todoRoutes

import (
	"github.com/gofiber/fiber/v2"
	"gorm-fiberv2-go/model/todo"
)

func SetupTodoRoutes(app *fiber.App) {
	app.Get("/todo", todo.GetTodo)
	app.Post("/todo", todo.PostTodo)
	//app.Delete("/todo/:id", todo.DeleteTodo)
	//app.Patch("/todo/:id", todo.PatchTodo)

}
