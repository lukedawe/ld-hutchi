package handlers

import (
	"lukedawe/hutchi/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GetCategoryToBreedsRequest struct {
	Name string `uri:"name" binding:"required"`
}

func (h *Handler) GetCategoryToBreeds(c *gin.Context) {
	var request GetCategoryToBreedsRequest
	if err := c.ShouldBindUri(&request); err != nil {
		HandleError(c, http.StatusBadRequest, err, "")
		return
	}

	// This is OK because the name is unique in the database.
	category, err := gorm.G[models.Category](h.DB).
		Preload("Breeds", nil).
		Where("name = ?", request.Name).
		First(c)

	if err != nil {
		HandleError(c, http.StatusBadRequest, err, "")
		return
	}

	response := make([]string, len(category.Breeds))
	for i, breed := range category.Breeds {
		response[i] = breed.Name
	}

	c.JSON(http.StatusOK, response)
}
