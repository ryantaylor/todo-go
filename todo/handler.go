package todo

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/go-chi/render"
	"net/http"
	"todo/models"
)

type HandlerRepository interface {
	CreateTodo(todoListID int, text string) (models.Todo, error)
	ListTodos(where sq.Eq) ([]models.Todo, error)
}

type Handler struct {
	repo HandlerRepository
}

func NewHandler(repo HandlerRepository) Handler {
	return Handler{repo}
}

type CreateRequest struct {
	Text string `json:"text"`
}

func (data *CreateRequest) Bind(req *http.Request) error {
	return nil
}

type Response struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

func NewResponse(record *models.Todo) *Response {
	return &Response{record.ID, record.Text}
}

func NewListResponse(records []models.Todo) []render.Renderer {
	var list []render.Renderer
	for _, record := range records {
		list = append(list, NewResponse(&record))
	}
	return list
}

func (response *Response) Render(writer http.ResponseWriter, req *http.Request) error {
	return nil
}

func (handler *Handler) CreateForTodoList(writer http.ResponseWriter, req *http.Request) {
	todoList := req.Context().Value("todoList").(models.TodoList)
	input := CreateRequest{}

	err := render.Bind(req, &input)
	if err != nil {
		panic(err)
	}

	record, err := handler.repo.CreateTodo(todoList.ID, input.Text)
	if err != nil {
		panic(err)
	}

	render.Status(req, http.StatusCreated)
	err = render.Render(writer, req, NewResponse(&record))
	if err != nil {
		panic(err)
	}
}

func (handler *Handler) ListForTodoList(writer http.ResponseWriter, req *http.Request) {
	todoList := req.Context().Value("todoList").(models.TodoList)

	todos, err := handler.repo.ListTodos(sq.Eq{"todo_list_id": todoList.ID})
	if err != nil {
		panic(err)
	}

	err = render.RenderList(writer, req, NewListResponse(todos))
	if err != nil {
		panic(err)
	}
}
