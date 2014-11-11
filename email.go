package main

import (
	"bytes"
	"fmt"
	"net/smtp"
	"strings"
	"text/template"
)

const emailTemplate = `From: {{.From}}
To: {{.To}}
Subject: {{.Subject}}
MIME-version: 1.0
Content-Type: text/plain; charset="UTF-8"

This email is sent by journalCheck.

System Events
=-=-=-=-=-=-=

{{.Body}}`

func SendEmail(host string, port int, userName string, password string, to []string, subject string, message string) (err error) {

	parameters := &struct {
		From    string
		To      string
		Subject string
		Body    string
	}{
		userName,
		strings.Join([]string(to), ","),
		subject,
		message,
	}

	buffer := new(bytes.Buffer)

	t, err := template.New("emailTemplate").Parse(emailTemplate)
	if err != nil {

		return
	}
	err = t.Execute(buffer, parameters)
	if err != nil {

		return
	}

	auth := smtp.PlainAuth("", userName, password, host)

	err = smtp.SendMail(
		fmt.Sprintf("%s:%d", host, port),
		auth,
		userName,
		to,
		buffer.Bytes())

	return
}
