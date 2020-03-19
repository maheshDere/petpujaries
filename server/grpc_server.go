package server

import (
	"fmt"
	"log"
	"net"
	"petpujaris/config"
	"petpujaris/downloader"
	"petpujaris/downloader/meals"
	"petpujaris/downloader/user"
	"petpujaris/email"
	"petpujaris/filemanager"
	"petpujaris/logger"
	"petpujaris/repository"
	"petpujaris/uploader"
	"petpujaris/workers"
	"strconv"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type GRPC struct {
	Port int
}

func (gc GRPC) getGRPCPort() string {
	return ":" + strconv.Itoa(gc.Port)
}

const WORKERS = 10

func (gc GRPC) Start() error {
	dbconfig := config.GetDBConfig()
	pgClient, err := repository.NewPgClient(dbconfig)
	if err != nil {
		return err
	}

	fileService := filemanager.NewCSVFileService(true, ',', -1)
	//fileService := filemanager.NewXLSXFileService()

	mealRegistry := repository.NewMealsRegistry(pgClient)
	userRegistry := repository.NewUserRegistry(pgClient)
	mealScheduler := repository.NewMealSchedulerRegistry(pgClient)
	emailConfig := config.GetEmailConfig()
	emailClient := email.NewEmailClient()
	emailService := email.NewEmailService(emailConfig.Email, emailConfig.Password, email.EmailSubject, emailClient)
	workerPool := workers.NewPool(WORKERS, mealRegistry, userRegistry, emailService, mealScheduler)

	uploaderService := uploader.NewUploaderService(workerPool)
	uploaderHandler := uploader.NewUploaderHandler(uploaderService, fileService)

	downloaderService := user.NewUserFileService(userRegistry)

	downloaderMealsService := meals.NewMealsFileService(mealRegistry)
	fileHandler := downloader.NewFileHandler(downloaderService, downloaderMealsService)

	Server := grpc.NewServer()

	uploader.RegisterUploadServiceServer(Server, uploaderHandler)

	downloader.RegisterDownloadServiceServer(Server, fileHandler)

	logger.LogInfo(logrus.Fields{"Port": gc.Port}, "Server started")

	l, err := net.Listen("tcp", gc.getGRPCPort())
	if err != nil {
		logger.LogError(err, "GRPC Start", fmt.Sprintf("Could not Listen to : %v", gc.getGRPCPort()))
		log.Fatalf("Could not Listen to : %v , %v", gc.getGRPCPort(), err)
	}

	return Server.Serve(l)
}
