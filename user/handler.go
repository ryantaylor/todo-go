package user

import (
	"github.com/go-chi/render"
	"net/http"
	"todo/infra"
	"todo/models"
)

type HandlerRepository interface {
	CreateUser(email string) (*models.User, error)
	FindUserByID(id int) (*models.User, error)
}

type Handler struct {
	repo HandlerRepository
}

func NewHandler(repo HandlerRepository) Handler {
	return Handler{repo}
}

type CreateRequest struct {
	Email string `json:"email"`
}

func (data *CreateRequest) Bind(_ *http.Request) error {
	return nil
}

type Response struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

func NewResponse(record *models.User) *Response {
	return &Response{record.ID, record.Email}
}

func (response *Response) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (handler *Handler) Create(writer http.ResponseWriter, req *http.Request) {
	input := CreateRequest{}

	err := render.Bind(req, &input)
	if err != nil {
		panic(err)
	}

	record, err := handler.repo.CreateUser(input.Email)
	if err != nil {
		_ = render.Render(writer, req, infra.ErrorResponse(err))
		return
	}

	render.Status(req, http.StatusCreated)
	err = render.Render(writer, req, NewResponse(record))
	if err != nil {
		panic(err)
	}
}

func (handler *Handler) Find(writer http.ResponseWriter, req *http.Request) {
	user := req.Context().Value("user").(models.User)

	err := render.Render(writer, req, NewResponse(&user))
	if err != nil {
		panic(err)
	}
}
