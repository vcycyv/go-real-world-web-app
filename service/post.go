package service

import (
	"github.com/vcycyv/blog/assembler"
	"github.com/vcycyv/blog/domain"
	rep "github.com/vcycyv/blog/representation"
)

type postService struct {
	postRepo domain.PostRepository
}

func NewPostService(postRepo domain.PostRepository) domain.PostInterface {
	return &postService{
		postRepo,
	}
}

func (s *postService) Add(post rep.Post) (*rep.Post, error) {
	data, err := s.postRepo.Add(*assembler.PostAss.ToData(post))
	if err != nil {
		return &rep.Post{}, err
	}
	return assembler.PostAss.ToRepresentation(*data), nil
}

func (s *postService) Get(id string) (*rep.Post, error) {
	data, err := s.postRepo.Get(id)
	if err != nil {
		return nil, err
	}
	return assembler.PostAss.ToRepresentation(*data), nil
}

func (s *postService) GetAll() ([]*rep.Post, error) {
	posts, err := s.postRepo.GetAll()
	if err != nil {
		return nil, err
	}

	rtnVal := []*rep.Post{}
	for _, post := range posts {
		rtnVal = append(rtnVal, assembler.PostAss.ToRepresentation(*post))
	}
	return rtnVal, nil
}

func (s *postService) Update(post rep.Post) (*rep.Post, error) {
	data, err := s.postRepo.Update(*assembler.PostAss.ToData(post))
	if err != nil {
		return nil, err
	}

	return assembler.PostAss.ToRepresentation(*data), nil
}

func (s *postService) Delete(id string) error {
	err := s.postRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
