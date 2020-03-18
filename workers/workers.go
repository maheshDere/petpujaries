package workers

import (
	"context"
	"fmt"
	"petpujaris/email"
	"petpujaris/logger"
	"petpujaris/models"
	"petpujaris/repository"
	"petpujaris/utils"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/araddon/dateparse"
)

var mealsSheetHeader = []string{"MealName", "Description", "ImageUrl", "Price", "Calories", "IsActive", "Item", "Ingredients", "meal_type_id", "MealType", "restaurant_cuisine_id", "CuisineName", "Errors"}
var schedulerSheetHeader = []string{"MealID", "MealName", "Date", "Errors"}
var employeeSheetHeader = []string{"name", "email", "mobile_number", "is_active", "role_id", "resourceable_id", "resourceable_type", "Errors"}

const EMP_SHEET_COLUMN_LENGTH = 7
const MEALS_SHEET_COLUMN_LENGTH = 12
const MEALS_SCHEDULER_COLUMN_LENGTH = 3

type Pool struct {
	Workers               int
	Emailservice          email.EmailService
	MealRegistry          repository.MealRegistry
	UserRepository        repository.UserRegistry
	MealSchedulerRegistry repository.MealSchedulerRegistry
}

type errorLog struct {
	Records []string
	MealID  int64
}

func NewPool(workers int, mealRegistry repository.MealRegistry, userRepository repository.UserRegistry, es email.EmailService, mealSchedulerRegistry repository.MealSchedulerRegistry) Pool {
	return Pool{Workers: workers, MealRegistry: mealRegistry, UserRepository: userRepository, Emailservice: es, MealSchedulerRegistry: mealSchedulerRegistry}
}

func (p Pool) Run(ctx context.Context, module string, userID int64, data [][]string) [][]string {
	tasks := make(chan []string, len(data))
	errorlog := make(chan errorLog, len(data))
	var errorRecords [][]string
	var wg sync.WaitGroup
	switch strings.ToLower(module) {
	case "meal":
		for w := 1; w <= p.Workers; w++ {
			go p.mealWorker(ctx, w, &wg, tasks, errorlog)
		}
		errorRecords = append(errorRecords, mealsSheetHeader)

	case "employee":
		for w := 1; w <= p.Workers; w++ {
			go p.UserWorker(ctx, &wg, tasks, errorlog)
		}
		errorRecords = append(errorRecords, employeeSheetHeader)

	case "mealscheduler":
		for w := 1; w <= p.Workers; w++ {
			go p.SchedulerWorker(ctx, &wg, &userID, tasks, errorlog)
		}
		errorRecords = append(errorRecords, mealsSheetHeader)
	}

	for t := 1; t < len(data); t++ {
		wg.Add(1)
		tasks <- data[t]
	}

	close(tasks)

	for t := 1; t < len(data); t++ {
		errs := <-errorlog
		if len(errs.Records) != 0 {
			errorRecords = append(errorRecords, errs.Records)
			if errs.MealID != 0 {
				p.MealRegistry.Delete(ctx, errs.MealID)
			}
		}
	}

	fmt.Println("errorRecords ", errorRecords)
	wg.Wait()
	return errorRecords
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
			errorRecord = append(errorRecord, strings.Join(errs[:], ","))
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

		itemSplit := strings.Split(t[6], ",")
		errData := p.createMealSubWorker(ctx, id, itemSplit)
		if len(errData) != 0 {
			errorRecord = append(errorRecord, t...)
			errorRecord = append(errorRecord, strings.Join(errData[:], ","))
			errorlog <- errorLog{Records: errorRecord, MealID: id}
			continue
		}

		ingredientSplit := strings.Split(t[7], ",")
		errData = p.createMealIngredientSubWorker(ctx, id, ingredientSplit)
		if len(errData) != 0 {
			errorRecord = append(errorRecord, t...)
			errorRecord = append(errorRecord, strings.Join(errData[:], ","))
			errorlog <- errorLog{Records: errorRecord, MealID: id}
			continue
		}

		errorlog <- errorLog{Records: errorRecord}

	}
}

