package handlers

import (
	"lukedawe/hutchi/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GetBreedRequest struct {
	Name string `uri:"name" binding:"required"`
}

func (h *Handler) GetBreed(c *gin.Context) {
	var request GetBreedRequest
	if err := c.ShouldBindUri(&request); err != nil {
		HandleError(c, http.StatusBadRequest, err, "")
		return
	}

	// Breeds are not unique in the database, so we need to return a list.
	breeds, err := gorm.G[models.Breed](h.DB).
		Preload("Category", nil).
		Where("name = ?", request.Name).
		Find(c)

	if err != nil {
		HandleError(c, http.StatusBadRequest, err, "")
		return
	}

	response := models.BreedsToDtoWithCategory(breeds)

	c.JSON(http.StatusOK, response)
}
