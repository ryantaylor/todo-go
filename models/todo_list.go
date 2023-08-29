package models

type TodoList struct {
	ID        int
	UserID    int `db:"user_id"`
	Name      string
	Todos     []Todo
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}
