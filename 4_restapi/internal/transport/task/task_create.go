package transport_task

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/notIdealen/aihwbe.git/internal/domain"
	transport_error "github.com/notIdealen/aihwbe.git/internal/transport/error"
)

type CreateTaskRequestDTO struct {
	Title  string `json:"title"`
	Desc   string `json:"description"`
	Status string `json:"status"`
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	mrw, ok := w.(*domain.MyResponseWriter)
	if !ok {
		panic("Bad type of response writer")
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		mrw.StatusCode = http.StatusBadRequest
		mrw.Log = fmt.Sprintf("TaskHandler: CreateTask: ReadAll: %s", err)
		http.Error(mrw, "invalid body", mrw.StatusCode)
		return
	}
	defer r.Body.Close()

	var dto CreateTaskRequestDTO
	err = json.Unmarshal(body, &dto)
	if err != nil {
		mrw.StatusCode = http.StatusBadRequest
		mrw.Log = fmt.Sprintf("TaskHandler: CreateTask: Unmarshal: %s", err)
		http.Error(mrw, "invalid body", mrw.StatusCode)
		return
	}

	id, err := h.s.CreateTask(r.Context(), dto)
	if err != nil {
		re := transport_error.ConcertToResponseError(err)
		mrw.StatusCode = re.StatusCode
		mrw.Log = re.Log
		http.Error(mrw, re.Message, re.StatusCode)
		return
	}

	mrw.Header().Set("Content-Type", "application/json")
	mrw.StatusCode = http.StatusOK
	mrw.Write([]byte(id))
}
