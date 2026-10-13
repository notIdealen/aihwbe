package transport_task

import (
	"net/http"
)

func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	// id := r.PathValue("id")

	// mrw, ok := w.(*middleware.MyResponseWriter)
	// if !ok {
	// 	panic("Bad type of response writer")
	// }

	// //проверить id на правильность написания
	// task, err := h.s.GetTask(r.Context(), id)
	// if err != nil {
	// 	respErr := GetResponseError(err)
	// 	//headers?
	// 	mrw.StatusCode = respErr.StatusCode
	// 	mrw.WriteHeader(respErr.StatusCode)
	// 	mrw.Write([]byte(respErr.Message))
	// 	return
	// }
	// resp, err := json.Marshal(task)
	// if err != nil {
	// 	mrw.SetStatusCode(http.StatusInternalServerError)
	// 	return
	// }

	// mrw.Header().Set("Content-Type", "application/json")
	// mrw.SetStatusCode(http.StatusOK)
	// // w.WriteHeader(http.StatusOK)
	// mrw.Write(resp)
}
