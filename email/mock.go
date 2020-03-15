package email

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockEmailClient struct{ mock.Mock }

func (m *MockEmailClient) SendMail(ctx context.Context, from, password string, dest []string, msg []byte) error {
	args := m.Called(ctx, from, password, dest, msg)
	return args.Error(0)
}
