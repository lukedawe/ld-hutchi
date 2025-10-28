package main

import (
	"database/sql"
	"log"
	"lukedawe/hutchi/handlers"
	"lukedawe/hutchi/util"
	"net/http"
	"os"
	"strconv"
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

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalln(port)
	}

	dsn := util.GetDsn(
		os.Getenv("DB_HOST"),
		uint(port),
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
	}), &gorm.Config{TranslateError: true})

	if err != nil {
		log.Fatalln(err)
	}

	return gormDb
}

func setupRouter(DB *sql.DB) *gin.Engine {
	r := gin.Default()
	// Register middleware before roots
	r.Use(handlers.ErrorHandler())

	h := &handlers.Handler{DB: getGormDb(DB)}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/v1")

	/*
		There are valid security concerns for this API. Using database ID's
		as unique resource identifiers is not best practice, but considering
		the data stored within the database is not sensitive, it makes querying a lot
		easier.
	*/

	v1.GET("/breeds/categories/:page/:page_size", h.GetCategoriesToBreeds) // Get all categories mapped to breeds (paginated).
	v1.GET("/categories", h.GetCategories)                                 // Get all categories (without breed information).
	v1.GET("/category/:id", h.GetCategory)                                 // Get the category for a category name.
	v1.GET("/categories/:id/breeds", h.GetCategoryToBreeds)                // Get all the breeds for a particular breed.
	v1.GET("/breed/:id", h.GetBreed)                                       // Get all breeds with a particular name.

	v1.POST("/category", h.PostCategory)     // Add a category.
	v1.POST("/breed", h.PostBreed)           // Add a breed.
	v1.POST("/categories", h.PostCategories) // Batch add categories.

	v1.PUT("/category/*id", h.PutCategory) // Upsert category.
	v1.PUT("/breed/*id")                   // Upsert breed.

	v1.PATCH("/category/:id") // Update a category.
	v1.PATCH("/breed/:id")    // Update a breed.

	v1.DELETE("/breed/:id")    // Delete a breed.
	v1.DELETE("/category/:id") // Delete a category.

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

	log.Fatalln(r.Run().Error())

	defer func() {
		if DB != nil {
			DB.Close()
		}
	}()
}
