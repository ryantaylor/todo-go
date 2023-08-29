package user

import (
	sq "github.com/Masterminds/squirrel"
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
