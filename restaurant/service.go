package restaurant

type restaurantService struct{}

func (rs restaurantService) SaveBulkRestaurantData(restaurantData [][]string) error {
	p := NewPool(10, len(restaurantData), restaurantData)
	p.Run()
	return nil
}

func NewRestaurantService() RestaurantService {
	return restaurantService{}
}
