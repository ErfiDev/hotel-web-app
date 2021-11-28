package main

import (
	"fmt"
	"github.com/erfidev/hotel-web-app/models"
	"net/smtp"
)

func Listener() {
	go func() {
		for {
			mail := <- appConfig.MailChan
			Send(mail)
		}
	}()
}

func Send(msg models.MailData) {
	// Sender data.
	from := msg.From
	password := msg.Pass

	// Receiver email address.
	to := []string{
		msg.To,
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte(msg.Content)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}