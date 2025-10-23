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

	c.Status(http.StatusCreated)
}
