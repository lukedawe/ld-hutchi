package handlers

import (
	"lukedawe/hutchi/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h *Handler) GetCategories(c *gin.Context) {
	categories, err := gorm.G[models.Category](h.DB).
		Find(c)

	if err != nil {
		HandleError(c, http.StatusInternalServerError, err, "")
		return
	}

	categoryList := make([]string, len(categories))

	for i, category := range categories {
		categoryList[i] = category.Name
	}

	c.JSON(http.StatusOK, categoryList)
}
