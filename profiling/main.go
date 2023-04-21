package main

import (
	"fmt"
	"net/http"

	_ "net/http/pprof" // the _ means we are not using the package explicitly
)

// write hello world Go http server
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
	http.ListenAndServe(":8080", nil)
}
