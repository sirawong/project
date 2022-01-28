package implement

import (
	"bytes"
	"context"
	"html/template"

	"reservation/entities"
	"reservation/errs"
	"reservation/logs"
	"reservation/service/invitation/input"
)

func (impl *implementation) Invitation(ctx context.Context, invitation []*input.InvitationInput) (err error) {

	var doc bytes.Buffer
	t, err := template.ParseFiles("html/email.html")
	if err != nil {
		logs.Error(err)
		return errs.NewUnexpectedError()
	}

	for _, val := range invitation {
		err := t.Execute(&doc, val)
		if err != nil {
			logs.Error(err)
			return errs.NewUnexpectedError()
		}
		impl.SendGridRepo.Send(&entities.Email{
			To:           val.To,
			ToEmail:      val.To,
			From:         "test@example.com",
			FromEmail:    "test@example.com",
			Subject:      "Invitation Movie",
			ContentPlain: doc.String(),
			ContentHtml:  doc.String(),
		})
	}

	return nil
}
