package handlers

import (
	"lukedawe/hutchi/model"
	"lukedawe/hutchi/util"
	"net/http"

	"github.com/gin-gonic/gin"
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

	var categories []model.Category
	if err := h.DB.Model(&model.Category{}).
		Scopes(util.Paginate(params.Page, params.PageSize)).
		Preload("Breeds").
		Find(&categories).Error; err != nil {
		HandleError(c, http.StatusInternalServerError, err, "")
		return
	}

	c.JSON(http.StatusOK, categories)
}
