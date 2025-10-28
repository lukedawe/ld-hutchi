package services

import (
	"context"
	"errors"
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

func GetCategoryByName(db *gorm.DB, c context.Context, name string) (models.Category, error) {
	return gorm.G[models.Category](db).
		Preload("Breeds", nil).
		Where("name = ?", name).
		// This is OK because the name is unique in the database.
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
func UpsertCategory(db *gorm.DB, c context.Context, upsertCat models.Category, oldName string) error {
	// Begin transaction (this will cover a couple of queries that we must perform).
	return db.
		WithContext(c).
		Transaction(func(tx *gorm.DB) error {
			oldCategory, err := gorm.G[models.Category](tx).Where("name = ?", oldName).First(c)
			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}

			// This could still be 0 iff the category does not already exist.
			upsertCat.ID = oldCategory.ID

			err = tx.
				Clauses(
					clause.OnConflict{
						Columns:   []clause.Column{{Name: "id"}},
						UpdateAll: true,
					}).
				Create(&upsertCat).Error

			if err != nil {
				return err
			}

			return tx.
				Model(&upsertCat).
				Unscoped().
				Association("Breeds").
				Unscoped().
				Replace(upsertCat.Breeds)
		})
}
