package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"net/http"
	"todo/user"
)

func main() {
	db := sqlx.MustConnect("pgx", "postgres://hellofreshdev:hellofreshdev@localhost:5432/todo?sslmode=disable")

	router := chi.NewRouter()

	router.Use(render.SetContentType(render.ContentTypeJSON))

	userRepo := user.NewRepository(db)
	userMiddleware := user.NewMiddleware(&userRepo)
	userHandler := user.NewHandler(&userRepo)

	router.Get("/status", func(writer http.ResponseWriter, req *http.Request) {
		writer.Write([]byte("ok"))
	})

	router.Route("/users", func(router chi.Router) {
		router.Post("/", userHandler.Create)
		router.Route("/{userID}", func(router chi.Router) {
			router.Use(userMiddleware.Context)
			router.Get("/", userHandler.Find)
		})
	})

	println("Server starting!")

	http.ListenAndServe(":8080", router)
}
