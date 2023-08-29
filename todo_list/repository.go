package todo_list

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"todo/db"
	"todo/models"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return Repository{db}
}

func (repo *Repository) CreateTodoList(userID int, name string) (models.TodoList, error) {
	record := models.TodoList{}
	attributes := sq.Eq{
		"user_id": userID,
		"name":    name,
	}

	err := db.CreateRecord(repo.db, &record, db.TableTodoLists, attributes)

	return record, err
}

func (repo *Repository) FindTodoListByID(id int, includeTodos bool) (models.TodoList, error) {
	record := models.TodoList{}

	err := db.FindRecordByID(repo.db, &record, db.TableTodoLists, id)

	if includeTodos {
		err = db.ListRecords(repo.db, &record.Todos, db.TableTodos, sq.Eq{"todo_list_id": record.ID})
	}

	return record, err
}

func (repo *Repository) FindTodoList(where sq.Eq, includeTodos bool) (models.TodoList, error) {
	record := models.TodoList{}

	err := db.FindRecord(repo.db, &record, db.TableTodoLists, where)

	if includeTodos {
		err = db.ListRecords(repo.db, &record.Todos, db.TableTodos, sq.Eq{"todo_list_id": record.ID})
	}

	return record, err
}

func (repo *Repository) ListTodoLists(where sq.Eq) ([]models.TodoList, error) {
	var records []models.TodoList

	err := db.ListRecords(repo.db, &records, db.TableTodoLists, where)

	return records, err
}
