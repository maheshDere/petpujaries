package user

import (
	"context"
	"errors"
	"petpujaris/logger"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	logger.Setup()
}

func TestGetPrimaryUserDetails(t *testing.T) {
	t.Run("when Invalid admin id pass to method", func(t *testing.T) {
		resourceableID := uint64(0)
		ctx := context.Background()
		ur, ufs := setup()
		expectedError := errors.New("expected error")
		expectedUserPrimaryDetials := make([][]string, 0)
		ur.On("GetResourceableID", ctx, uint64(1)).Return(resourceableID, expectedError).Once()
		t.Run("It should return primary user details", func(t *testing.T) {
			userPrimaryDetials, err := ufs.GetPrimaryUserDetails(ctx, uint64(1), uint64(3))
			assert.Error(t, err)
			assert.Equal(t, expectedUserPrimaryDetials, userPrimaryDetials)
			ur.AssertExpectations(t)
		})
	})
	t.Run("when valid admin id pass to method", func(t *testing.T) {
		resourceableID := uint64(7)
		ctx := context.Background()
		ur, ufs := setup()
		expectedUserPrimaryDetials := [][]string{[]string{"name", "email", "mobile_number", "is_active", "role_id", "resourceable_id", "resourceable_type"}, []string{" ", " ", " ", "true", "1", "7", "Company"}, []string{" ", " ", " ", "true", "1", "7", "Company"}, []string{" ", " ", " ", "true", "1", "7", "Company"}}
		ur.On("GetResourceableID", ctx, uint64(1)).Return(resourceableID, nil).Once()
		t.Run("It should return primary user details", func(t *testing.T) {
			userPrimaryDetials, err := ufs.GetPrimaryUserDetails(ctx, uint64(1), uint64(3))
			assert.NoError(t, err)
			assert.Equal(t, expectedUserPrimaryDetials, userPrimaryDetials)
			ur.AssertExpectations(t)
		})
	})
}

func setup() (*MockUserRegistry, UserFileService) {
	ur := new(MockUserRegistry)
	return ur, NewUserFileService(ur)
}
