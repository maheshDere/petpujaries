package user

import (
	"context"
	"fmt"
	"petpujaris/downloader"
	"petpujaris/logger"
)

type FileHandler struct {
	Service UserFileService
}

func NewFileHandler(service UserFileService) *FileHandler {
	return &FileHandler{Service: service}
}

func (ufh *FileHandler) DownloadEmployeeFileData(ctx context.Context, req *downloader.EmployeeFileDownloadRequest) (*downloader.EmployeeFileDownloadResponse, error) {
	var employeedetails downloader.EmployeeFileDownloadResponse

	resp, err := ufh.Service.GetPrimaryUserDetails(context.TODO(), req.AdminID, req.TotalEmployeeCount)
	if err != nil {
		logger.LogError(err, "downloader.user.handler", fmt.Sprintf("fail to get users primary data for admin id : %v and total employee count : %v", req.AdminID, req.TotalEmployeeCount))
		return &downloader.EmployeeFileDownloadResponse{}, err
	}

	for _, ed := range resp {
		var employeedata downloader.EmployeeData
		employeedata.EmployeeData = ed
		employeedetails.EmployeeDetails = append(employeedetails.EmployeeDetails, &employeedata)
	}

	return &employeedetails, nil
}
