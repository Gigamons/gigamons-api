package router

import (
	"net/http"

	"github.com/Gigamons/gigamons_api/constants"
	"github.com/pquerna/ffjson/ffjson"
)

// JSONAnswer is writting a Json Answer. {Response: 200, Message: "Hello World"} as example.
func JSONAnswer(status int, message interface{}, w http.ResponseWriter, r *http.Request) {
	res, err := ffjson.Marshal(constants.Response{Response: status, Message: message})
	if err != nil {
		ServerSideError(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(res)
}
