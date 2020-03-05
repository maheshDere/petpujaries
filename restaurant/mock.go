package restaurant

import "github.com/stretchr/testify/mock"

type MockRestaurantService struct {
	mock.Mock
}

func (m *MockRestaurantService) SaveBulkRestaurantData(restaurantData [][]string) error {
	args := m.Called(restaurantData)
	return args.Error(0)
}
