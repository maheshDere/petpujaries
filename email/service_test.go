package email

import (
	"context"
	"errors"
	"petpujaris/logger"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func init() {
	logger.Setup()
}

func TestSendEmail(t *testing.T) {
	t.Run("when WriteEmail method returns an error", func(t *testing.T) {
		_, _, _, client, es := setup()
		expectedError := errors.New("expected error")
		client.On("SendMail", context.Background(), "from@mail.com", "password", []string{"destinationmail@mail"}, mock.Anything).Return(expectedError).Once()
		t.Run("SendEmail should return an error", func(t *testing.T) {
			err := es.SendMail(context.Background(), []string{"destinationmail@mail"}, "password")
			assert.Equal(t, expectedError, err)
		})
		client.AssertExpectations(t)
	})
	t.Run("when WriteEmail method returns an error", func(t *testing.T) {
		_, _, _, client, es := setup()
		client.On("SendMail", context.Background(), "from@mail.com", "password", []string{"destinationmail@mail"}, mock.Anything).Return(nil).Once()
		t.Run("SendEmail should return an error", func(t *testing.T) {
			err := es.SendMail(context.Background(), []string{"destinationmail@mail"}, "password")
			assert.NoError(t, err)
		})
		client.AssertExpectations(t)
	})

}

func setup() (string, string, string, *MockEmailClient, EmailService) {
	client := new(MockEmailClient)
	from := "from@mail.com"
	password := "password"
	subject := "subject"
	return from, password, subject, client, NewEmailService(from, password, subject, client)
}
