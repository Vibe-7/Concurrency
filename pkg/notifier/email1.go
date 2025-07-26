package notifier

import (
	"fmt"
	"net/smtp"
)

type EmailNotifier struct {
	From     string
	To       string
	Password string
	SMTPHost string
	SMTPPort string
}

func NewEmailNotifier(from, to, password, smtpHost, smtpPort string) *EmailNotifier {
	return &EmailNotifier{
		From:     from,
		To:       to,
		Password: password,
		SMTPHost: smtpHost,
		SMTPPort: smtpPort,
	}
}

func (n *EmailNotifier) SendNotification(msg string) {
	auth := smtp.PlainAuth("", n.From, n.Password, n.SMTPHost)
	body := fmt.Sprintf("To: %s\r\nSubject: Notification\r\n\r\n%s", n.To, msg)
	err := smtp.SendMail(n.SMTPHost+":"+n.SMTPPort, auth, n.From, []string{n.To}, []byte(body))
	if err != nil {
		fmt.Printf("Failed to send email: %v\n", err)
	}
}
