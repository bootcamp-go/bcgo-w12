package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-web/structure/app/internal"
	"go-web/structure/app/platform/tools"
	"io"
	"net/http"

	"github.com/bootcamp-go/web/response"
)

func NewDefaultTask(sv internal.TaskService) *DefaultTask {
	return &DefaultTask{
		sv: sv,
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
	sv internal.TaskService
}


// Create is a method that creates a new task
func (d *DefaultTask) Create() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		// request
		// - read into bytes
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{"message": "invalid request body"})
			return
		}
		// - parse to map (dynamic)
		bodyMap := map[string]any{}
		if err := json.Unmarshal(bytes, &bodyMap); err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "invalid request body",
			})
			return
		}
		// validate required fields
		if err := tools.CheckFieldExistance(bodyMap, "title", "description", "done"); err != nil {
			var fieldError *tools.FieldError
			if errors.As(err, &fieldError) {
				response.JSON(w, http.StatusBadRequest, map[string]any{"message": fmt.Sprintf("%s is required", fieldError.Field)})
				return
			}

			response.JSON(w, http.StatusInternalServerError, map[string]any{"message": "internal server error"})
			return
		}
		// - parse json to struct (static)	
		var body TaskRequestBody
		if err := json.Unmarshal(bytes, &body); err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "invalid request body",
			})
			return
		}
		// - validate the task
		if body.Title == "" || len(body.Title) > 25 {
			response.JSON(w, http.StatusBadRequest, map[string]any{"message": "title is required and must be less than 25 characters"})
			return
		}
		if len(body.Description) > 50 {
			response.JSON(w, http.StatusBadRequest, map[string]any{"message": "description must be less than 250 characters"})
			return
		}

		// process
		// - serialize the request body into a task
		task := internal.Task{
			Title:       body.Title,
			Description: body.Description,
			Done:        body.Done,
		}
		// - save the task
		if err := d.sv.Save(&task); err != nil {
			switch {
			case errors.Is(err, internal.ErrTaskDuplicated):
				response.JSON(w, http.StatusConflict, map[string]any{"message": "task already exists"})
			case errors.Is(err, internal.ErrTaskInvalidField):
				response.JSON(w, http.StatusBadRequest, map[string]any{"message": "task is invalid"})
			default:
				response.JSON(w, http.StatusInternalServerError, map[string]any{"message": "internal server error"})
			}
			return
		}

		// response
		data := TaskJSON{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Done:        task.Done,
		}
		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "task created",
			"data":    data,
		})
	}
}