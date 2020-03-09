package uploader

import (
	"io"

	"github.com/stretchr/testify/mock"
)

type MockUploaderService struct {
	mock.Mock
}

func (m *MockUploaderService) SaveBulkdata(fileData [][]string) error {
	args := m.Called(fileData)
	return args.Error(0)
}

type MockXLSXFileService struct {
	mock.Mock
}

func (m *MockXLSXFileService) Reader(file io.Reader) ([][]string, error) {
	args := m.Called(file)
	mockresult := make([][]string, 0)
	return mockresult, args.Error(1)
}
