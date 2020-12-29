package quark

import (
	"errors"
	"fmt"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendEmail(toName, toEmail, subject, plainTextContent, htmlContent string) error {
	from := mail.NewEmail(os.Getenv("EMAIL_FROM_NAME"), os.Getenv("EMAIL_FROM"))

	to := mail.NewEmail(toName, toEmail)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))

	response, err := client.Send(message)
	if err != nil {
		return err
	}

	if response.StatusCode != 202 {
		return errors.New(fmt.Sprintf("%d", response.StatusCode) + " " + response.Body)
	}

	return nil
}

func SendEmailWithReplyTo(replyToName, replyToEmail, toName, toEmail, subject, plainTextContent, htmlContent string) error {
	from := mail.NewEmail(os.Getenv("EMAIL_FROM_NAME"), os.Getenv("EMAIL_FROM"))
	replyTo := mail.NewEmail(replyToName, replyToEmail)
	to := mail.NewEmail(toName, toEmail)

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	message.SetReplyTo(replyTo)

	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))

	response, err := client.Send(message)
	if err != nil {
		return err
	}

	if response.StatusCode != 202 {
		return errors.New(fmt.Sprintf("%d", response.StatusCode) + " " + response.Body)
	}

	return nil
}
