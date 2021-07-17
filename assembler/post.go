package assembler

import (
	"fmt"
	"net/http"

	"github.com/vcycyv/blog/entity"
	rep "github.com/vcycyv/blog/representation"
)

type PostAssembler struct{}

func NewPostAssembler() PostAssembler {
	return PostAssembler{}
}

func (s *PostAssembler) ToData(rep rep.Post) *entity.Post {
	return &entity.Post{
		Base: entity.Base{
			ID:        rep.ID,
			CreatedAt: rep.CreatedAt,
			UpdatedAt: rep.UpdatedAt,
		},

		Name: rep.Name,
		User: rep.User,
	}
}

func (s *PostAssembler) ToRepresentation(data entity.Post) *rep.Post {
	return &rep.Post{
		Base: rep.Base{
			ID:        data.ID,
			CreatedAt: data.CreatedAt,
			UpdatedAt: data.UpdatedAt,

			Links: []rep.ResourceLink{
				{
					Rel:    "self",
					Method: http.MethodGet,
					Href:   fmt.Sprintf("/posts/%s", data.ID),
				},
				{
					Rel:    "add-post",
					Method: http.MethodPost,
					Href:   "/posts",
				},
				{
					Rel:    "edit-post",
					Method: http.MethodPut,
					Href:   fmt.Sprintf("/posts/%s", data.ID),
				},
				{
					Rel:    "delete-post",
					Method: http.MethodDelete,
					Href:   fmt.Sprintf("/posts/%s", data.ID),
				},
			},
		},

		Name: data.Name,
		User: data.User,
	}
}
