package representation

type Book struct {
	Base

	Name string `json:"name"`
	User string `json:"user"`
}
