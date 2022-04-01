package dtos

type Todo struct {
	ID     string `db:"id"`
	Text   string `db:"todo_text"`
	Done   bool   `db:"done"`
	UserID string `db:"user_id"`
}
