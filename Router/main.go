package Router

import (
	"net/http"
	"os"
)

func MainRouter(w http.ResponseWriter, r *http.Request) {
	if _, err := os.Stat("public/index.html"); os.IsExist(err) {
		http.ServeFile(w, r, "public/index.html")
	}
}
