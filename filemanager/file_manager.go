package filemanager

import (
	"bytes"
	"encoding/csv"
	"io"
	"io/ioutil"
)

type XLSXFile struct {
}

type CSVFile struct {
	LazyQuotes      bool
	Delimiter       rune
	FieldsPerRecord int64
}

type ParseFile interface {
	Parse(r io.Reader) [][]string
}

func (file *XLSXFile) Parse(r io.Reader) [][]string {
	fileData := make([][]string, 0)

	return fileData
}

func (file *CSVFile) Parse(r io.Reader) [][]string {
	fileData := make([][]string, 0)
	byteData, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(bytes.NewReader(byteData))
	reader.LazyQuotes = file.LazyQuotes
	reader.Comma = file.Delimiter
	reader.FieldsPerRecord = file.FieldsPerRecord

	return fileData
}
