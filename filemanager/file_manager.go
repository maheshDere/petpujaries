package filemanager

import (
	"bytes"
	"encoding/csv"
	"io"
	"io/ioutil"
	"petpujaris/logger"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

type XLSXFile struct {
}

type FileOperation interface {
	Reader(r io.Reader) ([][]string, error)
}

func (xf XLSXFile) Reader(file io.Reader) ([][]string, error) {
	f, err := excelize.OpenReader(file)
	if err != nil {
		logger.LogError(err, "XLSXFile Reader", "error to get file pointer")
		return nil, err
	}

	sheetName := f.GetSheetName(1)
	result, err := f.GetRows(sheetName)
	if err != nil {
		logger.LogError(err, "XLSXFile Reader", "error to get file data")
		return nil, err
	}
	return result, nil
}

func NewXLSXFileService() FileOperation {
	return XLSXFile{}
}

type CSVFile struct {
	LazyQuotes      bool
	Delimiter       rune
	FieldsPerRecord int
}

func (cf CSVFile) Reader(r io.Reader) ([][]string, error) {
	var results [][]string
	byteData, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(bytes.NewReader(byteData))
	reader.LazyQuotes = cf.LazyQuotes
	reader.Comma = cf.Delimiter
	reader.FieldsPerRecord = cf.FieldsPerRecord

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

func NewCSVFileService(lazyQuotes bool, delimiter rune, fieldsPerRecord int) FileOperation {
	return CSVFile{LazyQuotes: lazyQuotes, Delimiter: delimiter, FieldsPerRecord: fieldsPerRecord}
}
