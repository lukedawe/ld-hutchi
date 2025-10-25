package handlers

import (
	"log"
	"lukedawe/hutchi/dtos/requests"
	"lukedawe/hutchi/dtos/responses"
	"lukedawe/hutchi/dtos/responses/errors"
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
		c.Error(errors.ErrBadRequestBinding.SetError(err))
		return
	}

	categories, err := services.GetCategoriesToBreeds(h.DB, c, request.Page, request.PageSize)
	if err != nil {
		c.Error(services.TranslateDbError(err))
		return
	}

	response := make([]responses.Category, len(categories))
	for i, category := range categories {
		response[i] = categoryModelToResponse(category)
	}

	c.JSON(http.StatusOK, response)
}

func (h *Handler) GetCategory(c *gin.Context) {
	var request requests.GetCategory
	if err := c.ShouldBindUri(&request); err != nil {
		c.Error(errors.ErrBadRequestBinding.SetError(err))
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
		c.Error(errors.ErrBadRequestBinding.SetError(err))
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

func (h *Handler) AddCategory(c *gin.Context) {
	var request requests.AddCategory
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.Error(errors.ErrBadRequestBinding.SetError(err))
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
	if err := services.CreateCategory(h.DB, c, &categoryModel); err != nil {
		c.Error(services.TranslateDbError(err))
		return
	}

	response := categoryModelToResponse(categoryModel)
	c.JSON(http.StatusCreated, response)
}

func (h *Handler) AddCategories(c *gin.Context) {
	var request requests.AddCategories
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.Error(errors.ErrBadRequestBinding.SetError(err))
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

	var response responses.Categories
	response.Categories = make([]responses.Category, len(categoryModels))
	for i, category := range categoryModels {
		response.Categories[i] = categoryModelToResponse(category)
	}

	c.JSON(http.StatusCreated, response)
}

// Helper functions for conversion between the DB model and the responses
func categoryModelToResponse(categoryModel models.Category) responses.Category {
	breeds := make([]responses.CategoryBreed, len(categoryModel.Breeds))
	for i, breed := range categoryModel.Breeds {
		breeds[i].Name = breed.Name
	}

	return responses.Category{
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
