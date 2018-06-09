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
	mux.PathPrefix("/static/").Handler(http.FileServer(http.Dir("./public/")))
	mux.Path("/").Handler(http.FileServer(http.Dir("./public/")))
	mux.NotFoundHandler = Rewrite()
}

func Rewrite() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/#"+r.URL.Path, http.StatusTemporaryRedirect)
	})
}

func AllowEveryone(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		next.ServeHTTP(w, r)
	})
}
