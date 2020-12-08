package personRoutes

import (
	"github.com/gofiber/fiber/v2"
	"gorm-fiberv2-go/model/person"
)

func SetupPersonRoutes(app *fiber.App) {
	app.Get("/person", person.GetAllPerson)
	app.Post("/person", person.PostPerson)
	app.Post("/person/evaluate", person.EvaluateAge)
}
