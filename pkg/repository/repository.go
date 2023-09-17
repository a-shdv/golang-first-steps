package repository

type Repository struct {
	Authorization
	TodoItem
	TodoList
}

func NewRepository() *Repository {
	return &Repository{}
}
