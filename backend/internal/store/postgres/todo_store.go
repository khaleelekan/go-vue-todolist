type TodoStore struct {
	db *sql.DB
}

func NewTodoStore(db *sql.DB) *TodoStore {
	return &TodoStore{db: db}
}
func (s *TodoStore) Create(todo *domain.Todo) (*domain.Todo, error) {
	query := `
	INSERT INTO todos (title, completed, user_id)
	VALUES ($1, $2, $3)
	RETURNING id, created_at;
	`
	err := s.db.QueryRow(
		query,
		todo.Title,
		todo.Completed,
		todo.UserID,
	).Scan(&todo.ID, &todo.CreatedAt)

	return todo, err
}

CREATE TABLE todos (
	id SERIAL PRIMARY KEY,
	title TEXT NOT NULL,
	completed BOOLEAN DEFAULT false,
	user_id INT NOT NULL,
	created_at TIMESTAMP DEFAULT now()
);
