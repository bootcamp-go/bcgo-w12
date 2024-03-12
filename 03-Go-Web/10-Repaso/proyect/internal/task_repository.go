package internal

type TaskRepository interface {
	FindAll() ([]Task, error)
	FindByID(id int) (Task, error)
}