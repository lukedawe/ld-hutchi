package handlers

import (
	"log"
	"lukedawe/hutchi/dtos/requests"
	"lukedawe/hutchi/dtos/responses"
	"lukedawe/hutchi/dtos/responses/errors"
	"lukedawe/hutchi/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// func (h *Handler) AddBreed(c *gin.Context) {
// 	var request requests.AddBreed

// 	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
// 		HandleError(c, http.StatusBadRequest, err, "")
// 		return
// 	}

// 	log.Println("Response binded successfully")

// 	// Find the category

// 	breedModel:=&models.Breed{
// 		Name: request.Name,
// 		Category: ,,
// 	}

// 	// Send to the database
// 	if err := gorm.G[models.Category](h.DB).Create(c, &categoryModel); err != nil {
// 		HandleError(c, http.StatusInternalServerError, err, "")
// 		return
// 	}

// 	c.Status(http.StatusCreated)
// }

func (h *Handler) GetBreed(c *gin.Context) {
	var request requests.GetBreed
	if err := c.ShouldBindUri(&request); err != nil {
		c.Error(errors.ErrBadRequestBinding.SetError(err))
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
