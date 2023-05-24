package app

import (
	"errors"
	"testing"
	"github.com/alemelomeza/cautious-journey/internal/domain"
	"github.com/alemelomeza/cautious-journey/internal/infra/db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetChild(t *testing.T) {
	t.Run("should retrive a child passing a name", func(t *testing.T) {
		c := domain.Child{
			Name: "Andres",
			Age:  3,
			Parents: []domain.Parent{
				domain.Parent{
					Name:         "Nelson",
					Relationship: "father",
				},
			},
		}
		repoMock := new(db.DBMock)
		repoMock.On("GetChild", mock.AnythingOfType("string")).Return(&c, nil)
		srv := NewChildService(repoMock, repoMock)

		cResult, err := srv.GetChild(c.Name)

		assert.Empty(t, err)
		assert.NotEmpty(t, cResult)
		assert.Equal(t, c.Name, cResult.Name)
		assert.Equal(t, c.Age, cResult.Age)
	})

	t.Run("Unhappy Path DB Error ", func(t *testing.T) {
		c := domain.Child{
			Name: "Andres",
			Age:  3,
			Parents: []domain.Parent{
				domain.Parent{
					Name:         "Nelson",
					Relationship: "father",
				},
			},
		}

		repoMock := new(db.DBMock)
		repoMock.On("GetChild", mock.AnythingOfType("string")).Return(&domain.Child{}, errors.New("DB Err 8001"))
		srv := NewChildService(repoMock, repoMock)

		cRes, err := srv.GetChild(c.Name)

		assert.Empty(t, cRes)
		assert.NotEmpty(t, err)
		assert.ErrorContains(t, err, ErrChildNotFound)
	})

}
