package services

import (
	"context"
	"lukedawe/hutchi/models"

	"gorm.io/gorm"
)

func CreateBreed(db *gorm.DB, c context.Context, breed *models.Breed) error {
	return gorm.G[models.Breed](db).Create(c, breed)
}

func GetBreeds(db *gorm.DB, c context.Context, id uint) ([]models.Breed, error) {
	// Breeds are not unique in the database, so we need to return a list.
	return gorm.G[models.Breed](db).
		Preload("Category", nil).
		Where("id = ?", id).
		Find(c)
}
