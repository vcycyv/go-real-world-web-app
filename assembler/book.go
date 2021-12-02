package assembler

import (
	"fmt"
	"net/http"

	"github.com/vcycyv/bookshop/entity"
	rep "github.com/vcycyv/bookshop/representation"
)

type BookAssembler struct{}

func NewBookAssembler() BookAssembler {
	return BookAssembler{}
}

func (s *BookAssembler) ToData(rep rep.Book) *entity.Book {
	return &entity.Book{
		Base: entity.Base{
			ID:        rep.ID,
			CreatedAt: rep.CreatedAt,
			UpdatedAt: rep.UpdatedAt,
		},

		Name: rep.Name,
		User: rep.User,
	}
}

func (s *BookAssembler) ToRepresentation(data entity.Book) *rep.Book {
	return &rep.Book{
		Base: rep.Base{
			ID:        data.ID,
			CreatedAt: data.CreatedAt,
			UpdatedAt: data.UpdatedAt,

			Links: []rep.ResourceLink{
				{
					Rel:    "self",
					Method: http.MethodGet,
					Href:   fmt.Sprintf("/books/%s", data.ID),
				},
				{
					Rel:    "add-book",
					Method: http.MethodPost,
					Href:   "/books",
				},
				{
					Rel:    "edit-book",
					Method: http.MethodPut,
					Href:   fmt.Sprintf("/books/%s", data.ID),
				},
				{
					Rel:    "delete-book",
					Method: http.MethodDelete,
					Href:   fmt.Sprintf("/books/%s", data.ID),
				},
			},
		},

		Name: data.Name,
		User: data.User,
	}
}
