package infra

import (
	"errors"
	"github.com/go-chi/render"
	"net/http"
	"todo/db"
)

type ErrResponse struct {
	Err            error `json:"-"`
	HTTPStatusCode int   `json:"-"`

	Message string `json:"message"`
}

func (res *ErrResponse) Render(writer http.ResponseWriter, req *http.Request) error {
	render.Status(req, res.HTTPStatusCode)
	return nil
}

func ErrorResponse(err error) render.Renderer {
	var duplicate *db.DuplicateError
	if errors.As(err, &duplicate) {
		return &ErrResponse{
			Err:            err,
			HTTPStatusCode: http.StatusUnprocessableEntity,
			Message:        err.Error(),
		}
	} else {
		return &ErrResponse{
			Err:            err,
			HTTPStatusCode: http.StatusInternalServerError,
			Message:        "Something went wrong!",
		}
	}
}
