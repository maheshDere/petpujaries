package restaurant

import (
	"net/http"
	"net/http/httptest"
	"petpujaris/config"
	"petpujaris/logger"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const dataUplodationURL = "/petpujaris/restaurant/csv/upload"

func TestRestaurantCSVHandler(t *testing.T) {
	err := config.SetupConfig()
	assert.NoError(t, err)
	config.LoadConfig()
	logger.Setup()

	t.Run("When users pass invalid csv file", func(t *testing.T) {
		csvStr := "row1,value1,value2\nrow2,value1,value2,"
		req := httptest.NewRequest("POST", dataUplodationURL, strings.NewReader(csvStr))
		req.Header.Add("Content-Type", "application/json")
		responseRecorder := httptest.NewRecorder()
		handler := setupRestaurantCSVHandler()
		handler(responseRecorder, req)
		t.Run("it should return StatusBadRequest ", func(t *testing.T) {
			actualResponse := responseRecorder.Body.String()
			assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
			assert.Empty(t, actualResponse)
		})
	})

	t.Run("When users pass csv file", func(t *testing.T) {
		csvStr := "row1,value1,value2\nrow2,value1,value2"
		req := httptest.NewRequest("POST", dataUplodationURL, strings.NewReader(csvStr))
		req.Header.Add("Content-Type", "application/json")
		responseRecorder := httptest.NewRecorder()
		handler := setupRestaurantCSVHandler()
		handler(responseRecorder, req)
		t.Run("it should return statusOK", func(t *testing.T) {
			actualResponse := responseRecorder.Body.String()
			assert.Equal(t, http.StatusOK, responseRecorder.Code)
			assert.Empty(t, actualResponse)
		})
	})
}
func setupRestaurantCSVHandler() http.HandlerFunc {
	return RestaurantCSVHandler()
}
