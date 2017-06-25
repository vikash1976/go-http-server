package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	bugsnag "github.com/bugsnag/bugsnag-go"
	"github.com/pkg/errors"
)

var logger *log.Entry
var logFilePath = "./log/logFile_"

type TStruct struct {
	cn     string
	street string
}

type panicHandler struct {
	http.Handler
}

func (h panicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	go func() {

		//passing on to bugsnag Recover
		defer bugsnag.Recover()
	}()
	h.Handler.ServeHTTP(w, r)
}

// Main function of our app. The entry point to our application.
// In this function we map requerst path to its handler.
func main() {

	bugsnag.Configure(bugsnag.Configuration{
		APIKey:       "05bad9d39a27f13235332a6897b10748",
		ReleaseStage: "trail",
		// more configuration options
	})
	log.SetFormatter(&log.JSONFormatter{})

	t := time.Now()
	logFileSuffix := fmt.Sprintf("%d%02d%02d%02d%02d%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	logFilename := logFilePath + logFileSuffix + ".txt"
	file, err := os.OpenFile(logFilename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer file.Close()
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Info("Failed to log to file, using default stderr")
	}

	http.Handle("/panic", panicHandler{http.HandlerFunc(panicPathHandler)})
	http.HandleFunc("/", myHandleFunc)

	err = http.ListenAndServe("localhost:9090", bugsnag.Handler(nil))

	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}

}
func init() {

}

func myHandleFunc(w http.ResponseWriter, req *http.Request) {

	logger = log.WithFields(log.Fields{

		"URL":           req.URL.Path,
		"Method":        req.Method,
		"RemoteAddress": req.RemoteAddr,
		"RequestURI":    req.RequestURI,
	})
	logger.Info("Processing Request")
	w.Header().Set("Content-Type", "text/plain")
	path := req.URL.Path[1:]

	_, err := isAuthenticated(req)
	if err != nil {
		err = errors.Wrap(err, "Authentication Failed")
		logger.Errorf("%v", err)
		logger.Errorf("%+v", err)
		bugsnag.Notify(err,
			bugsnag.SeverityError)

		http.Error(w, errors.Cause(err).Error(), http.StatusUnauthorized)
		return
	}
	auth := strings.SplitN(req.Header["Authorization"][0], " ", 2)
	payload, _ := base64.StdEncoding.DecodeString(auth[1])
	pair := strings.SplitN(string(payload), ":", 2)
	if strings.HasSuffix(path, "google.com") {
		fmt.Fprintf(w, "Hello Gopher, %s. Here is what you sent: %s %s\n", strings.TrimSuffix(path, "@google.com"), pair[0], pair[1])
		return
	}
	fmt.Fprintf(w, "Hello dear, %s.  Here is what you sent: %s %s\n", path, pair[0], pair[1])

}
func panicPathHandler(w http.ResponseWriter, req *http.Request) {

	//create a panic situation
	names := []string{"aname", "bname", "cname"}

	m := make(map[string]map[string]TStruct, len(names))
	for _, name := range names {
		m["uid"][name] = TStruct{cn: "Chaithra", street: "dp road"}
	}
}

func Validate(username, password string, err error) (bool, error) {
	var errL error
	if password == username+"!!" {
		fmt.Println("U and P matched")
		return true, nil
	}
	fmt.Println("U and P mismatched")
	// adding additional field to capture in log statement.
	// userId as identifier to the captured error
	logger.Data["userId"] = username
	if err == nil {
		errL = errors.New("Invalid Credentials")
	} else {
		errL = errors.Wrap(err, "Invalid Credentials")
	}
	return false, errL

}

func isAuthenticated(req *http.Request) (bool, error) {
	fmt.Println("isAuthenticated")
	var err error
	if req.Header["Authorization"] == nil {
		err = errors.New("Authorization missing")
		return false, err
	}
	auth := strings.SplitN(req.Header["Authorization"][0], " ", 2)

	if len(auth) != 2 || auth[0] != "Basic" {
		err = errors.Wrap(err, "Bad Authorization Syntax")
		return false, err
	}

	payload, _ := base64.StdEncoding.DecodeString(auth[1])
	pair := strings.SplitN(string(payload), ":", 2)

	if len(pair) != 2 {
		err = errors.Wrap(err, "Username or Password is missing")
		return false, err
	}

	if _, err1 := Validate(pair[0], pair[1], err); err1 != nil {
		fmt.Println("U and P didn't match")
		return false, err1
	}
	return true, nil
}
