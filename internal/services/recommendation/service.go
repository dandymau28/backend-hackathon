package recommendation

import (
	"errors"
	"log"
	"strings"
	"sync"

	"github.com/go-resty/resty/v2"
	"github.com/nix-united/golang-echo-boilerplate/internal/dto"
	"github.com/nix-united/golang-echo-boilerplate/internal/models"
	"github.com/nix-united/golang-echo-boilerplate/internal/repositories"
	"github.com/nix-united/golang-echo-boilerplate/internal/responses"
)

type RecommendationService interface {
	RecommendFood(userPrompt string) (responses.RecommendationResponse, error)
}

type recommendationService struct {
	allergyRepository repositories.AllergyRepository
	diseaseRepository repositories.DiseaseRepository
}

func NewRecommendationService(allergyRepository repositories.AllergyRepository, diseaseRepository repositories.DiseaseRepository) RecommendationService {
	return &recommendationService{
		allergyRepository: allergyRepository,
		diseaseRepository: diseaseRepository,
	}
}

func (s *recommendationService) RecommendFood(userPrompt string) (responses.RecommendationResponse, error) {
	var (
		wg    sync.WaitGroup
		mutex sync.Mutex
	)

	userId := 1

	allergies, err := s.allergyRepository.GetAllergiesByUser(userId)

	if err != nil {
		log.Printf("cannot get allergies: %v", err.Error())
		return responses.RecommendationResponse{}, err
	}

	diseases, err := s.diseaseRepository.GetDiseasesByUser(userId)

	if err != nil {
		log.Printf("cannot get diseases: %v", err.Error())
		return responses.RecommendationResponse{}, err
	}

	if len(allergies) == 0 || len(diseases) == 0 {
		log.Println("no allergies or diseases:")
		return responses.RecommendationResponse{
			DoMedicalSurvey: true,
		}, nil
	}

	var allergiesSlice []string
	if len(allergies) != 0 {
		allergiesSlice = extractAllergies(&allergies)
	}

	var allergySituation string
	for _, allergy := range allergies {
		if allergySituation != "" {
			allergySituation += ", "
		}
		allergySituation += allergy.Allergy.Name
	}

	var diseaseSituation string
	for _, disease := range diseases {
		if diseaseSituation != "" {
			diseaseSituation += ", "
		}
		diseaseSituation += disease.Disease.Name
	}

	additionalPrompt := "I have medical situation "

	if diseaseSituation != "" {
		additionalPrompt += diseaseSituation
	}

	if allergySituation != "" {
		additionalPrompt += " and allergy to "
		additionalPrompt += allergySituation
	}

	var result dto.GenerateResponse

	client := resty.New()

	resp, err := client.R().SetHeader("Content-Type", "application/json").
		SetBody(dto.GenerateRequest{
			Prompt:           userPrompt,
			AdditionalPrompt: additionalPrompt,
			SessionID:        "123123123",
		}).
		SetResult(&result).
		Post("http://8.215.31.210:8000/generate")

	if err != nil {
		log.Printf("cannot generate recommendation: %v", err.Error())
		return responses.RecommendationResponse{}, err
	}

	if resp.StatusCode() != 200 {
		log.Printf("status code not 200")
		return responses.RecommendationResponse{}, errors.New("status code not 200")
	}

	recommendedFoods := result.FoodNames
	log.Printf("recommendedFoods: %v", recommendedFoods)

	// resty to prompt
	searchResult := []dto.Search{
		dto.Search{
			MerchantID:   1,
			MerchantName: "Sate Madura",
			Ratings:      4,
			Foods: []dto.SearchFood{
				dto.SearchFood{
					FoodID:    1,
					FoodName:  "Chicken Satay",
					FoodPrice: 25000,
					Ingredients: []dto.SearchFoodIngredient{
						dto.SearchFoodIngredient{
							IngredientID:   1,
							IngredientName: "Kacang",
						},
						dto.SearchFoodIngredient{
							IngredientID:   2,
							IngredientName: "Ayam",
						},
					},
				},
				dto.SearchFood{
					FoodID:    1,
					FoodName:  "Seafood Fried Rice",
					FoodPrice: 30000,
					Ingredients: []dto.SearchFoodIngredient{
						dto.SearchFoodIngredient{
							IngredientID:   1,
							IngredientName: "Udang",
						},
						dto.SearchFoodIngredient{
							IngredientID:   2,
							IngredientName: "Cumi",
						},
					},
				},
				dto.SearchFood{
					FoodID:    1,
					FoodName:  "Fried Rice",
					FoodPrice: 20000,
					Ingredients: []dto.SearchFoodIngredient{
						dto.SearchFoodIngredient{
							IngredientID:   1,
							IngredientName: "Daging",
						},
					},
				},
			},
		},
	}

	selectedFoods := make([]responses.Merchant, 0, 3)

	for _, itemSearch := range searchResult {
		wg.Add(1)
		go func(item dto.Search) {
			defer wg.Done()
			aspects := map[string]bool{
				"foodMatch":      false,
				"ingredientSafe": true,
			}

			var sortedFoods []responses.FoodRecommendation

			for _, food := range item.Foods {
				if aspects["foodMatch"] = contains(recommendedFoods, strings.ToLower(food.FoodName)); aspects["foodMatch"] {
					for _, ingredient := range food.Ingredients {
						if contains(allergiesSlice, ingredient.IngredientName) {
							aspects["ingredientSafe"] = true
							sortedFoods = append(sortedFoods, responses.FoodRecommendation{
								FoodId: food.FoodID,
								Name:   food.FoodName,
								Price:  food.FoodPrice,
							})
						}
					}
				}
			}

			if aspects["foodMatch"] && aspects["ingredientSafe"] {
				mutex.Lock()
				selectedMerchant := responses.Merchant{
					MerchantID:   item.MerchantID,
					MerchantName: item.MerchantName,
					Ratings:      string(item.Ratings),
					Foods:        sortedFoods,
				}
				selectedFoods = append(selectedFoods, selectedMerchant)
				mutex.Unlock()
			}

		}(itemSearch)
	}
	// scan and takeout inedible ingredients

	// cache data

	wg.Wait()

	return responses.RecommendationResponse{
		TransactionID:   "123123123123",
		DoMedicalSurvey: false,
		Recommendations: selectedFoods,
	}, nil
}

func contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func extractAllergies(allergies *[]models.MedicalAllergy) (result []string) {
	for _, allergy := range *allergies {
		result = append(result, allergy.Allergy.Name)
	}
	return result
}
