package uploader

import (
	"context"
	"io"

	"github.com/stretchr/testify/mock"
)

type MockUploaderService struct {
	mock.Mock
}

func (m *MockUploaderService) SaveBulkdata(ctx context.Context, moduleName string, userID int64, fileData [][]string) [][]string {
	_ = m.Called(ctx, moduleName, userID, fileData)
	mockresult := make([][]string, 0)
	return mockresult
}

type MockXLSXFileService struct {
	mock.Mock
}

func (m *MockXLSXFileService) Reader(file io.Reader) ([][]string, error) {
	args := m.Called(file)
	mockresult := make([][]string, 0)
	return mockresult, args.Error(1)
}
