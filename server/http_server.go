package server

import (
	"net/http"
	"petpujaris/filemanager"
	"petpujaris/logger"
	"petpujaris/restaurant"

	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type HTTP struct {
	Port int
}

func (hs HTTP) getPort() string {
	return ":" + strconv.Itoa(hs.Port)
}

func (hs HTTP) Start() error {
	restaurantService := restaurant.NewRestaurantService()
	fileOperation := filemanager.NewXLSXFileService()
	//fileOperation := filemanager.NewCSVFileService(true, ',', -1)

	restaurantCSVHandler := restaurant.RestaurantCSVHandler(restaurantService, fileOperation)

	router := mux.NewRouter()
	restaurantRouter := router.PathPrefix("/petpujaris/restaurant").Subrouter()
	restaurantRouter.HandleFunc("/csv/upload", restaurantCSVHandler).Methods(http.MethodPost)

	logger.LogInfo(logrus.Fields{"Port": hs.Port}, "Server started")
	return http.ListenAndServe(hs.getPort(), router)
}
