package repository

import "go-web/unit-testing/app/internal"

// NewTaskMap returns a new TaskMap.
func NewTaskMap(m map[int]internal.Task) *TaskMap {
	// default values
	if m == nil {
		m = make(map[int]internal.Task)
	}

	return &TaskMap{
		m: m,
	}
}

// TaskMap implements TaskRepository using a map.
type TaskMap struct {
	// m is the map that stores tasks.
	// - key: task id
	// - value: task
	m map[int]internal.Task
}

// FindByID returns the task with the given ID.
func (r *TaskMap) FindByID(id int) (task internal.Task, err error) {
	// check if the task exists
	task, ok := r.m[id]
	if !ok {
		err = internal.ErrTaskNotFound
		return
	}

	return
}