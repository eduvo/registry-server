package main

import (
	"github.com/jordan-wright/email"
	"net/smtp"
)

func sendmail() error {
	e := email.NewEmail()
	e.From = "Test <from@example.com>"
	e.To = []string{"test@example.com"}
	e.Bcc = []string{"test_bcc@example.com"}
	e.Cc = []string{"test_cc@example.com"}
	e.Subject = "A Subject"
	e.Text = "Text Body in plain text."
	e.HTML = "<h1>Text Body in HTML.</h1>"
	return e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "test@gmail.com", "password123", "smtp.gmail.com"))
}
