package responses

type RecommendationResponse struct {
	TransactionID   string     `json:"transactionId"`
	DoMedicalSurvey bool       `json:"doMedicalSurvey"`
	Recommendations []Merchant `json:"recommendations"`
}

type Merchant struct {
	MerchantID   int                  `json:"merchantId"`
	MerchantName string               `json:"merchantName"`
	Ratings      string               `json:"ratings"`
	Foods        []FoodRecommendation `json:"foods"`
}

type FoodRecommendation struct {
	FoodId int    `json:"foodId"`
	Name   string `json:"name"`
	Price  int    `json:"price"`
}
