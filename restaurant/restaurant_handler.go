package restaurant

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"petpujaris/logger"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func RestaurantCSVHandler(restaurantService RestaurantService) http.HandlerFunc {
	restaurantCSVHandler := func(rw http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		csvData, err := ParseXLSXFile(r)
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

func ParseXLSXFile(req *http.Request) ([][]string, error) {
	var results [][]string
	file, _, err := req.FormFile("demo")
	if err != nil {
		logger.LogError(err, "ParseXLSXFile", "error to get file from request")
		return nil, err
	}

	f, err := excelize.OpenReader(file)
	if err != nil {
		logger.LogError(err, "ParseXLSXFile", "error to get file pointer")
		return nil, err
	}

	sheetName := f.GetSheetName(1)
	data, err := f.GetRows(sheetName)
	if err != nil {
		logger.LogError(err, "ParseXLSXFile", "error to get file data")
		return nil, err
	}

	for key, rows := range data {
		fmt.Printf("key : %v row : %v\n", key, rows)
	}
	return results, nil
}
