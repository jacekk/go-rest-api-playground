package mailling

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/jordan-wright/email"
)

type MailMessage struct {
	Message    string   `validate:"required|minLen:10"`
	Subject    string   `validate:"required|maxLen:30"`
	Recipients []string `validate:"required"`
}

func getSendAddr() string {
	host := os.Getenv("MAIL_HOST")
	port := os.Getenv("MAIL_PORT")

	return fmt.Sprintf("%s:%s", host, port)
}

func getSenderAuth() smtp.Auth {
	user := os.Getenv("MAIL_USER")
	pass := os.Getenv("MAIL_PASS")
	host := os.Getenv("MAIL_HOST")

	return smtp.PlainAuth("", user, pass, host)
}

func SendEmail(msg MailMessage) error {
	e := email.NewEmail()

	e.From = os.Getenv("MAIL_USER")
	e.To = msg.Recipients // or `e.Bcc` or `e.Cc`
	e.Subject = msg.Subject
	e.Text = []byte(msg.Message)
	e.HTML = []byte(fmt.Sprintf("<h1>%s</h1>", msg.Message))

	err := e.Send(getSendAddr(), getSenderAuth())

	return err
}
