package services

import (
	"context"
	"lukedawe/hutchi/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func UpsertBreed(db *gorm.DB, c context.Context, upsertBreed *models.Breed) error {
	return db.WithContext(c).Transaction(func(tx *gorm.DB) error {
		return tx.Clauses(
			clause.OnConflict{
				Columns:   []clause.Column{{Name: "id"}},
				UpdateAll: true,
			},
		).Create(&upsertBreed).Error
	})
}

func UpdateBreedName(db *gorm.DB, c context.Context, updatedBreed *models.Breed) error {
	return db.WithContext(c).Transaction(func(tx *gorm.DB) error {
		rowsAffected, err := gorm.G[models.Breed](db).
			Where("id = ?", updatedBreed.ID).
			Update(c, "name", updatedBreed.Name)

		if err != nil {
			return err
		}

		if rowsAffected == 0 {
			return gorm.ErrRecordNotFound
		}

		model, err := gorm.G[models.Breed](tx).Where("id = ?", updatedBreed.ID).First(c)

		if err != nil {
			return err
		}

		updatedBreed = &model
		return nil
	})
}
