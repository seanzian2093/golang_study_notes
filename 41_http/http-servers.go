package main

import (
    "fmt"
    "net/http"
)

// Functions serving as handlers take a http.ResponseWriter and a http.Request as arguments.
// The response writer is used to fill in the HTTP response. 
// Here our simple response is just “hello\n”
func hello(w http.ResponseWriter, req *http.Request) {

    fmt.Fprintf(w, "hello\n")
}

// a more complicated handler
func headers(w http.ResponseWriter, req *http.Request) {

    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func main() {

	// register our handlers on server routes using the http.HandleFunc convenience function. 
	// It sets up the default router in the net/http package and takes a function as an argument.
    http.HandleFunc("/hello", hello)
    http.HandleFunc("/headers", headers)

	// call the ListenAndServe with the port and a handler. nil tells it to use the default router we’ve just set up.
    http.ListenAndServe(":8090", nil)
}