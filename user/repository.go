package user

import (
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"todo/db"
)

const (
	tableName = "users"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return Repository{db}
}

func (repo *Repository) CreateUser(email string) (Record, error) {
	record := Record{}

	query, args, err := db.Builder.Insert(tableName).SetMap(squirrel.Eq{
		"email": email,
	}).Suffix(db.ReturningAll).ToSql()
	if err != nil {
		return record, err
	}

	err = repo.db.QueryRowx(query, args...).StructScan(&record)

	return record, err
}

func (repo *Repository) FindUserByID(id int) (Record, error) {
	record := Record{}

	query, args, err := db.Builder.Select("*").From(tableName).Where(squirrel.Eq{"id": id}).ToSql()
	if err != nil {
		return record, err
	}

	err = repo.db.Get(&record, query, args...)

	return record, err
}
