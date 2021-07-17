package mock

import (
	logger "github.com/sirupsen/logrus"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func CreateDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"),
		&gorm.Config{NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		}})
	if err != nil {
		logger.Fatal("failed to open db")
	}

	return db
}
