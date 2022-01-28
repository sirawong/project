package sendgrid

import (
	"log"
	"reservation/entities"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func (sendGrid *SendGrid) Send(data *entities.Email) (bool, error) {
	from := mail.NewEmail("Movie project", "movie-project@email.com")
	subject := data.Subject
	to := mail.NewEmail(data.To, data.ToEmail)
	plainTextContent := "Movie project"
	htmlContent := data.ContentHtml

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	apiKey := sendGrid.sendgridApiKey
	client := sendgrid.NewSendClient(apiKey)
	checkSend, err := client.Send(message)
	log.Println(checkSend)
	if err != nil {
		return false, err
	}
	return true, nil
}
