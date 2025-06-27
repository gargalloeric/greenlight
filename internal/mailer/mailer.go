package mailer

import (
	"bytes"
	"embed"
	"html/template"
	"log"
	"time"

	"github.com/wneessen/go-mail"
)

//go:embed "templates"
var templateFS embed.FS

type Mailer struct {
	client *mail.Client
	sender string
}

func New(host string, port int, username, password, sender string) Mailer {
	client, err := mail.NewClient(
		host,
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithPort(port),
		mail.WithTimeout(5*time.Second),
		mail.WithTLSPolicy(mail.TLSOpportunistic),
		mail.WithUsername(username),
		mail.WithPassword(password),
	)
	if err != nil {
		log.Fatalf("failed to create new mail delivery client: %s\n", err)
	}

	return Mailer{
		client: client,
		sender: sender,
	}
}

func (m Mailer) Send(recipient, templateFile string, data any) error {
	tmpl, err := template.New("email").ParseFS(templateFS, "templates/"+templateFile)
	if err != nil {
		return err
	}

	subject := new(bytes.Buffer)
	if err := tmpl.ExecuteTemplate(subject, "subject", data); err != nil {
		return err
	}

	plainBody := new(bytes.Buffer)
	if err := tmpl.ExecuteTemplate(plainBody, "plainBody", data); err != nil {
		return err
	}

	htmlBody := new(bytes.Buffer)
	if err := tmpl.ExecuteTemplate(htmlBody, "htmlBody", data); err != nil {
		return err
	}

	message := mail.NewMsg()
	if err := message.To(recipient); err != nil {
		return err
	}

	if err := message.From(m.sender); err != nil {
		return err
	}
	message.Subject(subject.String())
	message.SetBodyString(mail.TypeTextPlain, plainBody.String())
	message.AddAlternativeString(mail.TypeTextHTML, htmlBody.String())

	if err := m.client.DialAndSend(message); err != nil {
		return err
	}

	return nil
}
