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

func CreateCategory(db *gorm.DB, c context.Context, category models.Category) error {
	return gorm.G[models.Category](db).Create(c, &category)
}

func CreateCategories(db *gorm.DB, c context.Context, category []models.Category) error {
	return gorm.G[models.Category](db).CreateInBatches(c, &category, 10)
}

/*
At the moment this is redundant, because we are finding the category by name, and then
only updating the name (basically a no-op), but in the future if there is more data, then
we will want to upsert the data.

GORM performs all single operations as a transaction, but this will require a few
operations to be performed.
*/
func UpsertCategory(db *gorm.DB, c context.Context, category models.Category, oldName string) error {
	// Begin transaction (this will cover a few queries that we must perform).
	return db.
		WithContext(c).
		Transaction(func(tx *gorm.DB) error {

			// oldCategory, err := gorm.G[models.Category](tx).Where("name = ?", oldName).First(c)
			// if err != nil {
			// 	// Error can be ignored if it's only saying that a record was not found.
			// 	if !errors.Is(err, gorm.ErrRecordNotFound) {
			// 		return err
			// 	}
			// }

			// // TODO: https://stackoverflow.com/questions/77935191/how-to-upsert-a-gorm-model-with-its-association

			// if err := gorm.G[models.Category](
			// 	tx.
			// 		Clauses(clause.OnConflict{
			// 			Columns:   []clause.Column{{Name: "name"}},
			// 			DoUpdates: clause.AssignmentColumns([]string{"name"}), // Upsert data.
			// 		})).
			// 	Create(c, &category); err != nil {
			// 	return err
			// }

			// if err := tx.
			// 	Model(category).
			// 	Association("breeds").
			// 	Clear(); err != nil {
			// 	return err
			// }

			// // TODO: Make this idempotent!
			// if err := tx.Model(category).
			// 	Association("breeds").
			// 	Append(category.Breeds); err != nil {
			// 	return err
			// }

			return nil
		})
}
