package todo

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm-fiberv2-go/repository"
	"gorm.io/gorm"
	"strconv"
)

type Todo struct {
	gorm.Model
	//ID        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

func MigrateTodo(sql *gorm.DB) {
	sql.AutoMigrate(&Todo{})
	fmt.Println("Todo Entity migrated")
}

func GetTodo(ctx *fiber.Ctx) error {
	db := repository.ConnectMysql()
	var todosDB []Todo
	MigrateTodo(db)
	db.Find(&todosDB)
	return ctx.Status(fiber.StatusOK).JSON(todosDB)
}

func GetTodoById(ctx *fiber.Ctx) error {
	db := repository.ConnectMysql()
	MigrateTodo(db)
	var todo Todo
	paramID := ctx.Params("id")
	idTodo, err := strconv.Atoi(paramID)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Id",
		})
	}
	db.First(&todo, idTodo)
	return ctx.Status(fiber.StatusOK).JSON(todo)

}

func PostTodo(ctx *fiber.Ctx) error {
	db := repository.ConnectMysql()
	//var todoDB Todo
	MigrateTodo(db)
	type request struct {
		Name      string `json:"name"`
		Completed bool   `json:"completed"`
	}
	var body request
	err := ctx.BodyParser(&body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	particularTodo := Todo{
		Name:      body.Name,
		Completed: body.Completed,
	}

	db.Create(&particularTodo)
	return ctx.Status(fiber.StatusCreated).JSON(particularTodo)

}

func DeleteTodo(ctx *fiber.Ctx) error {
	db := repository.ConnectMysql()
	var todo Todo
	MigrateTodo(db)
	paramID := ctx.Params("id")
	idTodo, err := strconv.Atoi(paramID)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Id",
		})
	}
	db.Delete(&todo, idTodo)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"Successfully": "Todo was deleted",
	})
}

func PatchTodo(ctx *fiber.Ctx) error {
	db := repository.ConnectMysql()
	var todo Todo
	MigrateTodo(db)
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
	idTodo, err := strconv.Atoi(paramID)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Id",
		})
	}
	db.First(&todo, idTodo)
	todo.Name = body.Name
	todo.Completed = body.Completed
	db.Save(&todo)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"Successfully": "Todo was updated",
	})
}
