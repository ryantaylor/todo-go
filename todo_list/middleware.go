package todo_list

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
	"todo/models"
)

type MiddlewareRepository interface {
	FindTodoList(where sq.Eq, includeTodos bool) (models.TodoList, error)
}

type Middleware struct {
	repo MiddlewareRepository
}

func NewMiddleware(repo MiddlewareRepository) Middleware {
	return Middleware{repo}
}

func (step *Middleware) Context(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		var todoList models.TodoList
		var err error
		userRecord := req.Context().Value("user").(models.User)

		if todoListID, err := strconv.Atoi(chi.URLParam(req, "todoListID")); todoListID != 0 {
			todoList, err = step.repo.FindTodoList(sq.Eq{"id": todoListID, "user_id": userRecord.ID}, false)
		} else {
			panic(err)
		}
		if err != nil {
			panic(err)
		}

		ctx := context.WithValue(req.Context(), "todoList", todoList)
		next.ServeHTTP(writer, req.WithContext(ctx))
	})
}
