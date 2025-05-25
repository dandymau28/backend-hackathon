package responses

type MerchantResponse struct {
	MerchantID   int           `json:"merchantId"`
	MerchantName string        `json:"merchantName"`
	Foods        []FoodDisplay `json:"foods"`
}

type FoodDisplay struct {
	Name    string `json:"name"`
	Ratings string `json:"ratings"`
	Price   int    `json:"price"`
	Image   string `json:"image"`
}
