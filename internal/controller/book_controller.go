package controller

import (
	"encoding/json"
	"go-lib-api/internal/dto"
	"go-lib-api/internal/service"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type BookController struct {
	service service.BookService
}

func NewBookController(s service.BookService) *BookController {
	return &BookController{service: s}
}

func (c *BookController) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateBookRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	resp, err := c.service.CreateBook(r.Context(), req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func (c *BookController) Get(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	resp, err := c.service.GetBook(r.Context(), id)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}