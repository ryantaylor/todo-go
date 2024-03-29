package todo

import (
	sq "github.com/Masterminds/squirrel"
	"todo/db"
	"todo/models"
)

type RepositoryDB interface {
	db.CreateRecordDB
	db.ListRecordDB
}

type Repository struct {
	db RepositoryDB
}

func NewRepository(db RepositoryDB) Repository {
	return Repository{db}
}

func (repo *Repository) CreateTodo(todoListID int, text string) (models.Todo, error) {
	record := models.Todo{}
	attributes := sq.Eq{
		"todo_list_id": todoListID,
		"text":         text,
	}

	err := db.CreateRecord(repo.db, &record, db.TableTodos, attributes)

	return record, err
}

func (repo *Repository) ListTodos(where sq.Eq) ([]models.Todo, error) {
	var records []models.Todo

	err := db.ListRecords(repo.db, &records, db.TableTodos, where)

	return records, err
}
