package mailer

import (
	"bytes"
	"crypto/tls"
	"github.com/go-mail/mail"
	"github.com/spf13/viper"
	"html/template"
)

type Mail struct {
	From     string
	To       string
	Cc       string
	Subject  string
	Template string
	Data     interface{}
}

func (m Mail) SendMail() error {

	var temp = template.Must(template.ParseFiles("./public/templates/mailer/" + m.Template))
	var body bytes.Buffer

	err := temp.ExecuteTemplate(&body, m.Template, &m.Data)
	if err != nil {
		return err
	}

	msg := mail.NewMessage()
	msg.SetHeader("From", m.From)
	msg.SetHeader("To", m.To)
	msg.SetHeader("Subject", m.Subject)
	msg.SetBody("text/html", body.String())

	mailer := mail.NewDialer(viper.GetString("smtp.host"), viper.GetInt("smtp.port"), viper.GetString("smtp.username"), viper.GetString("smtp.password_generator"))
	mailer.TLSConfig = &tls.Config{}

	if err = mailer.DialAndSend(msg); err != nil {
		return err
	}
	return nil
}
