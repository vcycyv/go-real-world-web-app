package entity

import "time"

type Base struct {
	ID        string `gorm:"column:id;type:uuid;primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
