package email

import (
	"bytes"
	"context"
	"fmt"
	"mime/quotedprintable"
	"petpujaris/logger"
	"strings"
)

type emailService struct {
	From     string
	Password string
	Subject  string
	Client   EmailClient
}

func (em emailService) WriteEmail(ctx context.Context, dest []string, contentType, bodyMessage string) (string, error) {

	header := make(map[string]string)
	header["From"] = em.From

	receipient := ""

	for _, user := range dest {
		receipient = receipient + user
	}

	header["To"] = receipient
	header["Subject"] = em.Subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = fmt.Sprintf("%s; charset=\"utf-8\"", contentType)
	header["Content-Transfer-Encoding"] = "quoted-printable"
	header["Content-Disposition"] = "inline"

	message := ""

	for key, value := range header {
		message += fmt.Sprintf("%s: %s\r\n", key, value)
	}

	var encodedMessage bytes.Buffer

	finalMessage := quotedprintable.NewWriter(&encodedMessage)
	_, err := finalMessage.Write([]byte(bodyMessage))
	if err != nil {
		logger.LogError(err, "email.service.WriteEmail", fmt.Sprintf("error in write mail for user %v", dest))
		return "", err
	}
	finalMessage.Close()

	message += "\r\n" + encodedMessage.String()

	return message, nil
}

func (em emailService) WriteHTMLEmail(ctx context.Context, dest []string, bodyMessage string) (string, error) {
	return em.WriteEmail(ctx, dest, "text/html", bodyMessage)
}

func (em emailService) SendMail(ctx context.Context, dest []string, password string) (err error) {
	emailBody := CreateEmail(dest[0], password)
	bodyMessage, err := em.WriteHTMLEmail(ctx, dest, emailBody)
	if err != nil {
		logger.LogError(err, "email.service.Send mail", fmt.Sprintf("error in send mail for user %v", dest))
		return
	}

	msg := "From: " + em.From + "\n" +
		"To: " + strings.Join(dest, ",") + "\n" +
		"Subject: " + em.Subject + "\n" + bodyMessage

	err = em.Client.SendMail(ctx, em.From, em.Password, dest, []byte(msg))
	if err != nil {
		logger.LogError(err, "email.service.Send mail", fmt.Sprintf("error in send mail from client for user %v", dest))
		return
	}

	return nil
}

func NewEmailService(from, password, subject string, client EmailClient) EmailService {
	return emailService{
		From:     from,
		Password: password,
		Subject:  subject,
		Client:   client,
	}
}
