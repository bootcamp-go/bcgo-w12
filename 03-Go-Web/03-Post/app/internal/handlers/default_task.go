package handlers

import (
	"encoding/json"
	"go-web/post/app/internal"
	"go-web/post/app/platform/web"
	"io"
	"net/http"
)

func NewDefaultTask(tasks map[int]internal.Task, lastID int) *DefaultTask {
	// default values
	defaultTasks := make(map[int]internal.Task)
	defaultLastID := 0
	if tasks != nil {
		defaultTasks = tasks
	}
	if lastID != 0 {
		defaultLastID = lastID
	}

	return &DefaultTask{
		tasks:  defaultTasks,
		lastID: defaultLastID,
	}
}

type TaskJSON struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type TaskRequestBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Done		bool   `json:"done"`
}

// DefaultTask is a struct that contains the handlers for the tasks
type DefaultTask struct {
	// tasks is a map that stores the tasks
	// - key: id
	// - value: internal.Task (structured data)
	tasks map[int]internal.Task
	// lastID is the last identifier used
	lastID int
}

// Create is a method that creates a new task
func (d *DefaultTask) Create() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		// validate token
		token := r.Header.Get("Authorization")
		if token != "12345" {
			web.ResponseJSON(w, http.StatusUnauthorized, map[string]any{"error": "unauthorized"})
			return
		}

		// request
		// - read into bytes
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			web.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "invalid request body"})
			return
		}

		// - parse to map (dynamic)
		bodyMap := map[string]any{}
		if err := json.Unmarshal(bytes, &bodyMap); err != nil {
			web.ResponseJSON(w, http.StatusBadRequest, map[string]any{
				"error": "invalid request body",
			})
			return
		}
		// - validate required keys
		if _, ok := bodyMap["title"]; !ok {
			web.ResponseJSON(w, http.StatusBadRequest, map[string]any{
				"error": "title is required",
			})
			return
		}
		if _, ok := bodyMap["description"]; !ok {
			web.ResponseJSON(w, http.StatusBadRequest, map[string]any{
				"error": "description is required",
			})
			return
		}
		if _, ok := bodyMap["done"]; !ok {
			web.ResponseJSON(w, http.StatusBadRequest, map[string]any{
				"error": "done is required",
			})
			return
		}

		// - parse json to struct (static)	
		var body TaskRequestBody
		if err := json.Unmarshal(bytes, &body); err != nil {
			web.ResponseJSON(w, http.StatusBadRequest, map[string]any{
				"error": "invalid request body",
			})
			return
		}

		// if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		// 	web.ResponseJSON(w, http.StatusBadRequest, map[string]any{
		// 		"error": "invalid request body",
		// 	})
		// 	// w.Header().Set("Content-Type", "application/json")
		// 	// w.WriteHeader(http.StatusBadRequest)
		// 	// json.NewEncoder(w).Encode(map[string]any{
		// 	// 	"error": "invalid request body",
		// 	// })
		// 	return
		// }

		// process
		// - increment the last identifier
		d.lastID++
		// - serialize the request body into a task
		task := internal.Task{
			ID:          d.lastID,
			Title:       body.Title,
			Description: body.Description,
			Done:        body.Done,
		}
		// - validate the task
		if task.Title == "" || len(task.Title) > 25 {
			web.ResponseJSON(w, http.StatusBadRequest, map[string]any{
				"error": "title is required and must be less than 25 characters",
			})
			return
		}
		// - store the task
		d.tasks[task.ID] = task

		// response
		data := TaskJSON{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Done:        task.Done,
		}
		web.ResponseJSON(w, http.StatusCreated, map[string]any{
			"message": "task created",
			"data":    data,
		})
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.WriteHeader(http.StatusCreated)
	// 	json.NewEncoder(w).Encode(map[string]any{
	// 		"message": "task created",
	// 		"data":    data,
	// 	})
	// }
	}
}