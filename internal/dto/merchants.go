package dto

type Search struct {
	MerchantID   int
	MerchantName string
	Ratings      int
	Foods        []SearchFood
}

type SearchFood struct {
	FoodID      int
	FoodName    string
	FoodPrice   int
	Ingredients []SearchFoodIngredient
}

type SearchFoodIngredient struct {
	IngredientID   int
	IngredientName string
}
