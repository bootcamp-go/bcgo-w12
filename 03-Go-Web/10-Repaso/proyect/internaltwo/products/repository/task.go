package repository

type TaskScheme struct {
	ID    int
	Name  string
	Price float64
}

type TaskRepository interface {
	FindAll() ([]TaskScheme, error)
}