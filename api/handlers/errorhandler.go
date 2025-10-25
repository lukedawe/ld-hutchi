package handlers

import (
	"log"
	"lukedawe/hutchi/dtos/responses/errors"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next() // Process the request.

		if len(ctx.Errors) > 0 {

			log.Println("Error detected.")

			err := ctx.Errors.Last().Err
			// Map the error to the API-specific error code.
			response, ok := err.(errors.ErrorResponse)
			if !ok {
				response = errors.ErrInternalUnknown.SetError(err)
			}
			log.Printf("API Error [%d]: %s - Request: %s %s",
				response.Status, err.Error(), ctx.Request.Method, ctx.Request.URL.Path)
			if gin.IsDebugging() {
				ctx.AbortWithStatusJSON(response.Status, response)
				return
			}
			ctx.AbortWithStatusJSON(response.Status, response.ToProductionErrorStruct())
			return
		}
	}
}
