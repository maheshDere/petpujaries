package restaurant

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"petpujaris/config"
	"petpujaris/logger"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const dataUplodationURL = "/petpujaris/restaurant/csv/upload"

func TestRestaurantCSVHandler(t *testing.T) {
	err := config.SetupConfig()
	assert.NoError(t, err)
	config.LoadConfig()
	logger.Setup()

	t.Run("When users pass invalid csv file", func(t *testing.T) {
		b, _ := generateCSVData(t)

		req := httptest.NewRequest("POST", dataUplodationURL, &b)
		responseRecorder := httptest.NewRecorder()
		mockrestaurantService, handler := setupRestaurantCSVHandler()
		mockrestaurantService.On("SaveBulkRestaurantData", mock.Anything).Return(nil)

		handler(responseRecorder, req)
		t.Run("it should return StatusBadRequest ", func(t *testing.T) {
			actualResponse := responseRecorder.Body.String()
			assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
			assert.Empty(t, actualResponse)
			mockrestaurantService.AssertNotCalled(t, "SaveBulkRestaurantData")
		})
	})

	t.Run("When users pass csv file", func(t *testing.T) {
		b, w := generateCSVData(t)
		req := httptest.NewRequest("POST", dataUplodationURL, &b)
		req.Header.Set("Content-Type", w.FormDataContentType())
		responseRecorder := httptest.NewRecorder()
		mockrestaurantService, handler := setupRestaurantCSVHandler()
		mockrestaurantService.On("SaveBulkRestaurantData", mock.Anything).Return(nil)

		handler(responseRecorder, req)
		t.Run("it should return statusOK", func(t *testing.T) {
			actualResponse := responseRecorder.Body.String()
			assert.Equal(t, http.StatusOK, responseRecorder.Code)
			assert.Empty(t, actualResponse)
			mockrestaurantService.AssertExpectations(t)
		})
	})
}
func setupRestaurantCSVHandler() (*MockRestaurantService, http.HandlerFunc) {
	mockrestaurantService := new(MockRestaurantService)
	return mockrestaurantService, RestaurantCSVHandler(mockrestaurantService)
}

func generateCSVData(t *testing.T) (bytes.Buffer, *multipart.Writer) {
	var b bytes.Buffer
	var fw io.Writer
	w := multipart.NewWriter(&b)
	csvData := strings.NewReader("row1,value1,value2\nrow2,value1,value2")
	fw, err := w.CreateFormFile("csvfile", "testcsv.csv")
	assert.NoError(t, err)
	_, err = io.Copy(fw, csvData)
	assert.NoError(t, err)
	w.Close()
	return b, w
}
