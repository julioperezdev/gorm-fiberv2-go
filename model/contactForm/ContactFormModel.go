package contactForm

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm-fiberv2-go/repository"
	"gorm.io/gorm"
	"log"
	"net/smtp"
)

type ContactForm struct {
	gorm.Model
	Name    string `json:"name"`
	Gender  string `json:"gender"`
	Content string `json:"content"`
}

func MigrateContactForm(sql *gorm.DB) {
	sql.AutoMigrate(&ContactForm{})
	fmt.Println("Contact Form Entity migrated")
}

func GetAllContactForm(ctx *fiber.Ctx) error {
	db := repository.ConnectMysql()
	MigrateContactForm(db)
	var contactForms []ContactForm
	db.Find(&contactForms)
	return ctx.Status(fiber.StatusOK).JSON(contactForms)
}

func SaveContactForm(ctx *fiber.Ctx) error {
	db := repository.ConnectMysql()
	MigrateContactForm(db)
	type request struct {
		Name    string `json:"name"`
		Gender  string `json:"gender"`
		Content string `json:"content"`
	}
	var body request
	err := ctx.BodyParser(&body)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse JSON",
		})
	}
	particularContactForm := ContactForm{
		Name:    body.Name,
		Gender:  body.Gender,
		Content: body.Content,
	}
	db.Create(&particularContactForm)
	return ctx.Status(fiber.StatusCreated).JSON(particularContactForm)

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
