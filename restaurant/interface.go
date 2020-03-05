package restaurant

type RestaurantService interface {
	SaveBulkRestaurantData(restaurantData [][]string) error
}
