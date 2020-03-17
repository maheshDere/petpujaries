package email

import "context"

type EmailService interface {
	WriteEmail(ctx context.Context, dest []string, contentType, bodyMessage string) (string, error)
	WriteHTMLEmail(ctx context.Context, dest []string, bodyMessage string) (string, error)
	SendMail(ctx context.Context, dest []string, password string) (err error)
}

type EmailClient interface {
	SendMail(ctx context.Context, from, password string, dest []string, msg []byte) error
}
