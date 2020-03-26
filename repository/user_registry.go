package repository

import (
	"context"
	"fmt"
	"petpujaris/logger"
	"petpujaris/models"
)

type userRegistry struct {
	client Client
}

const RollID = 3

func (ur userRegistry) Save(ctx context.Context, user models.User) (userID int64, err error) {
	rows, err := ur.client.Query(ctx, CreateUserQuery, user.Name, user.Email, user.MobileNumber, user.IsActive, user.Password, user.RoleID, user.ResourceableID, user.ResourceableType, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		logger.LogError(err, "repository.saveUser", fmt.Sprintf("save user : %v", user))
		return userID, err
	}

	defer rows.Close()
	if rows.Next() {
		if err := rows.Scan(&userID); err != nil {
			return userID, err
		}
	}

	return
}

func (ur userRegistry) GetResourceableID(ctx context.Context, ID uint64) (uint64, error) {
	var resourceableID uint64
	err := ur.client.QueryRow(ctx, GetResourceableIDQuery, ID, RollID).Scan(&resourceableID)
	if err != nil {
		return resourceableID, err
	}
	return resourceableID, nil
}

func (ur userRegistry) SaveProfile(ctx context.Context, profile models.Profile) (err error) {
	_, err = ur.client.Exec(ctx, CreateUserProfileQuery, profile.EmployeeID, profile.Credits, profile.NotificationsEnabled, profile.UserID, profile.MealTypeID, profile.CreatedAt, profile.UpdatedAt)
	if err != nil {
		logger.LogError(err, "repository.saveUserProfile", fmt.Sprintf("save user profile : %v", profile))
		return
	}

	return
}

func NewUserRegistry(pg Client) userRegistry {
	return userRegistry{client: pg}
}
