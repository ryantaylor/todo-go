package models

type Todo struct {
	ID         int
	TodoListID int `db:"todo_list_id"`
	Text       string
	CreatedAt  string `db:"created_at"`
	UpdatedAt  string `db:"updated_at"`
}
