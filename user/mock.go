package user

import (
	"context"
	"petpujaris/models"

	"github.com/stretchr/testify/mock"
)

type MockRegistry struct{ mock.Mock }

func (e *MockRegistry) FindUserByID(ctx context.Context, userID string) (models.User, error) {
	args := e.Called(ctx, userID)
	return args.Get(0).(models.User), args.Error(1)
}
