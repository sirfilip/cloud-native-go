package api

import "net/http"

func Echo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-type", "text/plain")
	message := r.URL.Query().Get("message")
	w.Write([]byte(message))
}
