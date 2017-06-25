package main

import (
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http/pprof"
)

var done chan bool

func main() {

	http.HandleFunc("/", handleRequest)

	log.Fatalln(http.ListenAndServe(":8080", http.DefaultServeMux))
	
}
func handleRequest(w http.ResponseWriter, req *http.Request) {
	// make a new channel
	done = make(chan bool)
	for i := 0; i < 3; i++ {
		go doSomething()
	}
	// wait for all the goroutines to finish, and return
	for i := 0; i < 3; i++ {
		<-done
	}
	w.Write([]byte("Done processing"))
}

func doSomething() {
	// signal we are done doing something
	defer func() { done <- true }()
	// perform a web request
	resp, err := http.Get("http://localhost:50055/debug/requests")
	if err != nil {
		log.Fatal(err)
	}
	//defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		_, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
	}
}
