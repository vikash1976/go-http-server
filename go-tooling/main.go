// Main function of our app. The entry point to our application.
// In this function we map requerst path to its handler.
package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"strings"
)

// Main function of our app. The entry point to our application.
// In this function we map requerst path to its handler.
func main() {
	http.HandleFunc("/", myHandleFunc)
	log.Fatalln(http.ListenAndServe(":8080", nil))
	//http.ListenAndServe(":8080", nil)

}

// The handler function to root request path. In this function
// we check whether the path ends with google.com and respond accordingly.
func myHandleFunc(w http.ResponseWriter, req *http.Request) {
	//w.Header().Set("Content-Type", "text/plain")
	fmt.Printf("Request: %v\n", req)
	path := req.URL.Path[1:]

	if strings.HasSuffix(path, "google.com") {
		fmt.Fprintf(w, "Hello gopher, %s\n", strings.TrimSuffix(path, "@google.com"))
		return
	}
	fmt.Fprintf(w, "Hello dear, %s\n", path)

}
