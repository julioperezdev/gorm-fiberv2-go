package contactForm

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
	"net/smtp"
)

type ContactForm struct {
	gorm.Model
	name    string `json:"name"`
	gender  string `json:"gender"`
	content string `json:"content"`
}

func MigrateContactForm(sql *gorm.DB) {
	sql.AutoMigrate(&ContactForm{})
	fmt.Println("Contact Form Entity migrated")
}

func BodyRequest(ctx *fiber.Ctx) ContactForm {
	type request struct {
		name    string `json:"name"`
		gender  string `json:"gender"`
		content string `json:"content"`
	}
	var body request
	err := ctx.BodyParser(&body)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse JSON",
		})
	}
	particularContactForm := ContactForm{
		name:    body.name,
		gender:  body.gender,
		content: body.content,
	}
	return particularContactForm

}

func SendMailContactForm(ctx *fiber.Ctx) error {
	auth := smtp.PlainAuth(
		"",
		"ceo@protobot.dev",
		"XXXXXXXXXXXXXXXXXXX",
		"smtpout.secureserver.net")

	to := []string{"perezjulioernesto@gmail.com"}
	msg := []byte("To: recipient@example.net\r\n" +
		"Subject: discount Gophers!\r\n" +
		"\r\n" +
		"This is the email body.\r\n")
	err := smtp.SendMail(
		"smtpout.secureserver.net:465",
		auth,
		"ceo@protobot.dev",
		to,
		msg)
	if err != nil {
		log.Fatal(err)
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"Successfully": "Mail was send",
	})
}
