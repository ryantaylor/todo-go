package user

import (
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"todo/db"
	"todo/models"
)

type RepositoryDB interface {
	db.CreateRecordDB
	db.FindRecordDB
}

type Repository struct {
	db RepositoryDB
}

func NewRepository(db RepositoryDB) Repository {
	return Repository{db}
}

func (repo *Repository) CreateUser(email string) (*models.User, error) {
	record := models.User{}
	attributes := sq.Eq{"email": email}

	err := db.CreateRecord(repo.db, &record, db.TableUsers, attributes)
	if err != nil {
		var e *pgconn.PgError
		if errors.As(err, &e) && e.Code == pgerrcode.UniqueViolation {
			return nil, &db.DuplicateError{Message: fmt.Sprintf("A user with email %v already exists!", email)}
		}
		return nil, err
	}

	return &record, err
}

func (repo *Repository) FindUserByID(id int) (*models.User, error) {
	record := models.User{}

	err := db.FindRecordByID(repo.db, &record, db.TableUsers, id)
	if err != nil {
		return nil, err
	}

	return &record, err
}
