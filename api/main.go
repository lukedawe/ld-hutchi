package main

import (
	"database/sql"
	"log"
	"lukedawe/hutchi/handlers"
	"lukedawe/hutchi/util"
	"net/http"
	"os"

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

	v1.GET("/all_dogs", h.GetAllDogs)
	v1.POST("/add_category", h.AddCategory)

	return r
}

func main() {
	// gin.SetMode(gin.ReleaseMode)

	DB := connectDB("dev.env")

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
