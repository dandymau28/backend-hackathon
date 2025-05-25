package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nix-united/golang-echo-boilerplate/internal/repositories"
	"github.com/nix-united/golang-echo-boilerplate/internal/requests"
	"github.com/nix-united/golang-echo-boilerplate/internal/responses"
	s "github.com/nix-united/golang-echo-boilerplate/internal/server"
	"github.com/nix-united/golang-echo-boilerplate/internal/services/recommendation"
)

type RecommendationHandler struct {
	server *s.Server
}

func NewRecommendationHandler(server *s.Server) *RecommendationHandler {
	return &RecommendationHandler{server: server}
}

func (recommendationHandler *RecommendationHandler) Recommmend(c echo.Context) error {
	recommendationRequest := new(requests.Recommendation)

	if err := c.Bind(recommendationRequest); err != nil {
		log.Println("gagal")
		return err
	}

	if err := recommendationRequest.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty or not valid")
	}

	allergyRepository := repositories.NewAllergyRepository(recommendationHandler.server.DB)
	diseaseRepository := repositories.NewDiseaseRepository(recommendationHandler.server.DB)

	recomService := recommendation.NewRecommendationService(allergyRepository, diseaseRepository)
	response, err := recomService.RecommendFood(recommendationRequest.Prompt)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Server error")
	}

	return responses.Response(c, http.StatusOK, response)
}
