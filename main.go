package main

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/simukti/sqldb-logger/logadapter/zapadapter"
	"go.uber.org/zap"
	"net/http"
	"todo/todo"
	"todo/todo_list"
	"todo/user"
)

func main() {
	zapCfg := zap.NewDevelopmentConfig()
	zapCfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel) // whatever minimum level
	zapCfg.DisableCaller = true
	logger, _ := zapCfg.Build()

	dsn := "postgres://hellofreshdev:hellofreshdev@localhost:5432/todo?sslmode=disable"
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		panic(err)
	}

	loggerAdapter := zapadapter.New(logger)
	db = sqldblogger.OpenDriver(dsn, db.Driver(), loggerAdapter)
	dbx := sqlx.NewDb(db, "pgx")

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(render.SetContentType(render.ContentTypeJSON))

	userRepo := user.NewRepository(dbx)
	userMiddleware := user.NewMiddleware(&userRepo)
	userHandler := user.NewHandler(&userRepo)

	todoListRepo := todo_list.NewRepository(dbx)
	todoListMiddleware := todo_list.NewMiddleware(&todoListRepo)
	todoListHandler := todo_list.NewHandler(&todoListRepo)

	todoRepo := todo.NewRepository(dbx)
	todoHandler := todo.NewHandler(&todoRepo)

	router.Get("/status", func(writer http.ResponseWriter, req *http.Request) {
		writer.Write([]byte("ok"))
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

	println("Server starting!")

	http.ListenAndServe(":8080", router)
}
