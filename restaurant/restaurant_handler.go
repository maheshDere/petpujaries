package restaurant

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"petpujaris/logger"
)

func RestaurantCSVHandler(restaurantService RestaurantService) http.HandlerFunc {
	restaurantCSVHandler := func(rw http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		csvData, err := parseCSVFile(r)
		if err != nil {
			logger.LogError(err, "RestaurantCSVHandler", "CSV file can not parse")
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		fmt.Println(csvData)
		restaurantService.SaveBulkRestaurantData(csvData)
		rw.WriteHeader(http.StatusOK)

	}
	return http.HandlerFunc(restaurantCSVHandler)
}

func parseCSVFile(req *http.Request) ([][]string, error) {
	file, _, err := req.FormFile("csvfile")
	if err != nil {
		return nil, err
	}
	byteData, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(bytes.NewReader(byteData))
	reader.LazyQuotes = true
	reader.Comma = ','
	reader.FieldsPerRecord = -1
	var results [][]string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		results = append(results, record)
	}
	return results, nil
}
