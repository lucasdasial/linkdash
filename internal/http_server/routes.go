package httpserver

import (
	"fmt"
	"net/http"
)

func registerRoutes(server *http.ServeMux) {
	server.HandleFunc("GET /", home)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}
