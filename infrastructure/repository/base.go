package repository

import (
	logger "github.com/sirupsen/logrus"

	"gorm.io/gorm"

	"github.com/google/uuid"
	"github.com/vcycyv/bookshop/entity"
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
		logger.Fatal("failed to register uuid hook")
	}

	sqlDB, _ := db.DB()
	sqlDB.SetConnMaxIdleTime(10)
	sqlDB.SetMaxOpenConns(100)

	migrate(db)
}

func migrate(db *gorm.DB) {
	err := db.AutoMigrate(&entity.Book{})
	if err != nil {
		logger.Fatal("migration failed.")
	}
}
