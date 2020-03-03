package server

import (
	"net/http"
	"petpujaris/config"
	"petpujaris/repository"
	"petpujaris/user"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type HTTP struct {
	Port int
}

type Server interface {
	Start() error
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

	router := mux.NewRouter()
	joshRouter := router.PathPrefix("/josh/petpujarires").Subrouter()

	joshRouter.HandleFunc("/user/{userID}", FindUserByIDHandler).Methods(http.MethodGet)

	server := negroni.Classic()
	server.UseHandler(router)

	server.Run(hs.getPort())
	return nil
}
