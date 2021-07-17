package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/vcycyv/blog/handler"
	"github.com/vcycyv/blog/infrastructure"
	infra "github.com/vcycyv/blog/infrastructure"
	"github.com/vcycyv/blog/infrastructure/repository"
	"github.com/vcycyv/blog/middleware"
	"github.com/vcycyv/blog/service"
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
	postRepo := repository.NewPostRepo(db)
	postService := service.NewPostService(postRepo)
	postHandler := handler.NewPostHandler(postService, authService)
	authHandler := handler.NewAuthHandler(authService)

	r.POST("/auth", authHandler.GetAuth)

	api := r.Group("/")

	api.Use(middleware.NewJWTMiddleware(authService).JWT())
	{
		api.GET("/posts/:id", postHandler.Get)
		api.GET("/posts", postHandler.GetAll)
		api.POST("/posts", postHandler.Add)
		api.PUT("/posts/:id", postHandler.Update)
		api.DELETE("/posts/:id", postHandler.Delete)
	}
	return r
}
