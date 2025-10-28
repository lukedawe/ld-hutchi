package requests

import (
	"lukedawe/hutchi/handlers/dtos/responses/errors"

	"github.com/gin-gonic/gin"
)

// All types that implement this interface should return a response type
type ValidatableRequest interface {
	Validate() error
}

func ValidateRequestAndBindJson(ctx *gin.Context, request ValidatableRequest) error {
	if err := ctx.ShouldBindBodyWithJSON(request); err != nil {
		return errors.ErrBadRequestBinding.SetError(err)
	}

	if err := request.Validate(); err != nil {
		return errors.ErrBadRequestValidation.SetError(err)
	}

	return nil
}
