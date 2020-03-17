package user

import (
	"context"
	"fmt"
	"petpujaris/logger"
	"petpujaris/repository"
	"strconv"
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
	ResourceableID, err := ufs.UserRegistry.GetResourceableID(ctx, adminID)
	if err != nil {
		logger.LogError(err, "downloader.user", fmt.Sprintf("ResourceableID id not found for user with id %v", adminID))
		return primaryUserDetails, err
	}

	primaryUserDetails = createPrimaryUserData(ctx, ResourceableID, totalEmployeeCount)
	return primaryUserDetails, nil
}

func createPrimaryUserData(ctx context.Context, ResourceableID, totalEmployeeCount uint64) (primaryUserDetails [][]string) {
	strResourceableID := strconv.FormatUint(ResourceableID, 10)
	primaryUserDetails = make([][]string, 0)

	headers := []string{"name", "email", "mobile_number", "is_active", "role_id", "resourceable_id", "resourceable_type"}
	primaryUserDetails = append(primaryUserDetails, headers)

	for empCount := uint64(0); empCount < totalEmployeeCount; empCount++ {
		user := []string{" ", " ", " ", "true", USER_ROLL_ID, strResourceableID, USER_RESOURCEABLE_TYPE}
		primaryUserDetails = append(primaryUserDetails, user)
	}

	return
}

func NewUserFileService(ur repository.UserRegistry) UserFileService {
	return userFileService{ur}
}
