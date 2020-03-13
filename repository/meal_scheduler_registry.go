package repository

import (
	"context"
	"petpujaris/models"
)

type mealSchedulerRegistry struct {
	client Client
}

func (ms mealSchedulerRegistry) Save(ctx context.Context, schedulerRecord models.MealScheduler) error {
	_, err := ms.client.Exec(ctx, SaveMealSchedulerQuery, schedulerRecord.Date, schedulerRecord.MealID,
		schedulerRecord.UserID, schedulerRecord.CreatedAt, schedulerRecord.UpdatedAt)

	return err
}

func NewMealSchedulerRegistry(pg Client) MealSchedulerRegistry {
	return mealSchedulerRegistry{client: pg}
}
