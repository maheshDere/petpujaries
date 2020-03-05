package restaurant

import (
	"fmt"
	"net/http"
	"petpujaris/filemanager"
	"petpujaris/logger"
)

func RestaurantCSVHandler(restaurantService RestaurantService, fileOperation filemanager.FileOperation) http.HandlerFunc {
	restaurantCSVHandler := func(rw http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		file, _, err := r.FormFile("csvfile")
		if err != nil {
			logger.LogError(err, "RestaurantCSVHandler", "invalid parameter")
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		csvData, err := fileOperation.Reader(file)
		if err != nil {
			logger.LogError(err, "RestaurantCSVHandler", "file can not read")
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		fmt.Println(csvData)
		restaurantService.SaveBulkRestaurantData(csvData)
		rw.WriteHeader(http.StatusOK)

	}
	return http.HandlerFunc(restaurantCSVHandler)
}
