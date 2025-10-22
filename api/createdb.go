package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"lukedawe/hutchi/util"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

// Sets up its own connection to the database.
func PopulateDb(envFile string) {
	if err := godotenv.Load(envFile); err != nil {
		log.Fatalln(err.Error())
	}

	dsn := util.GetDsn(
		os.Getenv("DB_HOST"),
		5432,
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	conn, err := sql.Open("pgx", dsn)

	if err != nil {
		log.Fatalln(err.Error())
	}

	if err := conn.Ping(); err != nil {
		log.Fatalln(err.Error())
	}

	log.Println("Successfully connected to the database")

	defer conn.Close()

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
		Conn: conn,
	}), &gorm.Config{})

	if err != nil {
		log.Fatalln(err.Error())
	}

	for category, breeds := range allDogs {
		ctx := context.Background()
		gormCategory := &Category{Name: category}
		if err := gorm.G[Category](gormDb).Create(ctx, gormCategory); err != nil {
			log.Fatalln(err.Error())
		}

		for _, breed := range breeds {
			// Create the breeds the category contains
			if err := gorm.G[Breed](gormDb).Create(ctx, &Breed{Name: breed, Category: gormCategory.ID}); err != nil {
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
