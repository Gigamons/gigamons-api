package router

import (
	"net/http"
)

// ServerSideError Sends an serverside exception only with status code and raw plaintext.
func ServerSideError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)
	w.Write([]byte("ss_err"))
}
