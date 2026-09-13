package transport

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/notIdealen/aihwbe.git/internal/domain"
	"github.com/notIdealen/aihwbe.git/internal/middleware"
	"github.com/notIdealen/aihwbe.git/internal/server"
)

type Servicer interface {
	GetTask(context.Context, string) (*domain.Task, error)
	CreateTask(context.Context, CreateTaskRequestDTO) (string, error)
	DeleteTask(context.Context, string) error
	UpdateTask(context.Context, string, CreateTaskRequestDTO) (*domain.Task, error)
	GetList(context.Context) ([]domain.Task, error)
}

type TaskHandler struct {
	s Servicer
}

func NewTaskHandler(s Servicer) *TaskHandler {
	return &TaskHandler{
		s: s,
	}
}

func (h *TaskHandler) GetTask(rw http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	w, ok := rw.(*middleware.MyResponseWriter)
	if !ok {
		panic("Bad type of response writer")
	}

	//проверить id на правильность написания
	task, err := h.s.GetTask(r.Context(), id)
	if err != nil {
		writeResponseError(w, err)
		return
	}
	resp, err := json.Marshal(task)
	if err != nil {
		w.SetStatusCode(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.SetStatusCode(http.StatusOK)
	// w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

type CreateTaskRequestDTO struct {
	Title  string            `json:"title"`
	Desc   string            `json:"description"`
	Status domain.TaskStatus `json:"status"`
}

func (h *TaskHandler) CreateTask(rw http.ResponseWriter, r *http.Request) {
	var dto CreateTaskRequestDTO
	//проверить контент тайп, чтобы был json
	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	w, ok := rw.(*middleware.MyResponseWriter)
	if !ok {
		panic("Bad type of response writer")
	}

	err := json.Unmarshal(body, &dto)
	if err != nil {
		w.SetStatusCode(http.StatusBadRequest)
		w.Write([]byte(http.ErrBodyNotAllowed.Error()))
		return
	}

	if dto.Title == "" {
		w.SetStatusCode(http.StatusBadRequest)
		w.Write([]byte("Empty title"))
		return
	}

	if err := dto.Status.IsValidStatus(); err != nil {
		w.SetStatusCode(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	id, err := h.s.CreateTask(r.Context(), dto)
	if err != nil {
		// w.SetStatusCode(http.StatusBadGateway)
		writeResponseError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.SetStatusCode(http.StatusCreated)
	w.Write([]byte(id))
}

func (h *TaskHandler) UpdateTask(rw http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	w, ok := rw.(*middleware.MyResponseWriter)
	if !ok {
		panic("Bad type of response writer")
	}

	var dto CreateTaskRequestDTO
	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	err := json.Unmarshal(body, &dto)
	if err != nil {
		w.SetStatusCode(http.StatusBadRequest)
		w.Write([]byte(http.ErrBodyNotAllowed.Error()))
		return
	}

	if err := dto.Status.IsValidStatus(); err != nil {
		w.SetStatusCode(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	task, err := h.s.UpdateTask(r.Context(), id, dto)
	if err != nil {
		writeResponseError(w, err)
		return
	}
	resp, err := json.Marshal(task)
	if err != nil {
		w.SetStatusCode(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.SetStatusCode(http.StatusOK)
	w.Write(resp)
}

func (h *TaskHandler) DeleteTask(rw http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	w, ok := rw.(*middleware.MyResponseWriter)
	if !ok {
		panic("Bad type of response writer")
	}

	err := h.s.DeleteTask(r.Context(), id)
	if err != nil {
		writeResponseError(w, err)
		return
	}
	w.SetStatusCode(http.StatusNoContent)
}

func (h *TaskHandler) GetList(rw http.ResponseWriter, r *http.Request) {

	w, ok := rw.(*middleware.MyResponseWriter)
	if !ok {
		panic("Bad type of response writer")
	}

	tasks, err := h.s.GetList(r.Context())
	if err != nil {
		writeResponseError(w, err)
		return
	}

	var resp bytes.Buffer

	enc := json.NewEncoder(&resp)
	for _, task := range tasks {
		err = enc.Encode(task)
		if err != nil {
			writeResponseError(w, err)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.SetStatusCode(http.StatusOK)
	w.Write(resp.Bytes())
}

func (h *TaskHandler) GetRoutes() []server.Route {

	return []server.Route{
		{
			Method:  http.MethodGet,
			Path:    "/tasks/{id}",
			Handler: h.GetTask,
		},
		{
			Method:  http.MethodGet,
			Path:    "/tasks",
			Handler: h.GetList,
		},
		{
			Method:  http.MethodPost,
			Path:    "/tasks",
			Handler: h.CreateTask,
		},
		{
			Method:  http.MethodPut,
			Path:    "/tasks/{id}",
			Handler: h.UpdateTask,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/tasks/{id}",
			Handler: h.DeleteTask,
		},
	}
}

// func (h *TaskHandler) writeError(err error) domain.ResponseError {
// 	if errors.Is()

// }
