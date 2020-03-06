package server

import (
	"fmt"
	"log"
	"net"
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

	/* restaurantService := restaurant.NewRestaurantService()
	fileOperation := filemanager.NewXLSXFileService() */
	//fileOperation := filemanager.NewCSVFileService(true, ',', -1)
	uploaderHandler := uploader.NewUploaderHandler()
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
