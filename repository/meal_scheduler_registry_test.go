package repository

import (
	"context"
	"fmt"
	"petpujaris/config"
	"petpujaris/logger"
	"petpujaris/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var schedulerRegistry MealSchedulerRegistry

func init() {
	config.SetupConfig()
	logger.Setup()
	config.LoadConfig()
	dbConfig := config.GetDBConfig()
	fmt.Println(dbConfig)
	pgClient, err := NewPgClient(dbConfig)
	if err != nil {
		panic(err)
	}

	schedulerRegistry = NewMealSchedulerRegistry(pgClient)
}
func Test_mealSchedulerRegistry_Save(t *testing.T) {
	t.Run("when insert valid meals scheduler", func(t *testing.T) {
		schedulerRecord := models.MealScheduler{
			Date:      time.Now(),
			MealID:    1,
			UserID:    2,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		err := schedulerRegistry.Save(context.Background(), schedulerRecord)
		t.Run("it should not return an error", func(t *testing.T) {
			assert.NoError(t, err)
		})
	})
}
