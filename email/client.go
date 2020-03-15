package email

import (
	"context"
	"net/smtp"
)

type emailClient struct{}

const SMTPServer = "smtp.gmail.com"

func (c *emailClient) SendMail(ctx context.Context, from, password string, dest []string, msg []byte) error {
	return smtp.SendMail(SMTPServer+":587",
		smtp.PlainAuth("", from, password, SMTPServer),
		from, dest, []byte(msg))
}

func NewEmailClient() EmailClient {
	return &emailClient{}
}
