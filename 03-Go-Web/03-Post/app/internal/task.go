package internal

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