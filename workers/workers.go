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
	Workers        int
	TotalRecord    int
	Records        [][]string
	MealRegistry   repository.MealRegistry
	UserRepository repository.UserRepository
}

type errorLog struct {
	Records []string
}

func NewPool(workers int, totalRecord int, records [][]string, mealRegistry repository.MealRegistry) Pool {
	return Pool{Workers: workers, TotalRecord: totalRecord, Records: records, MealRegistry: mealRegistry}
}

func (p Pool) Run(ctx context.Context) {
	tasks := make(chan []string, p.TotalRecord)
	errorlog := make(chan errorLog, p.TotalRecord)
	
	var wg sync.WaitGroup
	switch module {
	case "meal":
		for w := 1; w <= p.Workers; w++ {
			go p.mealWorker(ctx, w, &wg, tasks, errorlog)
		}
	
	case "employee":
		for w := 1; w <= p.Workers; w++ {
			go p.UserWorker(ctx, &wg, tasks, errorlog)
		}
	
	for t := 1; t < p.TotalRecord; t++ {
		wg.Add(1)
		tasks <- data[t]
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
		wg.Done()
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

		ingredientSplit := strings.Split(t[7], ",")
		errData = p.createMealIngredientSubWorker(ctx, id, ingredientSplit)
		if len(errData) != 0 {
			errorRecord = append(errorRecord, t...)
			errorRecord = append(errorRecord, errData...)
			errorlog <- errorLog{Records: errorRecord}
			continue
		}

		errorlog <- errorLog{Records: errorRecord}

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
			MealsID:   mealId,
			Name:      itemSplit[t],
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		items <- mealItemRecord
	}

	close(items)

	var errorRecords []string
	for t := 0; t < len(itemSplit); t++ {
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
		wgSub.Done()
		if err := item.Validation(); err != nil {
			errorlog <- err
			continue
		}
		err := p.MealRegistry.SaveItem(ctx, item)
		errorlog <- err
	}

}

func (p Pool) createMealIngredientSubWorker(ctx context.Context, mealId int64, ingredientsSplit []string) []string {
	ingredients := make(chan models.Ingredients, len(ingredientsSplit))
	errorlog := make(chan error, len(ingredientsSplit))
	var wgSub sync.WaitGroup
	for i := 0; i < len(ingredientsSplit); i++ {
		go p.mealIngredientSubWorker(ctx, &wgSub, &mealId, ingredients, errorlog)
	}

	for t := 0; t < len(ingredientsSplit); t++ {
		wgSub.Add(1)
		mealIngredientsRecord := models.Ingredients{
			Name:      ingredientsSplit[t],
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		ingredients <- mealIngredientsRecord
	}

	close(ingredients)

	var errorRecords []string
	for t := 0; t < len(ingredientsSplit); t++ {
		errs := <-errorlog
		if errs != nil {
			errorRecords = append(errorRecords, errs.Error())
		}
	}
	wgSub.Wait()
	return errorRecords
}

func (p Pool) mealIngredientSubWorker(ctx context.Context, wgSub *sync.WaitGroup, mealId *int64, ingredients <-chan models.Ingredients, errorlog chan<- error) {
	for ingredient := range ingredients {
		wgSub.Done()
		if err := ingredient.Validation(); err != nil {
			errorlog <- err
			continue
		}

		id, err := p.MealRegistry.SaveIngredients(ctx, ingredient)
		if err != nil {
			errorlog <- err
			continue
		}

		mealsIngredients := models.MealsIngredients{
			IngredientID: id,
			MealsID:      *mealId,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}

		if err := mealsIngredients.Validation(); err != nil {
			errorlog <- err
			continue
		}

		err = p.MealRegistry.SaveMealIngredients(ctx, mealsIngredients)
		if err != nil {
			errorlog <- err
			continue
		}

		errorlog <- err
	}

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
		MealTypeID:          mealTypeID,
		RestaurantCuisineID: restaurantCuisineID,
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
	}
	return mealReord, errs
}

func (p Pool) UserWorker(ctx context.Context, wg *sync.WaitGroup, tasks <-chan []string, errorlog chan<- errorLog) {
	var errorRecord []string
	for task := range tasks {
		wg.Done()
		user, errs := parseUser(task)
		if len(errs) != 0 {
			errorRecord = append(errorRecord, task...)
			errorRecord = append(errorRecord, errs...)
			errorlog <- errorLog{Records: errorRecord}
			continue
		}

		err := p.UserRepository.Save(ctx, user)
		if err != nil {
			errorRecord = append(errorRecord, task...)
			errorRecord = append(errorRecord, err.Error())
			errorlog <- errorLog{Records: errorRecord}
			continue
		}
		errorlog <- errorLog{}
	}
}

func parseUser(task []string) (user models.User, errs []string) {
	var err error
	user.Name = task[0]
	user.Email = task[1]
	user.MobileNumber = task[2]
	user.IsActive, err = strconv.ParseBool(task[3])
	if err != nil {
		errs = append(errs, fmt.Sprintf("can not parse IsActive  value %s", task[3]))
	}

	roleID, err := strconv.ParseFloat(task[5], 64)
	if err != nil {
		errs = append(errs, fmt.Sprintf("can not parse role ID  value %s", task[5]))
	}
	user.RoleID = int(roleID)

	resourceableID, err := strconv.ParseFloat(task[6], 64)
	if err != nil {
		errs = append(errs, fmt.Sprintf("can not parse resourceableID  value %s", task[6]))
	}
	user.ResourceableID = int(resourceableID)

	user.ResourceableType = task[7]
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	if len(errs) != 0 {
		return user, errs
	}

	user.Password, err = user.GenerateHashedPassword()
	if err != nil {
		errs = append(errs, fmt.Sprintf("fail to generate password  for value %s", task[6]))
	}

	return
}
