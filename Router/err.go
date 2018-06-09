package Router

import (
	"net/http"

	// FirstParty
	"github.com/Gigamons/gigamons_api/constants"

	// ThirdParty
	"github.com/pquerna/ffjson/ffjson"
)

func ServerSideError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)
	w.Write([]byte("Server side Error"))
}

func Exception(ex string, w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(ex))
}

func JsonException(status int, message string, w http.ResponseWriter, r *http.Request) {
	res, err := ffjson.Marshal(constants.Response{Response: status, Message: message})
	if err != nil {
		ServerSideError(w, r)
		return
	}
	w.WriteHeader(status)
	w.Write(res)
}
