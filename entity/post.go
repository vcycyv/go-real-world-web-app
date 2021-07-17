package entity

type Post struct {
	Base

	Name string
	User string
}

// func (d *Post) BeforeCreate(tx *gorm.DB) error {
// 	d.ID = uuid.New().String()
// 	return nil
// }
