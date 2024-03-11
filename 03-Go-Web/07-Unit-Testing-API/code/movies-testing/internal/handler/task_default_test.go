package handler

import (
	"context"
	"encoding/json"
	"go-web/unit-testing/app/internal"
	"go-web/unit-testing/app/internal/repository"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

func TestTaskDefault_GetTask_HandlerFunc(t *testing.T) {
	t.Run("success - get task", func(t *testing.T) {
		// arrange
		// - repository: map
		db := map[int]internal.Task{
			1: {
				ID:          1,
				Name:        "task 1",
				Description: "description 1",
				Completed:   false,
			},
		}
		rp := repository.NewTaskMap(db)
		// - handler
		hd := NewTaskDefault(rp)
		hdFunc := hd.GetTask()

		// act
		req := httptest.NewRequest("GET", "/tasks/1", nil)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		res := httptest.NewRecorder()
		hdFunc(res, req)

		// assert
		expectedTask, err := json.Marshal(TaskJSON{
			ID:          1,
			Name:        "task 1",
			Description: "description 1",
			Completed:   false,
		})
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, res.Code)
		require.Equal(t, expectedTask, res.Body.Bytes())
		// require.JSONEq(t, `{
		// 	"id": 1,
		// 	"name": "task 1",
		// 	"description": "description 1",
		// 	"completed": false
		// }`, res.Body.String())
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
	})

	t.Run("fail - invalid id", func(t *testing.T) {
		// arrange
		// - repository: map
		rp := repository.NewTaskMap(nil)
		// - handler
		hd := NewTaskDefault(rp)
		hdFunc := hd.GetTask()

		// act
		req := httptest.NewRequest("GET", "/tasks/invalid", nil)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "invalid")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		res := httptest.NewRecorder()
		hdFunc(res, req)

		// assert
		require.Equal(t, http.StatusBadRequest, res.Code)
		require.Equal(t, "invalid id", res.Body.String())
		require.Equal(t, "text/plain; charset=utf-8", res.Header().Get("Content-Type"))
	})

	t.Run("fail - task not found", func(t *testing.T) {
		// arrange

		// act

		// assert

	})
}

func TestTaskDefault_GetTask_Chi(t *testing.T) {
	t.Run("success - get task", func(t *testing.T) {
		// arrange
		// - repository: map
		db := map[int]internal.Task{
			1: {
				ID:          1,
				Name:        "task 1",
				Description: "description 1",
				Completed:   false,
			},
		}
		rp := repository.NewTaskMap(db)
		// - handler
		hd := NewTaskDefault(rp)
		hdFunc := hd.GetTask()
		// - router
		rt := chi.NewRouter()
		rt.Get("/tasks/{id}", hdFunc)

		// act
		req := httptest.NewRequest("GET", "/tasks/1", nil)
		res := httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// assert
		require.Equal(t, http.StatusOK, res.Code)
		require.JSONEq(t, `{
			"id": 1,
			"name": "task 1",
			"description": "description 1",
			"completed": false
		}`, res.Body.String())
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
	})
}