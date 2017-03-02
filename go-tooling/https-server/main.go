package main

import (
	"encoding/base64"
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
	log.Fatalln(http.ListenAndServe("ap-pun-lp1408.internal.sungard.corp:8080", nil))
	//http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)

}

//var re = regexp.MustCompile("^(.+)@google.com$")
// The handler function to root request path. In this function
// we check whether the path ends with google.com and respond accordingly.
func myHandleFunc(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	path := req.URL.Path[1:]
	fmt.Printf("Request: %v\n", req)

	if req.Header["Authorization"] == nil {
		http.Error(w, "Authorization missing", http.StatusBadRequest)
		return
	}
	auth := strings.SplitN(req.Header["Authorization"][0], " ", 2)

	if len(auth) != 2 || auth[0] != "Basic" {
		http.Error(w, "bad syntax", http.StatusBadRequest)
		return
	}

	payload, _ := base64.StdEncoding.DecodeString(auth[1])
	pair := strings.SplitN(string(payload), ":", 2)

	if len(pair) != 2 || !Validate(pair[0], pair[1]) {
		http.Error(w, "authorization failed", http.StatusUnauthorized)
		return
	}
	//match := re.FindAllStringSubmatch(path, -1)
	if strings.HasSuffix(path, "google.com") {
		fmt.Fprintf(w, "Hello Gopher, %s. Here is what you sent: %s %s\n", strings.TrimSuffix(path, "@google.com"), pair[0], pair[1])
		return
	}
	fmt.Fprintf(w, "Hello dear, %s.  Here is what you sent: %s %s\n", path, pair[0], pair[1])

}

func Validate(username, password string) bool {

	fmt.Println(username)
	fmt.Println(password)
	/*if username == "auser" && password == "apass" {
		return true
	}
	return false*/
	return true
}
