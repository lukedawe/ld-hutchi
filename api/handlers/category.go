package handlers

import (
	"errors"
	"log"
	"lukedawe/hutchi/dtos/requests"
	"lukedawe/hutchi/dtos/responses"
	response_errors "lukedawe/hutchi/dtos/responses/errors"
	"lukedawe/hutchi/handlers/validation"
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

	category, err := services.GetCategoryByName(h.DB, c, request.Name)
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

	category, err := services.GetCategoryByName(h.DB, c, request.Name)
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
	var request requests.AddCategory
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.Error(response_errors.ErrBadRequestBinding.SetError(err))
		return
	}

	if err := validateAddCategoryRequest(request); err != nil {
		c.Error(err)
		return
	}

	categoryModel := models.Category{Name: request.Name}
	categoryModel.Breeds = make([]models.Breed, len(request.Breeds))
	for i, breed := range request.Breeds {
		categoryModel.Breeds[i] = models.Breed{
			Name: breed.Name,
		}
	}

	// Send to the database
	if err := services.CreateCategory(h.DB, c, categoryModel); err != nil {
		c.Error(services.TranslateDbError(err))
		return
	}

	response := categoryModelToResponse(categoryModel)
	c.JSON(http.StatusCreated, response)
}

func (h *Handler) PostCategories(c *gin.Context) {
	var request requests.AddCategories
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.Error(response_errors.ErrBadRequestBinding.SetError(err))
		return
	}

	// Validate each category individually.
	var err error
	for _, category := range request.Categories {
		err = errors.Join(validateAddCategoryRequest(category))
	}
	if err != nil {
		c.Error(err)
		return
	}

	categoryModels := make([]models.Category, len(request.Categories))
	for i, category := range request.Categories {
		categoryModels[i] = categoryRequestToModel(category)
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

	var body requests.PutCategoryBody
	if err := c.ShouldBindBodyWithJSON(&body); err != nil {
		c.Error(response_errors.ErrBadRequestInvalidJSON.SetError(err))
		return
	}

	if err := validatePutCategoryRequestUri(uri); err != nil {
		c.Error(err)
		return
	}

	if err := validatePutCategoryRequestBody(body); err != nil {
		c.Error(err)
		return
	}

	categoryModel := models.Category{Name: uri.Name}
	categoryModel.Breeds = make([]models.Breed, len(body.Breeds))
	for i, breed := range body.Breeds {
		categoryModel.Breeds[i] = models.Breed{
			Name: breed.Name,
		}
	}

	if err := services.UpsertCategory(h.DB, c, categoryModel); err != nil {
		c.Error(services.TranslateDbError(err))
		return
	}
}

// Helper functions for conversion between the DB model and the responses.
func categoryModelToResponse(categoryModel models.Category) responses.CategoryCreated {
	breeds := make([]responses.CategoryBreed, len(categoryModel.Breeds))
	for i, breed := range categoryModel.Breeds {
		breeds[i].Name = breed.Name
	}

	return responses.CategoryCreated{
		Name:   categoryModel.Name,
		Breeds: breeds,
	}
}

func categoryRequestToModel(categoryRequest requests.AddCategory) models.Category {
	breeds := make([]models.Breed, len(categoryRequest.Breeds))
	for j, breed := range categoryRequest.Breeds {
		breeds[j] = models.Breed{
			Name: breed.Name,
		}
	}

	return models.Category{
		Name:   categoryRequest.Name,
		Breeds: breeds,
	}
}

// Validates the add category request and returns nil or response-ready error.
func validateAddCategoryRequest(request requests.AddCategory) error {
	if err := validation.ValidateCategoryName(request.Name); err != nil {
		response := response_errors.ErrBadRequestInvalidParam.SetError(err)
		response.Message = err.Error() // Copy error message because it's user facing.
		return response
	}

	var err error
	for _, breed := range request.Breeds {
		err = errors.Join(validation.ValidateBreedName(breed.Name))
	}
	if err != nil {
		return err
	}

	return nil
}

func validatePutCategoryRequestUri(uri requests.PutCategoryUri) error {
	return validation.ValidateCategoryName(uri.Name)
}

func validatePutCategoryRequestBody(body requests.PutCategoryBody) error {
	var err error
	for _, breed := range body.Breeds {
		err = errors.Join(validation.ValidateBreedName(breed.Name))
	}
	if err != nil {
		return err
	}
	return nil
}
