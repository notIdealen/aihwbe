package transport_task

import "net/http"

func (h *TaskHandler) DeleteTask(rw http.ResponseWriter, r *http.Request) {
	// id := r.PathValue("id")

	// w, ok := rw.(*middleware.MyResponseWriter)
	// if !ok {
	// 	panic("Bad type of response writer")
	// }

	// err := h.s.DeleteTask(r.Context(), id)
	// if err != nil {
	// 	writeResponseError(w, err)
	// 	return
	// }
	// w.SetStatusCode(http.StatusNoContent)
}
