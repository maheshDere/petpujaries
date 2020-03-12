package server

import (
	"fmt"
	"log"
	"net"
	"petpujaris/config"
	"petpujaris/filemanager"
	"petpujaris/logger"
	"petpujaris/repository"
	"petpujaris/uploader"
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

func (gc GRPC) Start() error {
	dbconfig := config.GetDBConfig()
	pgClient, err := repository.NewPgClient(dbconfig)
	if err != nil {
		return err
	}

	mealRegistry := repository.NewMealsRegistry(pgClient)
	uploaderService := uploader.NewUploaderService(mealRegistry)
	fileService := filemanager.NewXLSXFileService()
	//fileService := filemanager.NewCSVFileService(true, ',', -1)
	ur := repository.NewUserRegistry(pgClient)
	uploaderService := uploader.NewUploaderService(ur)

	uploaderHandler := uploader.NewUploaderHandler(uploaderService, fileService)
	Server := grpc.NewServer()
	uploader.RegisterUploadServiceServer(Server, uploaderHandler)

	logger.LogInfo(logrus.Fields{"Port": gc.Port}, "Server started")

	l, err := net.Listen("tcp", gc.getGRPCPort())
	if err != nil {
		logger.LogError(err, "GRPC Start", fmt.Sprintf("Could not Listen to : %v", gc.getGRPCPort()))
		log.Fatalf("Could not Listen to : %v , %v", gc.getGRPCPort(), err)
	}

	return Server.Serve(l)
}
