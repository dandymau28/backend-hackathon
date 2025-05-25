package responses

type Checkout struct {
	TransactionID string         `json:"transactionId"`
	Foods         []FoodCheckout `json:"foods"`
}

type FoodCheckout struct {
	ID                     int      `json:"id"`
	Name                   string   `json:"name"`
	Quantity               int      `json:"quantity"`
	TotalPrice             int      `json:"totalPrice"`
	Nutrients              []string `json:"nutrients"`
	Calories               int      `json:"calories"`
	HighlightedIngredients []string `json:"highlightedIngredients"`
	Label                  string   `json:"label"`
}
