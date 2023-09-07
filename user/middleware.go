package user

import (
	"context"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
	"todo/models"
)

type MiddlewareRepository interface {
	FindUserByID(id int) (*models.User, error)
}

type Middleware struct {
	repo MiddlewareRepository
}

func NewMiddleware(repo MiddlewareRepository) Middleware {
	return Middleware{repo}
}

func (step *Middleware) Context(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		var user *models.User
		var err error

		if userID, err := strconv.Atoi(chi.URLParam(req, "userID")); userID != 0 {
			user, err = step.repo.FindUserByID(userID)
		} else {
			panic(err)
		}
		if err != nil {
			panic(err)
		}

		ctx := context.WithValue(req.Context(), "user", user)
		next.ServeHTTP(writer, req.WithContext(ctx))
	})
}
