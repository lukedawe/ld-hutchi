package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"lukedawe/hutchi/model"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Sets up its own connection to the database.
func PopulateDb(envFile string, DB *sql.DB) {
	// Now we have the connection to the database, we can attempt to add
	// the data.
	allDogs, err := readFile("common/dogs.json")
	if err != nil {
		log.Fatalln(err.Error())
	}

	if allDogs == nil {
		log.Fatalln("File contains no elements.")
	}

	// Now we have the dogs, we can send them to the database.
	gormDb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: DB,
	}), &gorm.Config{})

	if err != nil {
		log.Fatalln(err.Error())
	}

	for category, breeds := range allDogs {
		ctx := context.Background()
		gormCategory := &model.Category{Name: category}
		if err := gorm.G[model.Category](gormDb).Create(ctx, gormCategory); err != nil {
			log.Fatalln(err.Error())
		}

		for _, breed := range breeds {
			// Create the breeds the category contains
			if err := gorm.G[model.Breed](gormDb).Create(ctx, &model.Breed{Name: breed, CategoryID: gormCategory.ID}); err != nil {
				log.Fatalln(err.Error())
			}
		}
	}
}

func readFile(fileName string) (map[string][]string, error) {
	fileBytes, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var fileMap map[string][]string
	if err := json.Unmarshal(fileBytes, &fileMap); err != nil {
		return nil, err
	}

	return fileMap, nil
}
