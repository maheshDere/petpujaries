package user

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func FindByID(Service Service) http.HandlerFunc {
	findByIDHandler := func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userID := vars["userID"]
		_, err := uuid.Parse(userID)
		if userID == "" || err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		user, err := Service.FindUserByID(r.Context(), userID)
		if err != nil {
			rw.WriteHeader(http.StatusNotFound)
			return
		}

		jsonResponse, err := json.Marshal(user)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.WriteHeader(http.StatusOK)
		rw.Write(jsonResponse)

	}
	return http.HandlerFunc(findByIDHandler)
}
