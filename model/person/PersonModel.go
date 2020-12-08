package person

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm-fiberv2-go/repository"
	"gorm.io/gorm"
)

const PreRequiredAge = 18

type Person struct {
	gorm.Model
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func MigratePerson(sql *gorm.DB) {
	sql.AutoMigrate(&Person{})
	fmt.Println("Person Entity migrated")
}

func GetAllPerson(ctx *fiber.Ctx) error {
	db := repository.ConnectMysql()
	MigratePerson(db)
	var persons []Person
	db.Find(&persons)
	return ctx.Status(fiber.StatusOK).JSON(persons)
}

func PostPerson(ctx *fiber.Ctx) error {
	db := repository.ConnectMysql()
	MigratePerson(db)
	type request struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	var body request
	err := ctx.BodyParser(&body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	particularPerson := Person{
		Name: body.Name,
		Age:  body.Age,
	}

	db.Create(&particularPerson)
	return ctx.Status(fiber.StatusCreated).JSON(particularPerson)

}

func PostPerson2(ctx *fiber.Ctx) Person {
	type request struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	var body request
	err := ctx.BodyParser(&body)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse JSON",
		})
	}
	particularPerson := Person{
		Name: body.Name,
		Age:  body.Age,
	}
	return particularPerson

}

func EvaluateAge(ctx *fiber.Ctx) error {

	particularPerson := PostPerson2(ctx)
	if particularPerson.Age < PreRequiredAge {
		return ctx.Status(fiber.StatusPreconditionFailed).JSON(fiber.Map{
			"Precondition failed": "You need have 18 o more years to be saved under Database",
		})
	}
	db := repository.ConnectMysql()
	MigratePerson(db)
	db.Create(&particularPerson)
	return ctx.Status(fiber.StatusOK).JSON(particularPerson)

}
