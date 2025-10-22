package handlers

import (
	"io"
	"lukedawe/hutchi/generated/proto_dogs"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/proto"
)

func (h *Handler) AddCategory(c *gin.Context) {
	addMessage := &proto_dogs.CategoryAdd{}
	message, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if err := proto.Unmarshal(message, addMessage); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.Status(http.StatusCreated)
}
