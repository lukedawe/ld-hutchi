package handlers

import (
	"log"
	"lukedawe/hutchi/handlers/dtos/requests"
	"lukedawe/hutchi/handlers/dtos/responses"
	error_responses "lukedawe/hutchi/handlers/dtos/responses/errors"
	response_errors "lukedawe/hutchi/handlers/dtos/responses/errors"
	"lukedawe/hutchi/models"
	"lukedawe/hutchi/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath v1

// @Summary Get breed by ID.
// @Schemes
// @Description Get the breed by the ID.
// @Tags breeds
// @Produce json
// @Success 200 {object} []responses.BreedWithCategory
// @Router /breed/{id} [get]
func (h *Handler) GetBreed(c *gin.Context) {
	var body requests.GetBreed
	if err := c.ShouldBindUri(&body); err != nil {
		c.Error(error_responses.ErrBadRequestBinding.SetError(err))
		return
	}

	breeds, err := services.GetBreeds(h.DB, c, body.Id)
	if err != nil {
		c.Error(services.TranslateDbError(err))
		return
	}

	response := make([]responses.BreedWithCategory, len(breeds))
	for i, breed := range breeds {
		response[i] = responses.BreedWithCategory{
			Name:     breed.Name,
			Category: breed.Category.Name,
		}
	}

	log.Println("Got breed: ", response)

	c.JSON(http.StatusOK, response)
}

func (h *Handler) PostBreed(c *gin.Context) {
	var body requests.AddBreed
	if err := requests.ValidateRequestAndBindJson(c, &body); err != nil {
		c.Error(err)
		return
	}
	breedModel := &models.Breed{
		Name:       body.Name,
		CategoryID: body.CategoryId,
	}

	// Send to the database
	if err := services.CreateBreed(h.DB, c, breedModel); err != nil {
		c.Error(services.TranslateDbError(err))
		return
	}

	response := responses.BreedCreated{
		Name:       breedModel.Name,
		Id:         breedModel.ID,
		CategoryId: breedModel.CategoryID,
	}

	// Create the response
	c.JSON(http.StatusCreated, response)
}

func (h *Handler) PutBreed(c *gin.Context) {
	var uri requests.PutBreedUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.Error(response_errors.ErrBadRequestInvalidParam.SetError(err))
		return
	}

	// NOTE: For time's sake I'm reusing the AddBreed request because they are
	//	going to be the same, but really this should have it's own struct.
	var body requests.AddBreed
	if err := requests.ValidateRequestAndBindJson(c, &body); err != nil {
		c.Error(err)
		return
	}

	model := models.Breed{Name: body.Name, ID: uri.Id, CategoryID: body.CategoryId}

	if err := services.UpsertBreed(h.DB, c, &model); err != nil {
		c.Error(services.TranslateDbError(err))
		return
	}

	response := breedModelToResponse(model)
	c.JSON(http.StatusCreated, response)
}

func (h *Handler) PatchBreed(c *gin.Context) {
	var uri requests.PatchBreedUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.Error(response_errors.ErrBadRequestBinding.SetError(err))
		return
	}

	var body requests.PatchBreedBody
	if err := requests.ValidateRequestAndBindJson(c, &body); err != nil {
		c.Error(err)
		return
	}

	model := models.Breed{
		Name: body.Name,
		ID:   uri.Id,
	}

	if err := services.UpdateBreedName(h.DB, c, &model); err != nil {
		c.Error(services.TranslateDbError(err))
		return
	}

	response := breedModelToResponse(model)
	c.JSON(http.StatusOK, response)
}

func (h *Handler) DeleteBreed(c *gin.Context) {
	var uri requests.DeleteBreedUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.Error(response_errors.ErrBadRequestBinding.SetError(err))
		return
	}

	if err := services.DeleteBreed(h.DB, c, uri.Id); err != nil {
		c.Error(services.TranslateDbError(err))
		return
	}

	c.Status(http.StatusOK)
}

func breedModelToResponse(model models.Breed) responses.BreedCreated {
	return responses.BreedCreated{
		Id:         model.ID,
		Name:       model.Name,
		CategoryId: model.CategoryID,
	}
}
