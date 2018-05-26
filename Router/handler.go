package Router

import (
	"net/http"

	"github.com/Gigamons/gigamons_api/server"
)

func Route() {
	mux := server.GetMux()
	mux.Use(AllowEveryone)
	mux.HandleFunc("/api/user/{user}", UserRouter)
	mux.HandleFunc("/api/user", UserRouter)
}

func AllowEveryone(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}
