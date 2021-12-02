package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/vcycyv/bookshop/handler"
	"github.com/vcycyv/bookshop/infrastructure"
	infra "github.com/vcycyv/bookshop/infrastructure"
	"github.com/vcycyv/bookshop/infrastructure/repository"
	"github.com/vcycyv/bookshop/middleware"
	"github.com/vcycyv/bookshop/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	router := initRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", infra.AppSetting.HTTPPort),
		Handler:        router,
		ReadTimeout:    time.Duration(60) * time.Second,
		WriteTimeout:   time.Duration(60) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}

func createDB() *gorm.DB {
	var (
		dbName, user, password, host string
		port                         int
	)

	dbName = infra.DatabaseSetting.DBName
	user = infra.DatabaseSetting.DBUser
	password = infra.DatabaseSetting.DBPassword
	host = infra.DatabaseSetting.DBHost
	port = infra.DatabaseSetting.DBPort

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbName, port),
		PreferSimpleProtocol: true,
	}), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}})

	if err != nil {
		log.Fatal("failed to open database")
	}

	return db
}

func initRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.JSONAppErrorReporter())

	logrus.SetLevel(logrus.DebugLevel)
	r.Use(middleware.LoggerToFile(logrus.StandardLogger()))
	r.Use(middleware.CORS())
	gin.SetMode(infra.AppSetting.RunMode)

	db := createDB()
	repository.InitDB(db)

	authService := infrastructure.NewAuthService()
	bookRepo := repository.NewBookRepo(db)
	bookService := service.NewBookService(bookRepo)
	bookHandler := handler.NewBookHandler(bookService, authService)
	authHandler := handler.NewAuthHandler(authService)

	r.POST("/auth", authHandler.GetAuth)

	api := r.Group("/")

	api.Use(middleware.NewJWTMiddleware(authService).JWT())
	{
		api.GET("/books/:id", bookHandler.Get)
		api.GET("/books", bookHandler.GetAll)
		api.POST("/books", bookHandler.Add)
		api.PUT("/books/:id", bookHandler.Update)
		api.DELETE("/books/:id", bookHandler.Delete)
	}
	return r
}
