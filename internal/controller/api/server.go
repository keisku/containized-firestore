package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type server struct{}

// New .
func New() http.Handler {
	s := server{}
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("status ok"))
	})
	r.Route("/account", func(r chi.Router) {
		r.Post("/", s.CreateAccount)
		r.Get("/", s.LoadUsers)
		r.Get("/{id}", s.GetAccount)
		r.Post("/{id}", s.UpdateAccount)
		r.Delete("/{id}", s.DeleteAccount)
	})
	return r
}

// CreateAccount .
func (s server) CreateAccount(w http.ResponseWriter, r *http.Request) {}

// LoadUsers .
func (s server) LoadUsers(w http.ResponseWriter, r *http.Request) {}

// GetAccount .
func (s server) GetAccount(w http.ResponseWriter, r *http.Request) {}

// UpdateAccount .
func (s server) UpdateAccount(w http.ResponseWriter, r *http.Request) {}

// DeleteAccount .
func (s server) DeleteAccount(w http.ResponseWriter, r *http.Request) {}
