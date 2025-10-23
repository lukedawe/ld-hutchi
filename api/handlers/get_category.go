package handlers

import (
	"log"
	"lukedawe/hutchi/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GetCategoryRequest struct {
	Name string `uri:"name" binding:"required"`
}

func (h *Handler) GetCategory(c *gin.Context) {
	var request GetCategoryRequest
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

	response := category.ToResponse()

	log.Println(response)

	c.JSON(http.StatusOK, response)
}
