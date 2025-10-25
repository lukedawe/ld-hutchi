package services

import (
	"context"
	"lukedawe/hutchi/models"
	"lukedawe/hutchi/services/scopes"

	"gorm.io/gorm"
)

func GetCategories(db *gorm.DB, c context.Context) ([]models.Category, error) {
	return gorm.G[models.Category](db).Find(c)
}

func GetCategoriesToBreeds(db *gorm.DB, c context.Context, page uint, pageSize uint) ([]models.Category, error) {
	return gorm.G[models.Category](db).
		Scopes(scopes.Paginate(page, pageSize)).
		Preload("Breeds", nil).
		Find(c)
}

func GetCategoryByName(db *gorm.DB, c context.Context, name string) (models.Category, error) {
	return gorm.G[models.Category](db).
		Preload("Breeds", nil).
		Where("name = ?", name).
		// This is OK because the name is unique in the database.
		First(c)
}

func CreateCategory(db *gorm.DB, c context.Context, category *models.Category) error {
	return gorm.G[models.Category](db).Create(c, category)
}

func CreateCategories(db *gorm.DB, c context.Context, category []models.Category) error {
	return gorm.G[models.Category](db).CreateInBatches(c, &category, 10)
}
