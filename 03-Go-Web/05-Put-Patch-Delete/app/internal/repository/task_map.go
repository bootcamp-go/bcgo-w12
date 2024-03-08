package repository

import "go-web/methods/app/internal"

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

func (t *TaskMap) Update(task internal.Task) (err error) {
	// check if the task exists
	_, ok := (*t).db[task.ID]
	if !ok {
		err = internal.ErrTaskNotFound
		return
	}

	// check if the task already exists
	for _, t := range (*t).db {
		if t.ID != task.ID && t.Title == task.Title {
			err = internal.ErrTaskDuplicated
			return	
		}
	}

	// update the task
	(*t).db[task.ID] = task
	return
}

func (t *TaskMap) UpdatePartial(id int, fields map[string]any) (err error) {
	// check if the task exists
	task, ok := (*t).db[id]
	if !ok {
		err = internal.ErrTaskNotFound
		return
	}

	// iterate over the fields
	for key, value := range fields {
		switch key {
		case "Title", "title":
			title, ok := value.(string)
			if !ok {
				err = internal.ErrTaskInvalidField
				return
			}
			// check if the task already exists
			for _, t := range (*t).db {
				if t.ID != id && t.Title == title {
					err = internal.ErrTaskDuplicated
					return
				}
			}
			task.Title = title
		case "Description", "description":
			task.Description, ok = value.(string)
			if !ok {
				err = internal.ErrTaskInvalidField
				return
			}
		case "Done", "done":
			task.Done, ok = value.(bool)
			if !ok {
				err = internal.ErrTaskInvalidField
				return
			}
		default:
		}
	}

	// update the task
	(*t).db[id] = task
	return
}

func (t *TaskMap) Delete(id int) (err error) {	
	// check if the task exists
	_, ok := (*t).db[id]
	if !ok {
		err = internal.ErrTaskNotFound
		return
	}

	// delete the task
	delete((*t).db, id)
	return
}