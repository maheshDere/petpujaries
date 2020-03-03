package user

import (
	"context"
	"errors"
	"petpujaris/logger"
	"petpujaris/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var registryMock *MockRegistry

func init() {
	logger.Setup()
}

func TestFindUserByID(t *testing.T) {
	setup()
	userID := "user_id"
	ctx := context.Background()

	t.Run("when user not found in repository", func(t *testing.T) {
		expectederr := errors.New("expected err")
		registryMock.On("FindUserByID", ctx, userID).Return(models.User{}, expectederr).Once()
		t.Run("repository should return error", func(t *testing.T) {
			user, err := registryMock.FindUserByID(ctx, userID)
			assert.Equal(t, models.User{}, user)
			assert.Equal(t, expectederr, err)
			registryMock.AssertExpectations(t)
		})
	})
	t.Run("when user found in repository", func(t *testing.T) {
		expectedUser := models.User{
			CountryCode:  "IND",
			MobileNumber: "1234567890",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}
		registryMock.On("FindUserByID", ctx, userID).Return(expectedUser, nil).Once()
		t.Run("repository should return user", func(t *testing.T) {
			user, err := registryMock.FindUserByID(ctx, userID)
			assert.NoError(t, err)
			assert.Equal(t, expectedUser, user)
		})
	})
}

func setup() {
	registryMock = new(MockRegistry)
}
