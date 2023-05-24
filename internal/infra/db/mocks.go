package db

import (
	"github.com/alemelomeza/cautious-journey/internal/domain"
	"github.com/stretchr/testify/mock"
)

type DBMock struct {
	mock.Mock
}

func (m *DBMock) SaveChild(child domain.Child) error {
	args := m.Called(child)
	return args.Error(0)
}

func (m *DBMock) GetChild(name string) (*domain.Child, error) {
	args := m.Called(name)
	return args.Get(0).(*domain.Child), args.Error(1)
}

func (m *DBMock) GetChildrens() ([]*domain.Child, error) {
	args := m.Called()
	return args.Get(0).([]*domain.Child), args.Error(1)
}

func (m *DBMock) SaveParent(parent domain.Parent) error {
	args := m.Called(parent)
	return args.Error(0)
}

func (m *DBMock) GetParent(name string) (*domain.Parent, error) {
	args := m.Called(name)
	return args.Get(0).(*domain.Parent), args.Error(1)
}

func (m *DBMock) GetParents(childName string) ([]*domain.Parent, error) {
	args := m.Called(childName)
	return args.Get(0).([]*domain.Parent), args.Error(1)
}
