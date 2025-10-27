package handlers

import (
	"log"
	"lukedawe/hutchi/handlers/dtos/requests"
	"lukedawe/hutchi/handlers/dtos/responses"
	"lukedawe/hutchi/handlers/dtos/responses/errors"
	error_responses "lukedawe/hutchi/handlers/dtos/responses/errors"
	"lukedawe/hutchi/models"
	"lukedawe/hutchi/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetBreed(c *gin.Context) {
	var request requests.GetBreed
	if err := c.ShouldBindUri(&request); err != nil {
		c.Error(error_responses.ErrBadRequestBinding.SetError(err))
		return
	}

	breeds, err := services.GetBreeds(h.DB, c, request.Name)
	if err != nil {
		c.Error(services.TranslateDbError(err))
		return
	}

	response := make([]responses.BreedWithCategory, len(breeds))
	for i, breed := range breeds {
		response[i] = responses.BreedWithCategory{
			Name:     breed.Name,
			Category: breed.Category.Name,
		}
	}

	log.Println("Got breed: ", response)

	c.JSON(http.StatusOK, response)
}

func (h *Handler) PostBreed(c *gin.Context) {
	var request requests.AddBreed
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.Error(errors.ErrBadRequestBinding.SetError(err))
		return
	}

	if err := request.Validate(); err != nil {
		c.Error(err)
		return
	}

	// Find the category
	category, err := services.GetCategoryByName(h.DB, c, request.CategoryName)
	if err != nil {
		c.Error(services.TranslateDbError(err))
		return
	}

	breedModel := &models.Breed{
		Name:     request.Name,
		Category: category,
	}

	// Send to the database
	if err := services.CreateBreed(h.DB, c, breedModel); err != nil {
		c.Error(services.TranslateDbError(err))
		return
	}

	response := responses.BreedCreated{
		Name: breedModel.Name,
	}

	// Create the response
	c.JSON(http.StatusCreated, response)
}
