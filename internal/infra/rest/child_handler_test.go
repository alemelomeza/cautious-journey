package rest

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alemelomeza/cautious-journey/internal/app"
	"github.com/alemelomeza/cautious-journey/internal/domain"
	"github.com/alemelomeza/cautious-journey/internal/infra/db"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestChildHandler(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
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
		srv := app.NewChildService(repoMock, repoMock)
		h := NewChildHandler(srv)

		r := mux.NewRouter()
		r.HandleFunc("/child/{name}", h.Get)
		req := httptest.NewRequest(http.MethodGet, "/child/"+c.Name, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
		assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
	})

	t.Run("unhappy path", func(t *testing.T) {
		repoMock := new(db.DBMock)
		repoMock.On("GetChild", mock.AnythingOfType("string")).Return(&domain.Child{}, errors.New("any error"))
		srv := app.NewChildService(repoMock, repoMock)
		h := NewChildHandler(srv)

		r := mux.NewRouter()
		r.HandleFunc("/child/{name}", h.Get)
		req := httptest.NewRequest(http.MethodGet, "/child/Nelson", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})
}
