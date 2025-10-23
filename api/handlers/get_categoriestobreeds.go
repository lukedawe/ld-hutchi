package handlers

import (
	"lukedawe/hutchi/model"
	"lukedawe/hutchi/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type getAllDogsParams struct {
	page     uint `uri:"page" binding:"required"`
	pageSize uint `uri:"page_size" binding:"required"`
}

// NOTE: We would never normally have a call like this, which retrieves the whole database at a time,
// so the code here is not perfect (it copies the whole database into a protobufs message).
// Possibly I could make this endpoint a json that trivially takes the result of the query and returns
// it to the client.
func (h *Handler) GetCategoriesToBreeds(c *gin.Context) {
	var params getAllDogsParams
	if err := c.ShouldBindUri(&params); err != nil {
		HandleError(c, http.StatusBadRequest, err, "Uri Malformed")
		return
	}

	var categories []model.Category
	if err := h.DB.Model(&model.Category{}).
		Scopes(util.Paginate(params.page, params.pageSize)).
		Preload("Breeds").
		Find(&categories).Error; err != nil {
		HandleError(c, http.StatusInternalServerError, err, "")
		return
	}

	c.JSON(http.StatusFound, categories)
}
