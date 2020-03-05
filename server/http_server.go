package server

import (
	"net/http"
	"petpujaris/config"
	"petpujaris/filemanager"
	"petpujaris/logger"
	"petpujaris/repository"
	"petpujaris/restaurant"
	"petpujaris/user"
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
	dbconfig := config.GetDBConfig()
	pgClient, err := repository.NewPgClient(dbconfig)
	if err != nil {
		return err
	}

	dbRepository := repository.NewDBRegistry(pgClient)
	userService := user.NewUserService(dbRepository)
	FindUserByIDHandler := user.FindByID(userService)
	restaurantService := restaurant.NewRestaurantService()
	fileOperation := filemanager.NewXLSXFileService()
	//fileOperation := filemanager.NewCSVFileService(true, ',', -1)

	restaurantCSVHandler := restaurant.RestaurantCSVHandler(restaurantService, fileOperation)

	router := mux.NewRouter()
	restaurantRouter := router.PathPrefix("/petpujaris/restaurant").Subrouter()
	restaurantRouter.HandleFunc("/csv/upload", restaurantCSVHandler).Methods(http.MethodPost)

	userRouter := router.PathPrefix("/petpujaris").Subrouter()
	userRouter.HandleFunc("/user/{userID}", FindUserByIDHandler).Methods(http.MethodGet)

	logger.LogInfo(logrus.Fields{"Port": hs.Port}, "Server started")
	return http.ListenAndServe(hs.getPort(), router)
}
