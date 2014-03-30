package main

import (
    "bytes"
    "fmt"
    "net/smtp"
    "strings"
    "text/template"
)

func SendEmail(host string, port int, userName string, password string, to []string, subject string, message string) (err error) {

    parameters := &struct {
        From string
        To string
        Subject string
        Message string
    }{
        userName,
        strings.Join([]string(to), ","),
        subject,
        message,
    }

    buffer := new(bytes.Buffer)

    template := template.Must(template.New("emailTemplate").Parse(_EmailScript()))
    template.Execute(buffer, parameters)

    auth := smtp.PlainAuth("", userName, password, host)

    err = smtp.SendMail(
        fmt.Sprintf("%s:%d", host, port),
        auth,
        userName,
        to,
        buffer.Bytes())

    return err
}

// _EmailScript returns a template for the email message to be sent
func _EmailScript() (script string) {

    return `From: {{.From}}
            To: {{.To}}
            Subject: {{.Subject}}
            MIME-version: 1.0
            Content-Type: text/html; charset="UTF-8"

            {{.Message}}`
}