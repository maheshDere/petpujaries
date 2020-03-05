package restaurant

import (
	"io"

	"github.com/stretchr/testify/mock"
)

type MockRestaurantService struct {
	mock.Mock
}

func (m *MockRestaurantService) SaveBulkRestaurantData(restaurantData [][]string) error {
	args := m.Called(restaurantData)
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
