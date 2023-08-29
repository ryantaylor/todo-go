package db

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"todo/models"
)

type Record interface {
	models.User | models.TodoList | models.Todo
}

func CreateRecord[T Record](db *sqlx.DB, record *T, tableName string, attributes sq.Eq) error {
	query, args, err := Builder.Insert(tableName).SetMap(attributes).Suffix(ReturningAll).ToSql()
	if err != nil {
		return err
	}

	return db.QueryRowx(query, args...).StructScan(record)
}

func FindRecord[T Record](db *sqlx.DB, record *T, tableName string, where sq.Eq) error {
	query, args, err := Builder.Select("*").From(tableName).Where(where).ToSql()
	if err != nil {
		return err
	}

	return db.Get(record, query, args...)
}

func FindRecordByID[T Record](db *sqlx.DB, record *T, tableName string, id int) error {
	return FindRecord(db, record, tableName, sq.Eq{"id": id})
}

func ListRecords[T Record](db *sqlx.DB, records *[]T, tableName string, where sq.Eq) error {
	query, args, err := Builder.Select("*").From(tableName).Where(where).ToSql()
	if err != nil {
		return err
	}

	return db.Select(records, query, args...)
}
