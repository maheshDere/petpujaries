package restaurant

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"petpujaris/logger"
)

func RestaurantCSVHandler() http.HandlerFunc {
	restaurantCSVHandler := func(rw http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		csvData, err := parseCSVFile(r)
		if err != nil {
			logger.LogError(err, "RestaurantCSVHandler", "CSV file can not parse")
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		fmt.Println(csvData)
		rw.WriteHeader(http.StatusOK)

	}
	return http.HandlerFunc(restaurantCSVHandler)
}

func parseCSVFile(req *http.Request) ([][]string, error) {
	// parse POST body as csv
	reader := csv.NewReader(req.Body)
	var results [][]string
	for {
		// read one row from csv
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		// add record to result set
		results = append(results, record)
	}
	return results, nil
}
