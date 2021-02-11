package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/graphql-go/handler"
)

func RegisterRoutes(r *chi.Mux) *chi.Mux {
	/* GraphQL */
	graphQL := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})
	r.Use(middleware.Logger)
	r.Handle("/query", graphQL)
	return r
}
