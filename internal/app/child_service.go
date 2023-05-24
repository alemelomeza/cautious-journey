package app

import (
	"errors"

	"github.com/alemelomeza/cautious-journey/internal/domain"
)

const (
	ErrParentNotValid   string = "Parent not valid"
	ErrChildNotValid    string = "Child not valid"
	ErrChildNotFound    string = "Child not found try again"
	ErrChildrenNotFound string = "Childrens not found"
)

type ChildRepository interface {
	SaveChild(child domain.Child) error
	GetChild(name string) (*domain.Child, error)
	GetChildrens() ([]*domain.Child, error)
}

type ParentRepository interface {
	SaveParent(parent domain.Parent) error
	GetParent(name string) (*domain.Parent, error)
	GetParents(childName string) ([]*domain.Parent, error)
}

type ChildService struct {
	childRepo  ChildRepository
	parentRepo ParentRepository
}

func NewChildService(
	childRepo ChildRepository,
	parentRepo ParentRepository) *ChildService {
	return &ChildService{
		childRepo:  childRepo,
		parentRepo: parentRepo,
	}
}

func (s *ChildService) CreateChild(name string, age int, parents ...domain.Parent) error {
	c := domain.Child{
		Name:    name,
		Age:     age,
		Parents: parents,
	}
	if !c.Validate() {
		return errors.New(ErrChildNotValid)
	}
	for _, p := range c.Parents {
		if !p.Validate() {
			return errors.New(ErrParentNotValid)
		}
	}
	err := s.childRepo.SaveChild(c)
	if err != nil {
		return errors.New("error trying to save children")
	}
	for _, p := range c.Parents {
		err := s.parentRepo.SaveParent(p)
		if err != nil {
			return errors.New("error trying to saver a parent")
		}
	}
	return nil
}

func (s *ChildService) GetChild(name string) (*domain.Child, error) {
	c, err := s.childRepo.GetChild(name)
	if err != nil {
		return nil, errors.New(ErrChildNotFound)
	}
	return c, nil
}

func (s *ChildService) GetChildrens() ([]*domain.Child, error) {
	c, err := s.childRepo.GetChildrens()
	if err != nil {
		return nil, errors.New(ErrChildrenNotFound)
	}
	return c, nil
}
