package downloader

import (
	"context"
	"fmt"
	"petpujaris/downloader/meals"
	"petpujaris/downloader/user"
	"petpujaris/logger"
)

type FileHandler struct {
	userService user.UserFileService
	mealService meals.MealsFileService
}

func NewFileHandler(service user.UserFileService, mealService meals.MealsFileService) *FileHandler {
	return &FileHandler{userService: service, mealService: mealService}
}

func (ufh *FileHandler) DownloadEmployeeFileData(ctx context.Context, req *EmployeeFileDownloadRequest) (*EmployeeFileDownloadResponse, error) {
	var employeedetails EmployeeFileDownloadResponse

	resp, err := ufh.userService.GetPrimaryUserDetails(context.TODO(), req.AdminID, req.TotalEmployeeCount)
	if err != nil {
		logger.LogError(err, "downloader.user.handler", fmt.Sprintf("fail to get users primary data for admin id : %v and total employee count : %v", req.AdminID, req.TotalEmployeeCount))
		return &EmployeeFileDownloadResponse{}, err
	}

	for _, ed := range resp {
		var employeedata EmployeeData
		employeedata.EmployeeData = ed
		employeedetails.EmployeeDetails = append(employeedetails.EmployeeDetails, &employeedata)
	}

	return &employeedetails, nil
}

func (ufh *FileHandler) DownloadMealFileData(ctx context.Context, req *MealFileDownloadRequest) (*MealFileDownloadResponse, error) {
	var mealsResponse MealFileDownloadResponse

	resp, err := ufh.mealService.GetMealsDetails(context.TODO(), req.GetRestaurantID())
	if err != nil {
		logger.LogError(err, "downloader.meals.handler", fmt.Sprintf("fail to get restaurant meal data restaurantID : %v ", req.GetRestaurantID()))
		return &MealFileDownloadResponse{}, err
	}

	for _, md := range resp {
		var mealsData MealData
		mealsData.MealData = md
		mealsResponse.MealDetails = append(mealsResponse.MealDetails, &mealsData)
	}

	return &mealsResponse, nil
}

func (ufh *FileHandler) DownloadMealSchedulerFileData(ctx context.Context, req *MealSchedulerFileDownloadRequest) (*MealSchedulerFileDownloadResponse, error) {
	var mealsSchedulerResponse MealSchedulerFileDownloadResponse

	resp, err := ufh.mealService.GetMealsSchedulerDetails(context.TODO(), req.GetRestaurantID())
	if err != nil {
		logger.LogError(err, "downloader.meals.handler", fmt.Sprintf("fail to get restaurant meal data restaurantID : %v ", req.GetRestaurantID()))
		return &MealSchedulerFileDownloadResponse{}, err
	}

	for _, md := range resp {
		var mealSchedulerData MealSchedulerData
		mealSchedulerData.SchedulerData = md
		mealsSchedulerResponse.SchedulerDetails = append(mealsSchedulerResponse.SchedulerDetails, &mealSchedulerData)
	}

	return &mealsSchedulerResponse, nil
}
