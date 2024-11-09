package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	foodMap := make(map[string]food)

	foodMap["burger"] = food{preptime: 15}
	foodMap["chips"] = food{preptime: 10}
	foodMap["nuggets"] = food{preptime: 12}

	if food, exists := foodMap[order]; exists {
		return food.preptime
	}
	return 404
}

// func main() {
// 	fmt.Println(FoodDeliveryTime(""))
// 	fmt.Println(FoodDeliveryTime("chips"))
// 	fmt.Println(FoodDeliveryTime("nuggets"))
// 	fmt.Println(FoodDeliveryTime("burger") + FoodDeliveryTime("chips") + FoodDeliveryTime("nuggets"))
// }
