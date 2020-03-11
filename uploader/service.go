package uploader

import (
	"context"
	"petpujaris/repository"
	"petpujaris/workers"
)

type uploaderService struct {
	MealRegistry repository.MealRegistry
}

func (rs uploaderService) SaveBulkdata(ctx context.Context, module string, data [][]string) error {
	p := workers.NewPool(10, len(data), data, rs.MealRegistry)
	p.Run(ctx)
	return nil
}

func NewUploaderService(mealRegistry repository.MealRegistry) UploaderService {
	return uploaderService{MealRegistry: mealRegistry}
}
