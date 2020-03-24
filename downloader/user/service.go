package user

import (
	"context"
	"fmt"
	"petpujaris/logger"
	"petpujaris/repository"
)

type userFileService struct {
	UserRegistry repository.UserRegistry
}

type UserFileService interface {
	GetPrimaryUserDetails(ctx context.Context, adminID, totalEmployeeCount uint64) ([][]string, error)
}

const USER_ROLL_ID = "1"
const USER_RESOURCEABLE_TYPE = "Company"

func (ufs userFileService) GetPrimaryUserDetails(ctx context.Context, adminID, totalEmployeeCount uint64) ([][]string, error) {
	primaryUserDetails := make([][]string, 0)
	_, err := ufs.UserRegistry.GetResourceableID(ctx, adminID)
	if err != nil {
		logger.LogError(err, "downloader.user", fmt.Sprintf("ResourceableID id not found for user with id %v", adminID))
		return primaryUserDetails, err
	}

	primaryUserDetails = createPrimaryUserData(ctx, totalEmployeeCount)
	return primaryUserDetails, nil
}

func createPrimaryUserData(ctx context.Context, totalEmployeeCount uint64) (primaryUserDetails [][]string) {
	primaryUserDetails = make([][]string, 0)

	headers := []string{"name", "email", "mobile_number"}
	primaryUserDetails = append(primaryUserDetails, headers)

	for empCount := uint64(0); empCount < totalEmployeeCount; empCount++ {
		user := []string{" ", " ", " "}
		primaryUserDetails = append(primaryUserDetails, user)
	}

	return
}

func NewUserFileService(ur repository.UserRegistry) UserFileService {
	return userFileService{ur}
}
