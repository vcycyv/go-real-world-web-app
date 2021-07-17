package repository

import (
	"log"

	"gorm.io/gorm"

	"github.com/google/uuid"
	"github.com/vcycyv/blog/entity"
)

func InitDB(db *gorm.DB) {
	var createCallback = func(db *gorm.DB) {
		idField := db.Statement.Schema.LookUpField("id")
		if idField != nil {
			_ = idField.Set(db.Statement.ReflectValue, uuid.New().String())
		}
	}

	err := db.Callback().Create().Before("gorm:create").Register("uuid", createCallback)
	if err != nil {
		log.Fatal("failed to register uuid hook")
	}

	sqlDB, _ := db.DB()
	sqlDB.SetConnMaxIdleTime(10)
	sqlDB.SetMaxOpenConns(100)

	migrate(db)
}

func migrate(db *gorm.DB) {
	err := db.AutoMigrate(&entity.Post{})
	if err != nil {
		log.Fatal("migration failed.")
	}
}
