package filemanager

import (
	"errors"
	"os"
	"petpujaris/config"
	"petpujaris/logger"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXLSXFile_Reader(t *testing.T) {
	err := config.SetupConfig()
	assert.NoError(t, err)
	config.LoadConfig()
	logger.Setup()

	t.Run("when pass invalid file data", func(t *testing.T) {
		xlsxFile := XLSXFile{}
		csvStr := "row1,value1,value2\nrow2,value1,value2"
		expectedError := errors.New("zip: not a valid zip file")
		result, err := xlsxFile.Reader(strings.NewReader(csvStr))
		t.Run("it should return an error ", func(t *testing.T) {
			assert.Equal(t, expectedError, err)
			assert.Empty(t, result)
		})
	})
	t.Run("when pass valid data sheet", func(t *testing.T) {
		xlsxFile := XLSXFile{}
		file, err := os.Open("../test/Book.xlsx")
		assert.NoError(t, err)
		defer file.Close()
		result, err := xlsxFile.Reader(file)
		t.Run("it should return xlsx data ", func(t *testing.T) {
			assert.NoError(t, err)
			assert.NotEmpty(t, result)
		})
	})
}

func TestCSVFile_Reader(t *testing.T) {
	t.Run("When user pass invalid CSV data", func(t *testing.T) {
		cf := CSVFile{
			Delimiter: ',',
		}
		csvStr := "row1,value1,value2\nrow2,value1,value2,"
		expectedError := errors.New("record on line 2: wrong number of fields")
		result, err := cf.Reader(strings.NewReader(csvStr))
		t.Run("it should return an error ", func(t *testing.T) {
			assert.Equal(t, expectedError.Error(), err.Error())
			assert.Empty(t, result)
		})
	})

	t.Run("When user pass valid CSV data", func(t *testing.T) {
		cf := CSVFile{
			LazyQuotes:      true,
			Delimiter:       ',',
			FieldsPerRecord: -1,
		}
		csvStr := "row1,value1,value2\nrow2,value1,`value2"
		result, err := cf.Reader(strings.NewReader(csvStr))
		t.Run("it should return csv data ", func(t *testing.T) {
			assert.NoError(t, err)
			assert.NotEmpty(t, result)
		})
	})

}
