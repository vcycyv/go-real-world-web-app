package representation

type Post struct {
	Base

	Name string `json:"name"`
	User string `json:"user"`
}
