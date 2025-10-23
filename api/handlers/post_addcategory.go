package handlers

import (
	"log"
	"lukedawe/hutchi/models"
	"lukedawe/hutchi/requests"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h *Handler) AddCategory(c *gin.Context) {
	var request requests.AddCategoryMessage

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
