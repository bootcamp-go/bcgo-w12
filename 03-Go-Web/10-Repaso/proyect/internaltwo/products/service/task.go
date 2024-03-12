package service

type Task struct {
	ID    int
	Name  string
	Price float64
}

type TaskService interface {
	FindAll() ([]Task, error)
}