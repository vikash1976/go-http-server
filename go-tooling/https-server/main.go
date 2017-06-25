package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strings"
)

var logFilePath = "./log/regLogFile.txt"

type panicHandler struct {
	http.Handler
}

type TStruct struct {
	cn     string
	street string
}

func (h panicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	/*go func() {
		//passing on to bugsnag Recover
		defer bugsnag.Recover()
	}()*/
	h.Handler.ServeHTTP(w, r)
}

// Main function of our app. The entry point to our application.
// In this function we map requerst path to its handler.
func main() {
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Println("Failed to log to file, using default stderr")
	}
	http.Handle("/panic", panicHandler{http.HandlerFunc(panicPathHandler)})
	http.HandleFunc("/", myHandleFunc)
	log.Fatalln(http.ListenAndServe("10.253.98.20:8080", nil))
	//http.ListenAndServeTLS("ap-pun-lp1408.internal.sungard.corp:10443", "cert.pem", "key.pem", nil)

}

func myHandleFunc(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	path := req.URL.Path[1:]
	log.Printf("Request: %v\n", req)

	if _, err := isAuthenticated(w, req); err != nil {
		log.Printf("Req processing status: %v\n", err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	log.Println("Processing Request")
	auth := strings.SplitN(req.Header["Authorization"][0], " ", 2)

	payload, _ := base64.StdEncoding.DecodeString(auth[1])
	pair := strings.SplitN(string(payload), ":", 2)
	if strings.HasSuffix(path, "google.com") {
		fmt.Fprintf(w, "Hello Gopher, %s. Here is what you sent: %s %s\n", strings.TrimSuffix(path, "@google.com"), pair[0], pair[1])
		return
	}
	fmt.Fprintf(w, "Hello dear, %s.  Here is what you sent: %s %s\n", path, pair[0], pair[1])
	return

}
func panicPathHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("Request: %v\n", req)
	//create a panic situation
	names := []string{"aname", "bname", "cname"}

	m := make(map[string]map[string]TStruct, len(names))
	for _, name := range names {
		m["uid"][name] = TStruct{cn: "Chaithra", street: "dp road"}
	}
}

func Validate(username, password string) (bool, error) {

	fmt.Println(username)
	fmt.Println(password)
	if password == username+"!!" {
		return true, nil
	}
	return false, errors.New("Invalid Credentials")
}

func isAuthenticated(w http.ResponseWriter, req *http.Request) (bool, error) {
	if req.Header["Authorization"] == nil {
		err := errors.New("Authorization missing")
		return false, err
	}
	auth := strings.SplitN(req.Header["Authorization"][0], " ", 2)

	if len(auth) != 2 || auth[0] != "Basic" {
		err := errors.New("Bad Authorization Syntax")
		return false, err
	}

	payload, _ := base64.StdEncoding.DecodeString(auth[1])
	pair := strings.SplitN(string(payload), ":", 2)

	if len(pair) != 2 {
		err := errors.New("Username or Password is missing")
		return false, err
	}

	if _, err := Validate(pair[0], pair[1]); err != nil {
		return false, errors.New("Credentials validation failed")
	}
	return true, nil

}
