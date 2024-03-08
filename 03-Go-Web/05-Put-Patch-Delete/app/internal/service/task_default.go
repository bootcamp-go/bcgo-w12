package service

import "go-web/methods/app/internal"

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

func (t *TaskDefault) Update(task internal.Task) (err error) {
	// update the task
	err = t.rp.Update(task)
	return
}

func (t *TaskDefault) UpdatePartial(id int, fields map[string]interface{}) (err error) {
	// update the task partially
	err = t.rp.UpdatePartial(id, fields)
	return
}

func (t *TaskDefault) Delete(id int) (err error) {
	// delete the task
	err = t.rp.Delete(id)
	return
}
