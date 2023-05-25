package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/alemelomeza/cautious-journey/internal/app"
	"github.com/alemelomeza/cautious-journey/internal/infra/dto"
	"github.com/gorilla/mux"
)

type ChildHandler struct {
	childService *app.ChildService
}

func NewChildHandler(childService *app.ChildService) *ChildHandler {
	return &ChildHandler{
		childService: childService,
	}
}

func (h *ChildHandler) List(w http.ResponseWriter, r *http.Request) {
	c, err := h.childService.GetChildrens()
	if err != nil {
		log.Printf("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("%v", c)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, http.StatusText(http.StatusOK), fmt.Sprintf("%v", c))
}

func (h *ChildHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["name"] == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	c, err := h.childService.GetChild(vars["name"])
	if err != nil {
		log.Printf("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	cDTO := dto.ChildDTO{
		Name: c.Name,
		Age:  c.Age,
	}
	log.Printf("%v", c)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cDTO)
}
