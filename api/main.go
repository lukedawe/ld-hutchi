package main

import (
	"database/sql"
	"log"
	"lukedawe/hutchi/handlers"
	"lukedawe/hutchi/util"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectDB(envFile string) *sql.DB {
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

	DB, err := sql.Open("pgx", dsn)

	if err != nil {
		log.Fatalln(err.Error())
	}

	if err := DB.Ping(); err != nil {
		log.Fatalln(err.Error())
	}

	log.Println("Successfully connected to the database")

	// Configure the database connection.
	DB.SetMaxOpenConns(10)
	DB.SetConnMaxLifetime(time.Hour)

	return DB
}

func getGormDb(sqldb *sql.DB) *gorm.DB {
	gormDb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqldb,
	}), &gorm.Config{})

	if err != nil {
		if gin.IsDebugging() {
			log.Fatalln(err)
		}
	}

	return gormDb
}

func setupRouter(DB *sql.DB) *gin.Engine {
	r := gin.Default()

	h := &handlers.Handler{DB: getGormDb(DB)}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/v1")

	v1.GET("/breeds/categories/:page/:page_size", h.GetCategoriesToBreeds) // Get all categories mapped to breeds.
	v1.GET("/categories", h.GetCategories)                                 // Get all categories.
	v1.GET("/category/:name", h.GetCategory)                               // Get the category for a category name.
	v1.GET("/category/:name/breeds", h.GetCategoryToBreeds)                // Get all the breeds for a particular breed.
	v1.GET("/breed/:name", h.GetBreed)                                     // Get a particular breed.
	v1.POST("/category", h.AddCategory)                                    // Add a category.
	v1.POST("/breed")                                                      // Add a breed.
	v1.PUT("/categories")                                                  // Batch add categories
	v1.DELETE("/breed")                                                    // Delete a breed.
	v1.DELETE("/category")                                                 // Delete a category.

	return r
}

func main() {
	// gin.SetMode(gin.ReleaseMode)

	var envFileName string
	if gin.IsDebugging() {
		envFileName = "dev.env"
	} else {
		envFileName = ".env"
	}

	DB := connectDB(envFileName)

	if DB == nil {
		log.Fatalln("DB connection is nil")
	}

	r := setupRouter(DB)

	log.Fatalln(r.Run(":8081").Error())

	defer func() {
		if DB != nil {
			DB.Close()
		}
	}()
}
