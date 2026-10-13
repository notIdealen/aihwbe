package transport_task

import "net/http"

func (h *TaskHandler) UpdateTask(rw http.ResponseWriter, r *http.Request) {
	// id := r.PathValue("id")

	// w, ok := rw.(*middleware.MyResponseWriter)
	// if !ok {
	// 	panic("Bad type of response writer")
	// }

	// var dto CreateTaskRequestDTO
	// body, _ := io.ReadAll(r.Body)
	// defer r.Body.Close()

	// err := json.Unmarshal(body, &dto)
	// if err != nil {
	// 	w.SetStatusCode(http.StatusBadRequest)
	// 	w.Write([]byte(http.ErrBodyNotAllowed.Error()))
	// 	return
	// }

	// if err := dto.Status.IsValidStatus(); err != nil {
	// 	w.SetStatusCode(http.StatusBadRequest)
	// 	w.Write([]byte(err.Error()))
	// 	return
	// }

	// task, err := h.s.UpdateTask(r.Context(), id, dto)
	// if err != nil {
	// 	writeResponseError(w, err)
	// 	return
	// }
	// resp, err := json.Marshal(task)
	// if err != nil {
	// 	w.SetStatusCode(http.StatusInternalServerError)
	// 	return
	// }

	// w.Header().Set("Content-Type", "application/json")
	// w.SetStatusCode(http.StatusOK)
	// w.Write(resp)
}
