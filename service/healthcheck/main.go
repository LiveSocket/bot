package healthcheck

import (
	"fmt"
	"net/http"
)

// HTTPHandler A handler function for HTTP requests
type handler = func(w http.ResponseWriter, r *http.Request)

// Init Starts a standard healthcheck listener on port 80
func Init() {
	go start("80")
}

// start Starts a healthcheck on the provided port
func start(port string) {
	handlers := map[string]handler{
		"/": func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "healthy")
		},
	}
	err := open(port, handlers)
	panic(err)
}

// open Opens and starts a new HTTP server
func open(port string, handlers map[string]handler) error {
	for key, value := range handlers {
		http.HandleFunc(key, value)
	}

	return http.ListenAndServe(":"+port, nil)
}
