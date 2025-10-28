package handlers

import (
	"log"
	"lukedawe/hutchi/handlers/dtos/requests"
	"lukedawe/hutchi/handlers/dtos/responses"
	response_errors "lukedawe/hutchi/handlers/dtos/responses/errors"
	"lukedawe/hutchi/models"
	"lukedawe/hutchi/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetCategories(c *gin.Context) {
	categories, err := services.GetCategories(h.DB, c)
	if err != nil {
		c.Error(services.TranslateDbError(err))
		return
	}

	categoryList := make([]string, len(categories))
	for i, category := range categories {
		categoryList[i] = category.Name
	}

	c.JSON(http.StatusOK, categoryList)
}

// Paginated results for categories to breeds
func (h *Handler) GetCategoriesToBreeds(c *gin.Context) {
	var request requests.GetCategoriesToBreeds
	if err := c.ShouldBindUri(&request); err != nil {
		c.Error(response_errors.ErrBadRequestBinding.SetError(err))
		return
	}

	categories, err := services.GetCategoriesToBreeds(h.DB, c, request.Page, request.PageSize)
	if err != nil {
		c.Error(services.TranslateDbError(err))
		return
	}

	response := make([]responses.CategoryCreated, len(categories))
	for i, category := range categories {
		response[i] = categoryModelToResponse(category)
	}

	c.JSON(http.StatusOK, response)
}

func (h *Handler) GetCategory(c *gin.Context) {
	var request requests.GetCategory
	if err := c.ShouldBindUri(&request); err != nil {
		c.Error(response_errors.ErrBadRequestBinding.SetError(err))
		return
	}

	category, err := services.GetCategoryById(h.DB, c, request.Id)
	if err != nil {
		c.Error(services.TranslateDbError(err))
		return
	}

	response := categoryModelToResponse(category)
	log.Println(response)
	c.JSON(http.StatusOK, response)
}

func (h *Handler) GetCategoryToBreeds(c *gin.Context) {
	var request requests.GetCategoryToBreeds
	if err := c.ShouldBindUri(&request); err != nil {
		c.Error(response_errors.ErrBadRequestBinding.SetError(err))
		return
	}

	category, err := services.GetCategoryById(h.DB, c, request.Id)
	if err != nil {
		c.Error(services.TranslateDbError(err))
		return
	}

	response := make([]string, len(category.Breeds))
	for i, breed := range category.Breeds {
		response[i] = breed.Name
	}

	c.JSON(http.StatusOK, response)
}

func (h *Handler) PostCategory(c *gin.Context) {
	var body requests.AddCategoryJson
	if err := requests.ValidateRequestAndBindJson(c, &body); err != nil {
		c.Error(err)
		return
	}

	categoryModel := addCategoryRequestToModel(body)
	// Send to the database
	if err := services.CreateCategory(h.DB, c, categoryModel); err != nil {
		c.Error(services.TranslateDbError(err))
		return
	}

	response := categoryModelToResponse(categoryModel)
	c.JSON(http.StatusCreated, response)
}

func (h *Handler) PostCategories(c *gin.Context) {
	var body requests.AddCategories
	if err := requests.ValidateRequestAndBindJson(c, &body); err != nil {
		c.Error(err)
		return
	}

	categoryModels := make([]models.Category, len(body.Categories))
	for i, category := range body.Categories {
		categoryModels[i] = addCategoryRequestToModel(category)
	}

	if err := services.CreateCategories(h.DB, c, categoryModels); err != nil {
		c.Error(services.TranslateDbError(err))
		return
	}

	var response responses.CategoriesCreated
	response.Categories = make([]responses.CategoryCreated, len(categoryModels))
	for i, category := range categoryModels {
		response.Categories[i] = categoryModelToResponse(category)
	}

	c.JSON(http.StatusCreated, response)
}

func (h *Handler) PutCategory(c *gin.Context) {
	var uri requests.PutCategoryUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.Error(response_errors.ErrBadRequestBinding.SetError(err))
		return
	}

	var body requests.AddCategoryJson
	if err := requests.ValidateRequestAndBindJson(c, &body); err != nil {
		c.Error(err)
		return
	}

	categoryModel := addCategoryRequestToModel(body)
	categoryModel.ID = uri.Id

	if err := services.UpsertCategory(h.DB, c, &categoryModel); err != nil {
		c.Error(services.TranslateDbError(err))
		return
	}

	c.JSON(http.StatusCreated, categoryModelToResponse(categoryModel))
}

// Updates the name of the category.
func (h *Handler) PatchCategory(c *gin.Context) {
	var uri requests.PatchCategoryUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.Error(response_errors.ErrBadRequestInvalidParam.SetError(err))
		return
	}

	var body requests.PatchCategoryBody
	if err := requests.ValidateRequestAndBindJson(c, &body); err != nil {
		c.Error(err)
		return
	}

	model := &models.Category{
		Name: body.Name,
		ID:   uri.Id,
	}
	err := services.UpdateCategoryName(h.DB, c, model)
	if err != nil {
		c.Error(services.TranslateDbError(err))
		return
	}

	c.JSON(http.StatusOK, categoryModelToResponse(*model))
}

func (h *Handler) DeleteCategory(c *gin.Context) {
	var uri requests.DeleteCategoryUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.Error(response_errors.ErrBadRequestBinding.SetError(err))
		return
	}

	if err := services.DeleteCategory(h.DB, c, uri.Id); err != nil {
		c.Error(services.TranslateDbError(err))
		return
	}

	c.Status(http.StatusOK)
}

// Helper functions for conversion between the DB model and the responses.
func categoryModelToResponse(categoryModel models.Category) responses.CategoryCreated {
	breeds := make([]responses.CategoryBreed, len(categoryModel.Breeds))
	for i, breed := range categoryModel.Breeds {
		breeds[i].Name = breed.Name
		breeds[i].Id = breed.ID
	}

	return responses.CategoryCreated{
		Name:   categoryModel.Name,
		Breeds: breeds,
		Id:     categoryModel.ID,
	}
}

func addCategoryRequestToModel(body requests.AddCategoryJson) models.Category {
	model := models.Category{Name: body.Name}
	model.Breeds = make([]models.Breed, len(body.Breeds))
	for i, breed := range body.Breeds {
		model.Breeds[i] = models.Breed{
			Name: breed.Name,
		}
	}
	return model
}
