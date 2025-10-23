package handlers

import (
	"log"
	"lukedawe/hutchi/dtos"
	"lukedawe/hutchi/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AddCategoryRequest struct {
	Category dtos.Category `json:"category" binding:"required"`
}

func (h *Handler) AddCategory(c *gin.Context) {
	var request AddCategoryRequest

	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		HandleError(c, http.StatusBadRequest, err, "")
		return
	}

	log.Println("Response binded successfully")

	categoryModel := models.CategoryDtoToModel(&request.Category)

	// Send to the database
	if err := gorm.G[models.Category](h.DB).Create(c, &categoryModel); err != nil {
		HandleError(c, http.StatusInternalServerError, err, "")
		return
	}

	c.Status(http.StatusCreated)
}
