package todo

import (
	"github.com/gofiber/fiber/v2"
	"gorm-fiberv2-go/repository"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	//ID        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

//
//var todos = []Todo{
//	{ID: 1, Name: "julio", Completed: false},
//	{ID: 2, Name: "maria", Completed: true},
//}

func GetTodo(ctx *fiber.Ctx) error {
	sql := repository.ConnectMysql()
	var todosDB []Todo
	sql.Find(&todosDB)
	return ctx.Status(fiber.StatusOK).JSON(todosDB)
}

func PostTodo(ctx *fiber.Ctx) error {
	db := repository.ConnectMysql()
	var todoDB Todo
	//type request struct {
	//	Name      string `json:"name"`
	//	Completed bool   `json:"completed"`
	//}
	//var body request
	err := ctx.BodyParser(&todoDB)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	//particularTodo := Todo{
	//	ID:        len(todos) + 1,
	//	Name:      body.Name,
	//	Completed: body.Completed,
	//}
	//todos = append(todos, particularTodo)
	db.Create(&todoDB)
	return ctx.Status(fiber.StatusCreated).JSON(todoDB)

}

/*
func DeleteTodo(ctx *fiber.Ctx) error {
	paramID := ctx.Params("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Id",
		})
	}
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[0:i], todos[i+i:]...)
			return ctx.Status(fiber.StatusOK).JSON(todo)
		}
	}
	return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "Record not found",
	})
}

func PatchTodo(ctx *fiber.Ctx) error {

	type request struct {
		Name      string `json:"name"`
		Completed bool   `json:"completed"`
	}
	var body request
	err := ctx.BodyParser(&body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parser JSON",
		})
	}

	paramID := ctx.Params("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Id",
		})
	}
	for i, todo := range todos {
		if todo.ID == id {
			todos[i] = Todo{
				ID:        id,
				Name:      body.Name,
				Completed: body.Completed,
			}
			return ctx.Status(fiber.StatusOK).JSON(todos[i])
		}
	}
	return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "Record not found",
	})
}

*/
