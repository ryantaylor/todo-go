package db

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"todo/models"
)

type Record interface {
	models.User | models.TodoList | models.Todo
}

type CreateRecordDB interface {
	QueryRowx(query string, args ...interface{}) *sqlx.Row
}

func CreateRecord[T Record](db CreateRecordDB, record *T, tableName string, attributes sq.Eq) error {
	query, args, err := Builder.Insert(tableName).SetMap(attributes).Suffix(ReturningAll).ToSql()
	if err != nil {
		return err
	}

	return db.QueryRowx(query, args...).StructScan(record)
}

type FindRecordDB interface {
	Get(dest interface{}, query string, args ...interface{}) error
}

func FindRecord[T Record](db FindRecordDB, record *T, tableName string, where sq.Eq) error {
	query, args, err := Builder.Select("*").From(tableName).Where(where).ToSql()
	if err != nil {
		return err
	}

	return db.Get(record, query, args...)
}

func FindRecordByID[T Record](db FindRecordDB, record *T, tableName string, id int) error {
	return FindRecord(db, record, tableName, sq.Eq{"id": id})
}

func ListRecords[T Record](db *sqlx.DB, records *[]T, tableName string, where sq.Eq) error {
	query, args, err := Builder.Select("*").From(tableName).Where(where).ToSql()
	if err != nil {
		return err
	}

	return db.Select(records, query, args...)
}
