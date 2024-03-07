package repository

import "go-web/structure/app/internal"

// NewTaskMap creates a new TaskMap
func NewTaskMap(db map[int]internal.Task, lastID int) *TaskMap {
	// default values
	if db == nil {
		db = make(map[int]internal.Task)
	}

	// create the task map
	return &TaskMap{
		db:     db,
		lastID: lastID,
	}
}

// TaskMap is an implementation of the TaskRepository interface
// based on a map
type TaskMap struct {
	// db is a map that stores the tasks
	// - key: id
	// - value: internal.Task (structured data)
	db map[int]internal.Task

	// lastID is the last identifier used
	lastID int
}

func (t *TaskMap) Save(task *internal.Task) (err error) {
	// check if the task already exists
	for _, t := range (*t).db {
		if t.Title == (*task).Title {
			err = internal.ErrTaskDuplicated
			return
		}
	}

	// increment lastID
	(*t).lastID++

	// set the id of the task
	(*task).ID = (*t).lastID

	// save the task
	(*t).db[(*task).ID] = *task
	
	return
}