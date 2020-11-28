package api

import (
	"context"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	appUser "github.com/kskumgk63/containized-firestore/internal/application/user"
	firestoreUser "github.com/kskumgk63/containized-firestore/internal/infrastructure/firestore/user"
	"github.com/kskumgk63/containized-firestore/pkg/env"
)

type server struct {
	account CRUDHandler
}

// New .
func New() http.Handler {
	ctx := context.Background()
	projectID, err := env.ProjectID()
	if err != nil {
		panic(err)
	}
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		panic(err)
	}
	s := server{
		newAccountCRUDHandler(
			appUser.NewUseCase(
				firestoreUser.NewRepository(client),
			),
		),
	}
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("status ok"))
	})
	r.Route("/account", func(r chi.Router) {
		r.Post("/", s.account.Create)
		r.Get("/", s.account.Read)
		r.Get("/{id}", s.account.Read)
		r.Post("/{id}", s.account.Update)
		r.Delete("/{id}", s.account.Delete)
	})
	return r
}
