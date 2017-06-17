// Golang HTTP handler as a middleware

package main

import (
	"fmt"
	"net/http"
)

// Wrapper around http.FileServer handler
func LogHandler(h http.Handler) http.Handler {

	// HandlerFunc converts ordinary functions into wrapper
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(*r.URL)
		h.ServeHTTP(w, r)
	})
}

func main() {
	// http FileServer returns a http Handler
	fileHandler := http.FileServer(http.Dir("/tmp"))

	//
	wrappedHandler := LogHandler(fileHandler)
	http.ListenAndServe(":8080", wrappedHandler)
}