func (p Pool) createMealSubWorker(ctx context.Context, mealID int64, itemSplit []string) []string {
	items := make(chan models.Items, len(itemSplit))
	errorlog := make(chan error, len(itemSplit))
	var wgSub sync.WaitGroup
	for i := 0; i < len(itemSplit); i++ {
		go p.mealItemSubWorker(ctx, &wgSub, items, errorlog)
	}

	for t := 0; t < len(itemSplit); t++ {
		wgSub.Add(1)
		mealItemRecord := models.Items{
			MealsID:   mealID,
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

func (p Pool) createMealIngredientSubWorker(ctx context.Context, mealID int64, ingredientsSplit []string) []string {
	ingredients := make(chan models.Ingredients, len(ingredientsSplit))
	errorlog := make(chan error, len(ingredientsSplit))
	var wgSub sync.WaitGroup
	for i := 0; i < len(ingredientsSplit); i++ {
		go p.mealIngredientSubWorker(ctx, &wgSub, &mealID, ingredients, errorlog)
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

func (p Pool) mealIngredientSubWorker(ctx context.Context, wgSub *sync.WaitGroup, mealID *int64, ingredients <-chan models.Ingredients, errorlog chan<- error) {
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
			MealsID:      *mealID,
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
	if len(t) != MEALS_SHEET_COLUMN_LENGTH {
		errs = append(errs, fmt.Sprint("Invalid sheet ,contain less information then expected"))
		return models.Meals{}, errs
	}

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
	var err error
	for task := range tasks {
		wg.Done()
		var errorRecord []string
		user, errs := parseUser(task)
		if len(errs) != 0 {
			errorRecord = append(errorRecord, task...)
			errorRecord = append(errorRecord, strings.Join(errs[:], ","))
			errorlog <- errorLog{Records: errorRecord}
			continue
		}

		err = user.Validate()
		if err != nil {
			errorRecord = append(errorRecord, task...)
			errorRecord = append(errorRecord, err.Error())
			errorlog <- errorLog{Records: errorRecord}
			continue
		}

		key := utils.RandomString()
		user.Password, err = user.GenerateHashedPassword(key)
		if err != nil {
			errorRecord = append(errorRecord, task...)
			errorRecord = append(errorRecord, err.Error())
			errorlog <- errorLog{Records: errorRecord}
			continue
		}

		err = p.UserRepository.Save(ctx, user)
		if err != nil {
			errorRecord = append(errorRecord, task...)
			errorRecord = append(errorRecord, err.Error())
			errorlog <- errorLog{Records: errorRecord}
			continue
		}

		err := p.Emailservice.SendMail(ctx, []string{user.Email}, key)
		if err != nil {
			logger.LogError(err, "worker", fmt.Sprintf("fail to send email for user %v", user.Name))
		}

		errorlog <- errorLog{}
	}
}

func parseUser(task []string) (user models.User, errs []string) {
	var err error

	if len(task) != EMP_SHEET_COLUMN_LENGTH {
		errs = append(errs, fmt.Sprint("Invalid sheet ,contain less information then expected"))
		return
	}

	user.Name = task[0]
	user.Email = task[1]
	user.MobileNumber = task[2]
	user.IsActive, err = strconv.ParseBool(task[3])
	if err != nil {
		errs = append(errs, fmt.Sprintf("can not parse IsActive  value %s", task[3]))
	}

	roleID, err := strconv.ParseInt(task[4], 10, 64)
	if err != nil {
		errs = append(errs, fmt.Sprintf("can not parse role ID  value %s", task[5]))
	}
	user.RoleID = int(roleID)

	resourceableID, err := strconv.ParseInt(task[5], 10, 64)
	if err != nil {
		errs = append(errs, fmt.Sprintf("can not parse resourceableID  value %s", task[6]))
	}
	user.ResourceableID = int(resourceableID)

	user.ResourceableType = task[6]
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	return
}

func (p Pool) SchedulerWorker(ctx context.Context, wg *sync.WaitGroup, userID *int64, tasks <-chan []string, errorlog chan<- errorLog) {
	var errorRecord []string
	for task := range tasks {
		wg.Done()
		scheduler, errs := parseScheduler(task)
		if len(errs) != 0 {
			errorRecord = append(errorRecord, task...)
			errorRecord = append(errorRecord, strings.Join(errs[:], ","))
			errorlog <- errorLog{Records: errorRecord}
			continue
		}
		scheduler.UserID = *userID
		if err := scheduler.Validation(); err != nil {
			errorRecord = append(errorRecord, task...)
			errorRecord = append(errorRecord, err.Error())
			errorlog <- errorLog{Records: errorRecord}
			continue
		}

		err := p.MealSchedulerRegistry.Save(ctx, scheduler)
		if err != nil {
			errorRecord = append(errorRecord, task...)
			errorRecord = append(errorRecord, err.Error())
			errorlog <- errorLog{Records: errorRecord}
			continue
		}
		errorlog <- errorLog{}
	}
}

func parseScheduler(task []string) (models.MealScheduler, []string) {
	var errs []string
	if len(task) != MEALS_SCHEDULER_COLUMN_LENGTH {
		errs = append(errs, fmt.Sprint("Invalid sheet ,contain less information then expected"))
		return models.MealScheduler{}, errs
	}

	mealID, err := strconv.ParseInt(task[0], 10, 32)
	if err != nil {
		errs = append(errs, fmt.Sprintf("can not parse Meal ID value %s", task[0]))
	}
	schedulerDate, err := dateparse.ParseLocal(task[2])
	if err != nil {
		errs = append(errs, fmt.Sprintf("invalid date %s", task[2]))
	}
	scheduler := models.MealScheduler{
		MealID:    mealID,
		Date:      schedulerDate,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return scheduler, errs
}
