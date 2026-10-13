package transport_task

import "net/http"

func (h *TaskHandler) GetList(rw http.ResponseWriter, r *http.Request) {

	// w, ok := rw.(*middleware.MyResponseWriter)
	// if !ok {
	// 	panic("Bad type of response writer")
	// }

	// tasks, err := h.s.GetList(r.Context())
	// if err != nil {
	// 	writeResponseError(w, err)
	// 	return
	// }

	// var resp bytes.Buffer

	// enc := json.NewEncoder(&resp)
	// for _, task := range tasks {
	// 	err = enc.Encode(task)
	// 	if err != nil {
	// 		writeResponseError(w, err)
	// 		return
	// 	}
	// }

	// w.Header().Set("Content-Type", "application/json")
	// w.SetStatusCode(http.StatusOK)
	// w.Write(resp.Bytes())
}
