package main

import (
	"context"
	"log"
	"lukedawe/hutchi/generated/proto"
	"os"

	"github.com/jackc/pgx"
	"github.com/joho/godotenv"
	"google.golang.org/protobuf/encoding/protojson"
)

// Sets up its own connection to the database.
func CreateDb(envFile string) {
	if err := godotenv.Load(envFile); err != nil {
		log.Fatalln(err.Error())
	}

	config := pgx.ConnConfig{
		Database: os.Getenv("DB_NAME"),
		Host:     os.Getenv("DB_HOST"),
		Port:     5431, // TODO: Pain to get this out of the env file.
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	conn, err := pgx.Connect(config)
	if err != nil {
		log.Fatalln(err.Error())
	}

	if err := conn.Ping(context.TODO()); err != nil {
		log.Fatalln(err.Error())
	}

	log.Println("Successfully connected to the database")

	defer conn.Close()

	// Now we have the connection to the database, we can attempt to add
	// the data.
	fileBytes, err := os.ReadFile("./common/dogs.json")
	if err != nil {
		log.Fatalln(err.Error())
	}

	dogs := &proto.AllDogs{}
	if err := protojson.Unmarshal(fileBytes, dogs); err != nil {
		log.Fatalln(err.Error())
	}

	for _, category := range dogs.Categories {
		log.Println(category.Name)
	}

	// Now we have the dogs, we can send them to the database.
}
