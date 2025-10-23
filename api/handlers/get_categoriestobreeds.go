package handlers

import (
	"lukedawe/hutchi/models"
	"lukedawe/hutchi/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type categoriesToBreedsParams struct {
	Page     uint `uri:"page" binding:"required"`
	PageSize uint `uri:"page_size" binding:"required"`
}

// Paginated results for categories to breeds
func (h *Handler) GetCategoriesToBreeds(c *gin.Context) {
	var params categoriesToBreedsParams
	if err := c.ShouldBindUri(&params); err != nil {
		HandleError(c, http.StatusBadRequest, err, "Uri Malformed")
		return
	}

	categories, err := gorm.G[models.Category](h.DB).
		Scopes(util.Paginate(params.Page, params.PageSize)).
		Preload("Breeds", nil).
		Find(c)

	if err != nil {
		HandleError(c, http.StatusInternalServerError, err, "")
		return
	}

	response := models.CategoriesToDTO(categories)

	c.JSON(http.StatusOK, response)
}
