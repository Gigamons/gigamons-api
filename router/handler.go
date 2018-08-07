package router

import (
	"net/http"

	"github.com/Gigamons/gigamons_api/server"
)

// Route is setting the Routes.
func Route() {
	mux := server.GetMux()
	mux.Use(AllowEveryone)
	mux.HandleFunc("/api/v1/user/{user}", UserRouter)
	mux.HandleFunc("/api/v1/user", UserRouter)
	mux.HandleFunc("/api/v1/news/{page}", News)
	mux.HandleFunc("/api/v1/news", News)
}

// AllowEveryone is allowing everyone to use the api. (that other people can use this api)
func AllowEveryone(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		next.ServeHTTP(w, r)
	})
}
