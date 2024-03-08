package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-web/methods/app/internal"
	"go-web/methods/app/platform/tools"
	"io"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
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
	// vl internal.TaskValidator
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

// Update is a method that updates a task
func (d *DefaultTask) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - id from path
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid id")
			return
		}

		// - read into bytes
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid request body")
			return
		}
		// - parse to map (dynamic)
		bodyMap := map[string]any{}
		if err := json.Unmarshal(bytes, &bodyMap); err != nil {
			response.Text(w, http.StatusBadRequest, "invalid request body")
			return
		}
		// validate required fields
		if err := tools.CheckFieldExistance(bodyMap, "title", "description", "done"); err != nil {
			var fieldError *tools.FieldError
			if errors.As(err, &fieldError) {
				response.Text(w, http.StatusBadRequest, fmt.Sprintf("%s is required", fieldError.Field))
				return
			}

			response.Text(w, http.StatusInternalServerError, "internal server error")
			return
		}
		// - parse json to struct (static)
		var body TaskRequestBody
		if err := json.Unmarshal(bytes, &body); err != nil {
			response.Text(w, http.StatusBadRequest, "invalid request body")
			return
		}
		// - validate the task
		if body.Title == "" || len(body.Title) > 25 {
			response.Text(w, http.StatusBadRequest, "title is required and must be less than 25 characters")
			return
		}
		if len(body.Description) > 50 {
			response.Text(w, http.StatusBadRequest, "description must be less than 250 characters")
			return
		}

		// process
		// - serialize the request body into a task
		task := internal.Task{
			ID:          id,
			Title:       body.Title,
			Description: body.Description,
			Done:        body.Done,
		}
		// - update the task
		if err := d.sv.Update(task);err != nil {
			switch {
			case errors.Is(err, internal.ErrTaskNotFound):
				response.Text(w, http.StatusNotFound, "task not found")
			case errors.Is(err, internal.ErrTaskInvalidField):
				response.Text(w, http.StatusBadRequest, "task is invalid")
			case errors.Is(err, internal.ErrTaskDuplicated):
				response.Text(w, http.StatusConflict, "task already exists")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
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
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "task updated",
			"data":    data,
		})
	}
}

// UpdatePartial is a method that updates a task partially
func (d *DefaultTask) UpdatePartial() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - id from path
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid id")
			return
		}
		// - parse map from json
		bodyMap := make(map[string]any)
		if err := request.JSON(r, &bodyMap); err != nil {
			response.Text(w, http.StatusBadRequest, "invalid request body")
			return
		}
		// validate required fields
		if title, ok := bodyMap["title"]; ok {
			titleString, ok := title.(string)
			if !ok {
				response.Text(w, http.StatusBadRequest, "title is invalid")
				return
			}

			if titleString == "" || len(titleString) > 25 {
				response.Text(w, http.StatusBadRequest, "title is required and must be less than 25 characters")
				return
			}
		}
		if description, ok := bodyMap["description"]; ok {
			descriptionString, ok := description.(string)
			if !ok {
				response.Text(w, http.StatusBadRequest, "description is invalid")
				return
			}

			if len(descriptionString) > 50 {
				response.Text(w, http.StatusBadRequest, "description must be less than 250 characters")
				return
			}
		}

		// process
		// - update the task partially
		if err := d.sv.UpdatePartial(id, bodyMap); err != nil {
			switch {
			case errors.Is(err, internal.ErrTaskNotFound):
				response.Text(w, http.StatusNotFound, "task not found")
			case errors.Is(err, internal.ErrTaskInvalidField):
				response.Text(w, http.StatusBadRequest, "task is invalid")
			case errors.Is(err, internal.ErrTaskDuplicated):
				response.Text(w, http.StatusConflict, "task already exists")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		// response
		response.Text(w, http.StatusOK, "task updated")
	}
}

// Delete is a method that deletes a task
func (d *DefaultTask) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - id from path
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid id")
			return
		}

		// process
		if err := d.sv.Delete(id); err != nil {
			switch {
			case errors.Is(err, internal.ErrTaskNotFound):
				response.Text(w, http.StatusNotFound, "task not found")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		// response
		w.WriteHeader(http.StatusNoContent)
	}
}