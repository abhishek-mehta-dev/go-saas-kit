package email

import (
	"fmt"
	"net/smtp"
	"os"
)

type EmailSender interface {
	Send(to, subject, body string) error
}

type smtpSender struct {
	auth smtp.Auth
	from string
	host string
	port string
}

func NewSMTPSender() EmailSender {
	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")
	user := os.Getenv("SMTP_USER")
	pass := os.Getenv("SMTP_PASS")

	auth := smtp.PlainAuth("", user, pass, host)

	return &smtpSender{
		auth: auth,
		from: user,
		host: host,
		port: port,
	}
}

func (s *smtpSender) Send(to, subject, body string) error {
	addr := fmt.Sprintf("%s:%s", s.host, s.port)
	msg := []byte(fmt.Sprintf(
		"To: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\nContent-Type: text/html; charset=\"UTF-8\"\r\n\r\n%s",
		to, subject, body,
	))
	return smtp.SendMail(addr, s.auth, s.from, []string{to}, msg)
}
