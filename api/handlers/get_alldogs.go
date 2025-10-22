package handlers

import (
	"lukedawe/hutchi/model"
	"lukedawe/hutchi/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// NOTE: We would never normally have a call like this, which retrieves the whole database at a time,
// so the code here is not perfect (it copies the whole database into a protobufs message).
// Possibly I could make this endpoint a json that trivially takes the result of the query and returns
// it to the client.
func (h *Handler) GetAllDogs(c *gin.Context) {
	var categories []model.Category
	if err := h.DB.Model(&model.Category{}).
		Preload("Breeds").
		Find(&categories).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Encode into a protobufs response
	allDogs := util.CategoriesToAllDogs(categories)

	c.ProtoBuf(http.StatusFound, allDogs)
}
