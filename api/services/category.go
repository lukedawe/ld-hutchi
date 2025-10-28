package services

import (
	"context"
	"lukedawe/hutchi/models"
	"lukedawe/hutchi/services/scopes"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func GetCategoryById(db *gorm.DB, c context.Context, id uint) (models.Category, error) {
	return gorm.G[models.Category](db).
		Preload("Breeds", nil).
		Where("id = ?", id).
		First(c)
}

func CreateCategory(db *gorm.DB, c context.Context, category models.Category) error {
	return gorm.G[models.Category](db).Create(c, &category)
}

func CreateCategories(db *gorm.DB, c context.Context, category []models.Category) error {
	return gorm.G[models.Category](db).CreateInBatches(c, &category, 10)
}

/*
GORM performs all single operations as a transaction, but this will require a few
operations to be performed.
*/
func UpsertCategory(db *gorm.DB, c context.Context, upsertCat *models.Category) error {
	// Begin transaction (this will cover a couple of queries that we must perform).
	return db.
		WithContext(c).
		Transaction(func(tx *gorm.DB) error {
			err := tx.
				Clauses(
					clause.OnConflict{
						Columns:   []clause.Column{{Name: "id"}},
						UpdateAll: true,
					}).
				Create(&upsertCat).Error

			if err != nil {
				return err
			}

			// Replace all the associated breeds.
			return tx.
				Model(&upsertCat).
				Unscoped().
				Association("Breeds").
				Unscoped().
				Replace(upsertCat.Breeds)
		})
}

func UpdateCategoryName(db *gorm.DB, c context.Context, updatedCategory *models.Category) error {
	return db.WithContext(c).Transaction(func(tx *gorm.DB) error {
		rowsAffected, err := gorm.G[models.Category](tx).
			Where("id = ?", updatedCategory.ID).
			Update(c, "name", updatedCategory.Name)

		if err != nil {
			return err
		}

		if rowsAffected == 0 {
			return gorm.ErrRecordNotFound
		}

		model, err := gorm.G[models.Category](tx).Preload("Breeds", nil).Where("id = ?", updatedCategory.ID).First(c)
		if err != nil {
			return err
		}

		updatedCategory = &model
		return nil
	})
}

// This is OK because the delete will cascade to the breeds table.
func DeleteCategory(db *gorm.DB, c context.Context, id uint) error {
	rowsAffected, err := gorm.G[models.Category](db).Where("id = ?", id).Delete(c)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
