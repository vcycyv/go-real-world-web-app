package entity

type Book struct {
	Base

	Name string
	User string
}

// func (d *Book) BeforeCreate(tx *gorm.DB) error {
// 	d.ID = uuid.New().String()
// 	return nil
// }
