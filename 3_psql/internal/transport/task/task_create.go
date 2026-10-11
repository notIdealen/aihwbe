package transport_task

import (
	"encoding/json"
	"io"
	"net/http"

	middleware_router "github.com/notIdealen/aihwbe.git/internal/middleware/router"
	"github.com/notIdealen/aihwbe.git/internal/transport"
)

type CreateTaskRequestDTO struct {
	Title  string `json:"title"`
	Desc   string `json:"description"`
	Status string `json:"status"`
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	mrw, ok := w.(*middleware_router.MyResponseWriter)
	if !ok {
		panic("Bad type of response writer")
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		mrw.WriteHeader(http.StatusBadRequest)
		mrw.Write([]byte("invalid body"))
		return
	}
	defer r.Body.Close()

	var dto CreateTaskRequestDTO
	err = json.Unmarshal(body, &dto)
	if err != nil {
		mrw.WriteHeader(http.StatusBadRequest)
		mrw.Write([]byte("invalid body"))
		return
	}

	id, err := h.s.CreateTask(r.Context(), dto)
	if err != nil {
		clientErr := transport.ParseResponseError(err, mrw)
		mrw.Write([]byte(clientErr))
		return
	}

	mrw.Header().Set("Content-Type", "application/json")
	mrw.StatusCode = http.StatusOK
	w.Write([]byte(id))
}
