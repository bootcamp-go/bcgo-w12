package service

import "go-web/structure/app/internal"

// NewDefaultTask creates a new TaskDefault
func NewDefaultTask(rp internal.TaskRepository) *TaskDefault {
	// create the task default
	return &TaskDefault{
		rp: rp,
	}
}

// TaskDefault is an implementation of the TaskService interface
type TaskDefault struct {
	// repository is the task repository
	rp internal.TaskRepository
	// external services / apis
	// ...
}

func (t *TaskDefault) Save(task *internal.Task) (err error) {
	// save the task
	err = t.rp.Save(task)
	return
}