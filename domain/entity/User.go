package entity

type User struct {
	UserName          string
	RestaurantGroupID int64
	Name              string
	Role              string
}

func CreateUserFromSlide(raw interface{}) *User {
	result := &User{
		UserName:          raw.(map[string]interface{})["userName"].(string),
		RestaurantGroupID: int64(raw.(map[string]interface{})["restaurantGroupID"].(float64)),
		Name:              raw.(map[string]interface{})["name"].(string),
		Role:              raw.(map[string]interface{})["role"].(string),
	}
	return result
}
