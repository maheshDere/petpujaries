package workers

import (
	"context"
	"fmt"
	"petpujaris/models"
	"petpujaris/repository"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Pool struct {
	Workers      int
	TotalRecord  int
	Records      [][]string
	MealRegistry repository.MealRegistry
}

type errorLog struct {
	Records []string
}

func NewPool(workers int, totalRecord int, records [][]string, mealRegistry repository.MealRegistry) Pool {
	return Pool{Workers: workers, TotalRecord: totalRecord, Records: records, MealRegistry: mealRegistry}
}

func (p Pool) Run(ctx context.Context) {
	tasks := make(chan []string, p.TotalRecord)
	errorlog := make(chan errorLog, 1)
	var wg sync.WaitGroup
	for w := 1; w <= p.Workers; w++ {
		go p.mealWorker(ctx, w, &wg, tasks, errorlog)
	}

	for t := 1; t < p.TotalRecord; t++ {
		wg.Add(1)
		tasks <- p.Records[t]
	}

	close(tasks)
	var errorRecords [][]string
	for t := 1; t < p.TotalRecord; t++ {
		errs := <-errorlog
		if len(errs.Records) != 0 {
			errorRecords = append(errorRecords, errs.Records)
		}
	}

	fmt.Println("errorRecords ", errorRecords)

	wg.Wait()

}

func (p Pool) mealWorker(ctx context.Context, wid int, wg *sync.WaitGroup, tasks <-chan []string, errorlog chan<- errorLog) {
	for t := range tasks {
		var errorRecord []string
		mealRecord, errs := parseMealRecord(t)
		if err := mealRecord.Validation(); err != nil {
			errs = append(errs, err.Error())
		}
		if len(errs) != 0 {
			errorRecord = append(errorRecord, t...)
			errorRecord = append(errorRecord, errs...)
			errorlog <- errorLog{Records: errorRecord}
			continue
		}

		id, err := p.MealRegistry.Save(ctx, mealRecord)
		if err != nil {
			errorRecord = append(errorRecord, t...)
			errorRecord = append(errorRecord, err.Error())
			errorlog <- errorLog{Records: errorRecord}
			continue
		}
		fmt.Println(id)
		itemSplit := strings.Split(t[6], ",")
		errData := p.createMealSubWorker(ctx, id, itemSplit)
		if len(errData) != 0 {
			errorRecord = append(errorRecord, t...)
			errorRecord = append(errorRecord, errData...)
			errorlog <- errorLog{Records: errorRecord}
			continue
		}

		errorlog <- errorLog{Records: errorRecord}
		wg.Done()
	}
}

func (p Pool) createMealSubWorker(ctx context.Context, mealId int64, itemSplit []string) []string {
	items := make(chan models.Items, len(itemSplit))
	errorlog := make(chan error, len(itemSplit))
	var wgSub sync.WaitGroup
	for i := 0; i < len(itemSplit); i++ {
		go p.mealItemSubWorker(ctx, &wgSub, items, errorlog)
	}

	for t := 0; t < len(itemSplit); t++ {
		wgSub.Add(1)
		mealItemRecord := models.Items{
			MealsID:   int8(mealId),
			Name:      itemSplit[t],
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		fmt.Println("mealItemRecord", mealId, mealItemRecord)
		items <- mealItemRecord
	}

	close(items)

	var errorRecords []string
	for t := 1; t < len(itemSplit); t++ {
		errs := <-errorlog
		if errs != nil {
			errorRecords = append(errorRecords, errs.Error())
		}
	}
	wgSub.Wait()
	return errorRecords
}

func (p Pool) mealItemSubWorker(ctx context.Context, wgSub *sync.WaitGroup, items <-chan models.Items, errorlog chan<- error) {
	for item := range items {
		err := p.MealRegistry.SaveItem(ctx, item)
		errorlog <- err
	}
	wgSub.Done()
}

func parseMealRecord(t []string) (models.Meals, []string) {
	var errs []string
	price, err := strconv.ParseFloat(t[3], 32)
	if err != nil {
		errs = append(errs, fmt.Sprintf("can not parse Price value %s", t[3]))
	}

	calories, err := strconv.ParseFloat(t[4], 32)
	if err != nil {
		errs = append(errs, fmt.Sprintf("can not parse Calories value %s", t[4]))
	}

	isActive, err := strconv.ParseBool(t[5])
	if err != nil {
		errs = append(errs, fmt.Sprintf("can not parse ISActive value %s", t[5]))
	}

	mealTypeID, err := strconv.ParseInt(t[8], 10, 32)
	if err != nil {
		errs = append(errs, fmt.Sprintf("can not parse Meal Type ID value %s", t[8]))
	}

	restaurantCuisineID, err := strconv.ParseInt(t[10], 10, 32)
	if err != nil {
		errs = append(errs, fmt.Sprintf("can not parse Restaurant Cuisine ID value %s", t[10]))

	}

	mealReord := models.Meals{
		Name:                t[0],
		Description:         t[1],
		ImageURL:            t[2],
		Price:               float32(price),
		Calories:            float32(calories),
		ISActive:            isActive,
		MealTypeID:          int8(mealTypeID),
		RestaurantCuisineID: int8(restaurantCuisineID),
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
	}
	return mealReord, errs
}
