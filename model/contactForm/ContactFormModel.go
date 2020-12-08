package contactForm

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm-fiberv2-go/repository"
	"gorm.io/gorm"
	"html/template"
	"log"
	"net/mail"
	"net/smtp"
)

type ContactForm struct {
	gorm.Model
	Name    string `json:"name"`
	Gender  string `json:"gender"`
	Content string `json:"content"`
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
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

func SendMail2(ctx *fiber.Ctx) error {

	type Dest struct {
		Name string
	}
	from := mail.Address{"Naturismo", "ceo@protobot.dev"}
	to := mail.Address{"Julio", "perezjulioernesto@gmail.com"}
	subject := "Enviado desde Go"
	dest := Dest{Name: to.Address}

	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subject
	headers["Content-Type"] = `text/html; charset="UTF-8"`

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s : %s\r\n", k, v)
	}
	t, err := template.ParseFiles("/home/protobot/go/src/gorm-fiberv2-go/templates/mail.html")
	checkError(err)

	buf := new(bytes.Buffer)
	err = t.Execute(buf, dest)
	checkError(err)

	message += buf.String()

	servername := "smtpout.secureserver.net:465"
	host := "smtpout.secureserver.net"

	auth := smtp.PlainAuth("", "ceo@protobot.dev", "9513451julio", host)
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}
	conn, err := tls.Dial("tcp", servername, tlsConfig)
	checkError(err)

	client, err := smtp.NewClient(conn, host)
	checkError(err)

	err = client.Auth(auth)
	checkError(err)

	err = client.Mail(from.Address)
	checkError(err)

	err = client.Rcpt(to.Address)
	checkError(err)

	w, err := client.Data()
	checkError(err)

	_, err = w.Write([]byte(message))
	checkError(err)

	err = w.Close()
	checkError(err)

	client.Quit()
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"successfully": "Mail was send it",
	})

}
