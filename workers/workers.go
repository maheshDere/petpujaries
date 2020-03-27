package workers

import (
	"context"
	"database/sql"
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

var mealsSheetHeader = []string{"MealName", "Description", "ImageUrl", "Price", "Calories", "IsActive", "Item", "Ingredients", "Meal Type Id", "MealType", "Restaurant Cuisine Id", "CuisineName", "Errors"}
var schedulerSheetHeader = []string{"MealID", "MealName", "Date", "Errors"}
var employeeSheetHeader = []string{"name", "email", "mobile_number", "employee_id", "meal_type_id", "meal_type", "Errors"}

const EMP_SHEET_COLUMN_LENGTH = 6
const MEALS_SHEET_COLUMN_LENGTH = 12
const MEALS_SCHEDULER_COLUMN_LENGTH = 3
const DEFAULT_USER_STATUS = true
const EMPLOYEE_RESOURCEABLE_TYPE = "Company"
const EMPLOYEE_ROLL_ID = 1
const DEFAULT_NOTIFICATION_ENABLED = true
const DEFAULT_CREDITS = 0

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
		errorRecords = append(errorRecords, employeeSheetHeader)
		for w := 1; w <= p.Workers; w++ {
			resourceableID, err := p.UserRepository.GetResourceableID(ctx, uint64(userID))
			if err != nil {
				if err == sql.ErrNoRows {
					errorRecords = append(errorRecords, []string{" ", " ", " ", " ", " ", " ", "unauthorised user to upload employee details"})
				} else {
					errorRecords = append(errorRecords, []string{" ", " ", " ", " ", " ", " ", "Something went wrong please try again"})
				}
				return errorRecords
			}
			go p.UserWorker(ctx, &wg, tasks, errorlog, resourceableID)
		}

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
		if errMsgs := mealRecord.Validation(); len(errMsgs) != 0 {
			errs = append(errs, errMsgs...)
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
	errorlogs := make(chan []string, len(itemSplit))
	var wgSub sync.WaitGroup
	for i := 0; i < len(itemSplit); i++ {
		go p.mealItemSubWorker(ctx, &wgSub, items, errorlogs)
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
		errsMsgs := <-errorlogs
		if len(errsMsgs) != 0 {
			errorRecords = append(errorRecords, errsMsgs...)
		}
	}
	wgSub.Wait()
	return errorRecords
}

func (p Pool) mealItemSubWorker(ctx context.Context, wgSub *sync.WaitGroup, items <-chan models.Items, errorlog chan<- []string) {
	for item := range items {
		wgSub.Done()
		if errMsgs := item.Validation(); len(errMsgs) != 0 {
			errorlog <- errMsgs
			continue
		}
		err := p.MealRegistry.SaveItem(ctx, item)
		if err != nil {
			errorlog <- []string{err.Error()}
		}
		errorlog <- []string{}
	}
}

func (p Pool) createMealIngredientSubWorker(ctx context.Context, mealID int64, ingredientsSplit []string) []string {
	ingredients := make(chan models.Ingredients, len(ingredientsSplit))
	errorlog := make(chan []string, len(ingredientsSplit))
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
		if len(errs) != 0 {
			errorRecords = append(errorRecords, errs...)
		}
	}
	wgSub.Wait()
	return errorRecords

}

func (p Pool) mealIngredientSubWorker(ctx context.Context, wgSub *sync.WaitGroup, mealID *int64, ingredients <-chan models.Ingredients, errorlog chan<- []string) {
	for ingredient := range ingredients {
		wgSub.Done()
		if errMsgs := ingredient.Validation(); len(errMsgs) != 0 {
			errorlog <- errMsgs
			continue
		}

		id, err := p.MealRegistry.SaveIngredients(ctx, ingredient)
		if err != nil {
			errorlog <- []string{err.Error()}
			continue
		}

		mealsIngredients := models.MealsIngredients{
			IngredientID: id,
			MealsID:      *mealID,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}

		if errMsgs := mealsIngredients.Validation(); len(errMsgs) != 0 {
			errorlog <- errMsgs
			continue
		}

		err = p.MealRegistry.SaveMealIngredients(ctx, mealsIngredients)
		if err != nil {
			errorlog <- []string{err.Error()}
			continue
		}

		errorlog <- []string{}
	}

}

