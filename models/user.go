package models

type User struct {
	ID        int
	Email     string
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}
