package contactFormRoutes

import (
	"github.com/gofiber/fiber/v2"
	"gorm-fiberv2-go/model/contactForm"
)

func SetupContactFormRoutes(app *fiber.App) {
	app.Post("/mail", contactForm.SendMailContactForm)
}
