package user

import (
	"context"
	"petpujaris/models"

	"github.com/stretchr/testify/mock"
)

type MockUserRegistry struct{ mock.Mock }

func (m *MockUserRegistry) GetResourceableID(ctx context.Context, ID uint64) (uint64, error) {
	args := m.Called(ctx, ID)
	return args.Get(0).(uint64), args.Error(1)
}

func (m *MockUserRegistry) Save(ctx context.Context, user models.User) (err error) {
	args := m.Called(ctx, user)
	return args.Error(0)
}
