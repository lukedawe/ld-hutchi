package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AddCategory(c *gin.Context) {
	// addMessage := &proto_dogs.CategoryAdd{}
	// message, err := io.ReadAll(c.Request.Body)
	// if err != nil {
	// 	HandleError(c, http.StatusBadRequest, err, "")
	// 	return
	// }
	// if err := proto.Unmarshal(message, addMessage); err != nil {
	// 	HandleError(c, http.StatusBadRequest, err, "")
	// 	return
	// }

	// var request requests.AddCategoryMessage
	// if err := c.ShouldBindBodyWithJSON(&request); err != nil {
	// 	HandleError(c, http.StatusBadRequest, err, "")
	// }

	// if request != nil {

	// }

	c.Status(http.StatusCreated)
}
