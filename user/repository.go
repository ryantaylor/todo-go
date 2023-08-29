package user

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

func (repo *Repository) CreateUser(email string) (models.User, error) {
	record := models.User{}
	attributes := sq.Eq{"email": email}

	err := db.CreateRecord(repo.db, &record, db.TableUsers, attributes)

	return record, err
}

func (repo *Repository) FindUserByID(id int) (models.User, error) {
	record := models.User{}

	err := db.FindRecordByID(repo.db, &record, db.TableUsers, id)

	return record, err
}
