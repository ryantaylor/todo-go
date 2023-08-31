package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"net/http"
	"todo/db"
	"todo/todo"
	"todo/todo_list"
	"todo/user"
)

type ServerDatabase interface {
	db.CreateRecordDB
	db.ListRecordDB
	db.FindRecordDB
}

func Setup(db ServerDatabase) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(render.SetContentType(render.ContentTypeJSON))

	userRepo := user.NewRepository(db)
	userMiddleware := user.NewMiddleware(&userRepo)
	userHandler := user.NewHandler(&userRepo)

	todoListRepo := todo_list.NewRepository(db)
	todoListMiddleware := todo_list.NewMiddleware(&todoListRepo)
	todoListHandler := todo_list.NewHandler(&todoListRepo)

	todoRepo := todo.NewRepository(db)
	todoHandler := todo.NewHandler(&todoRepo)

	router.Get("/status", func(writer http.ResponseWriter, req *http.Request) {
		_, _ = writer.Write([]byte("ok"))
	})

	router.Route("/users", func(router chi.Router) {
		router.Post("/", userHandler.Create)
		router.Route("/{userID}", func(router chi.Router) {
			router.Use(userMiddleware.Context)
			router.Get("/", userHandler.Find)
			router.Route("/todo_lists", func(router chi.Router) {
				router.Get("/", todoListHandler.ListForUser)
				router.Post("/", todoListHandler.CreateForUser)
				router.Route("/{todoListID}", func(router chi.Router) {
					router.Use(todoListMiddleware.Context)
					router.Get("/", todoListHandler.FindForUser)
					router.Route("/todos", func(router chi.Router) {
						router.Get("/", todoHandler.ListForTodoList)
						router.Post("/", todoHandler.CreateForTodoList)
					})
				})
			})
		})
	})

	return router
}