func parseMealRecord(t []string) (models.Meals, []string) {
	var errs []string
	if len(t) != MEALS_SHEET_COLUMN_LENGTH {
		errs = append(errs, fmt.Sprint("Invalid sheet ,contain unexpected information"))
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

func (p Pool) UserWorker(ctx context.Context, wg *sync.WaitGroup, tasks <-chan []string, errorlog chan<- errorLog, resourceableID uint64) {
	var err error
	for task := range tasks {
		wg.Done()
		var errorRecord []string
		user, errs := parseUser(task, resourceableID)
		if len(errs) != 0 {
			errorRecord = append(errorRecord, task...)
			errorRecord = append(errorRecord, strings.Join(errs[:], ","))
			errorlog <- errorLog{Records: errorRecord}
			continue
		}

		errMsgs := user.Validate()
		if len(errMsgs) != 0 {
			errorRecord = append(errorRecord, task...)
			errorRecord = append(errorRecord, strings.Join(errMsgs[:], ","))
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

		userID, err := p.UserRepository.Save(ctx, user)
		if err != nil {
			errorRecord = append(errorRecord, task...)
			errorRecord = append(errorRecord, err.Error())
			errorlog <- errorLog{Records: errorRecord}
			continue
		}

		user.Profile.UserID = userID
		errMsg := user.Profile.Validate()
		if len(errMsg) != 0 {
			errorRecord = append(errorRecord, task...)
			errorRecord = append(errorRecord, strings.Join(errMsg[:], ","))
			errorlog <- errorLog{Records: errorRecord}
			continue
		}

		err = p.UserRepository.SaveProfile(ctx, user.Profile)
		if err != nil {
			p.UserRepository.Delete(ctx, userID)
			errorRecord = append(errorRecord, task...)
			errorRecord = append(errorRecord, err.Error())
			errorlog <- errorLog{Records: errorRecord}
			continue
		}

		go func() {
			err = p.Emailservice.SendMail(ctx, []string{user.Email}, key)
			if err != nil {
				logger.LogError(err, "worker", fmt.Sprintf("fail to send email for user %v", user.Name))
			}
		}()

		errorlog <- errorLog{}
	}
}

func parseUser(task []string, resourceableID uint64) (user models.User, errs []string) {
	if len(task) != EMP_SHEET_COLUMN_LENGTH {
		errs = append(errs, fmt.Sprint("Invalid sheet ,contain less information then expected"))
		return
	}

	time := time.Now()
	user.Name = task[0]
	user.Email = task[1]
	user.MobileNumber = task[2]
	user.Profile.EmployeeID = task[3]
	mealTypeID, err := strconv.ParseInt(task[4], 10, 64)
	if err != nil {
		errs = append(errs, fmt.Sprintf("can not parse Meal type ID value %s", task[4]))
	}

	user.Profile.MealTypeID = mealTypeID
	user.IsActive = DEFAULT_USER_STATUS
	user.RoleID = EMPLOYEE_ROLL_ID

	user.ResourceableID = resourceableID

	user.ResourceableType = EMPLOYEE_RESOURCEABLE_TYPE
	user.CreatedAt = time
	user.UpdatedAt = time
	user.Profile.Credits = DEFAULT_CREDITS
	user.Profile.NotificationsEnabled = DEFAULT_NOTIFICATION_ENABLED
	user.Profile.CreatedAt = time
	user.Profile.UpdatedAt = time
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
		if errMsgs := scheduler.Validation(); len(errMsgs) != 0 {
			errorRecord = append(errorRecord, task...)
			errorRecord = append(errorRecord, strings.Join(errMsgs[:], ","))
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
