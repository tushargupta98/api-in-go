package domain

import (
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	"github.com/tushargupta98/api-in-go/cache"
)

func DomainRouter(r chi.Router, db *sqlx.DB, cache cache.RedisClient) {
	repo := NewDomainRepository(db, cache)
	handler := &DomainHandler{repo}

	r.Get("/domain", handler.List)
	r.Post("/domain", handler.Create)
	r.Get("/domain/{id}", handler.Get)
	r.Put("/domain/{id}", handler.Update)
	r.Delete("/domain/{id}", handler.Delete)
}
