package internal

import "errors"

// Task is a struct that represents a task
type Task struct {
	// ID is the unique identifier of the task
	ID          int
	// Title is the title of the task
	Title       string
	// Description is the description of the task
	Description string
	// Done is a flag that indicates if the task is done
	Done        bool
}

var (
	// ErrTaskNotFound is an error that is used when the task is not found
	ErrTaskNotFound = errors.New("task not found")
	// ErrTaskDuplicate is an error that is used when the task already exists
	ErrTaskDuplicated = errors.New("task already exists")
	// ErrTaskInvalidField is an error that is used when the task is invalid
	ErrTaskInvalidField = errors.New("task is invalid")
	// ErrTaskInternal is an error that is used when the task can't be saved
	ErrTaskInternal = errors.New("task can't be processed")
)

// TaskRepository is an interface that represents a task repository
type TaskRepository interface {
	// Save saves a task
	Save(task *Task) (err error)
}

// TaskService is an interface that represents a task service
type TaskService interface {
	// Save saves a task
	Save(task *Task) (err error)
}