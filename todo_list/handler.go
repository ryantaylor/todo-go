package todo_list

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/go-chi/render"
	"net/http"
	"todo/models"
)

type HandlerRepository interface {
	CreateTodoList(userID int, name string) (models.TodoList, error)
	ListTodoLists(where sq.Eq) ([]models.TodoList, error)
}

type Handler struct {
	repo HandlerRepository
}

func NewHandler(repo HandlerRepository) Handler {
	return Handler{repo}
}

type CreateRequest struct {
	Name string `json:"name"`
}

func (data *CreateRequest) Bind(req *http.Request) error {
	return nil
}

type Response struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func NewResponse(record *models.TodoList) *Response {
	return &Response{record.ID, record.Name}
}

func NewListResponse(records []models.TodoList) []render.Renderer {
	var list []render.Renderer
	for _, record := range records {
		list = append(list, NewResponse(&record))
	}
	return list
}

func (response *Response) Render(writer http.ResponseWriter, req *http.Request) error {
	return nil
}

func (handler *Handler) CreateForUser(writer http.ResponseWriter, req *http.Request) {
	userRecord := req.Context().Value("user").(models.User)
	input := CreateRequest{}

	err := render.Bind(req, &input)
	if err != nil {
		panic(err)
	}

	record, err := handler.repo.CreateTodoList(userRecord.ID, input.Name)
	if err != nil {
		panic(err)
	}

	render.Status(req, http.StatusCreated)
	err = render.Render(writer, req, NewResponse(&record))
	if err != nil {
		panic(err)
	}
}

func (handler *Handler) ListForUser(writer http.ResponseWriter, req *http.Request) {
	userRecord := req.Context().Value("user").(models.User)

	todoLists, err := handler.repo.ListTodoLists(sq.Eq{"user_id": userRecord.ID})
	if err != nil {
		panic(err)
	}

	err = render.RenderList(writer, req, NewListResponse(todoLists))
	if err != nil {
		panic(err)
	}
}

func (handler *Handler) FindForUser(writer http.ResponseWriter, req *http.Request) {
	todoList := req.Context().Value("todoList").(models.TodoList)

	err := render.Render(writer, req, NewResponse(&todoList))
	if err != nil {
		panic(err)
	}
}
