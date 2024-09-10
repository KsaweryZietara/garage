package mail

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"path"
	"runtime"
)

const (
	headers             = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"
	NewEmployeeTemplate = "newEmployee.html"
)

type NewEmployee struct {
	GarageName string
	Code       string
}

type Config struct {
	Username string `env:"USERNAME"`
	Password string `env:"PASSWORD"`
	SmtpHost string `env:"SMTP_HOST"`
	SmtpPort string `env:"SMTP_PORT"`
}

type Mail struct {
	cfg Config
}

func New(cfg Config) *Mail {
	return &Mail{
		cfg: cfg,
	}
}

func (m *Mail) Send(to string, subject string, templateName string, templateData interface{}) error {
	_, currentPath, _, _ := runtime.Caller(0)
	templatePath := fmt.Sprintf("%s/resources/templates/%s", path.Join(path.Dir(currentPath), "../../../../"), templateName)

	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	if err = t.Execute(&body, templateData); err != nil {
		return err
	}

	auth := smtp.PlainAuth("", m.cfg.Username, m.cfg.Password, m.cfg.SmtpHost)

	msg := "Subject: " + subject + "\n" + headers + "\n\n" + body.String()

	if err = smtp.SendMail(m.cfg.SmtpHost+":"+m.cfg.SmtpPort, auth, m.cfg.Username, []string{to}, []byte(msg)); err != nil {
		return err
	}

	return nil
}
