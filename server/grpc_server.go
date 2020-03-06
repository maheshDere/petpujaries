package server

import (
	"fmt"
	"log"
	"net"
	"petpujaris/filemanager"
	"petpujaris/logger"
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

	uploaderService := uploader.NewUploaderService()
	fileService := filemanager.NewXLSXFileService()
	//fileService := filemanager.NewCSVFileService(true, ',', -1)

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
