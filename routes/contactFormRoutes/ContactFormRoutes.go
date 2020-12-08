package contactFormRoutes

import (
	"github.com/gofiber/fiber/v2"
	"gorm-fiberv2-go/model/contactForm"
)

func SetupContactFormRoutes(app *fiber.App) {

	app.Get("/contactForm", contactForm.GetAllContactForm)
	app.Post("/contactForm", contactForm.SaveContactForm)
	app.Post("/mail", contactForm.SendMailContactForm)

}
