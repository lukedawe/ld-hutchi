package handlers

import (
	"log"
	"lukedawe/hutchi/dtos/requests"
	"lukedawe/hutchi/dtos/responses"
	"lukedawe/hutchi/models"
	"lukedawe/hutchi/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetCategories(c *gin.Context) {
	categories, err := services.GetCategories(h.DB, c)

	if err != nil {
		HandleError(c, http.StatusInternalServerError, err, "")
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
		HandleError(c, http.StatusBadRequest, err, "Uri Malformed")
		return
	}

	categories, err := services.GetCategoriesToBreeds(h.DB, c, request.Page, request.PageSize)
	if err != nil {
		HandleError(c, http.StatusInternalServerError, err, "")
		return
	}

	response := make([]responses.Category, len(categories))
	for i, category := range categories {
		response[i] = categoryToResponse(&category)
	}

	c.JSON(http.StatusOK, response)
}

func (h *Handler) GetCategory(c *gin.Context) {
	var request requests.GetCategory
	if err := c.ShouldBindUri(&request); err != nil {
		HandleError(c, http.StatusBadRequest, err, "")
		return
	}

	category, err := services.GetCategory(h.DB, c, request.Name)

	if err != nil {
		HandleError(c, http.StatusBadRequest, err, "")
		return
	}

	response := categoryToResponse(&category)
	log.Println(response)
	c.JSON(http.StatusOK, response)
}

func (h *Handler) GetCategoryToBreeds(c *gin.Context) {
	var request requests.GetCategoryToBreeds
	if err := c.ShouldBindUri(&request); err != nil {
		HandleError(c, http.StatusBadRequest, err, "")
		return
	}

	category, err := services.GetCategory(h.DB, c, request.Name)
	if err != nil {
		HandleError(c, http.StatusBadRequest, err, "")
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
		HandleError(c, http.StatusBadRequest, err, "")
		return
	}

	categoryModel := models.Category{Name: request.Name}
	categoryModel.Breeds = make([]models.Breed, len(request.Breeds))
	for i, breed := range request.Breeds {
		categoryModel.Breeds[i] = models.Breed{
			Name: breed,
		}
	}
	// Send to the database
	if err := services.CreateCategory(h.DB, c, categoryModel); err != nil {
		HandleError(c, http.StatusInternalServerError, err, "")
		return
	}

	c.Status(http.StatusCreated)
}

// Helper functions for conversion between the DB model and the responses
func categoryToResponse(categoryModel *models.Category) responses.Category {
	breeds := make([]responses.Breed, len(categoryModel.Breeds))
	for i, breed := range categoryModel.Breeds {
		breeds[i] = responses.Breed{
			Name: breed.Name,
		}
	}

	return responses.Category{
		Name:   categoryModel.Name,
		Breeds: breeds,
	}
}
