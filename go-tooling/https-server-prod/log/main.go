/*package main

import (
	"fmt"
	"log"
	"os"
	"time"





	"github.com/mitchellh/panicwrap"
)

type T struct {
	cn     string
	street string
}

func main() {
	exitStatus, err := panicwrap.BasicWrap(reportPanics)
	if err != nil {
		// Something went wrong setting up the panic wrapper. Unlikely, but possible.
		panic(err)
	}

	// If exitStatus >= 0, then we're the parent process and the panicwrap
	// re-executed ourselves and completed. Just exit with the proper status.
	if exitStatus >= 0 {
		os.Exit(exitStatus)
	}

	//runWithoutPanicReporting()
	//panic("Main Deliberate!!!")
	// Let's say we panic
	//panic("oh shucks")
	go myFunc()

	time.Sleep(10 * time.Second)

}

var panicFilePath = "./panic/"

func reportPanics(msg string) {
	t := time.Now()
	panicFileSuffix := fmt.Sprintf("%d%02d%02d%02d%02d%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	panicFilename := panicFilePath + "abc_" + panicFileSuffix + ".txt"
	//panicFileSuffix := now.Year() + now.Month() + Now.Day() + now.Hour() + now.Minute() + now.Second()
	file, err := os.OpenFile(panicFilename, os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	if err == nil {
		//log.SetOutput(file)
		file.Write([]byte(msg))
	} else {
		log.Println("Failed to log panic to file, using default stderr", panicFileSuffix)
	}
	fmt.Printf("Panic Details: %v\n", msg)
	os.Exit(1)

}

func runWithoutPanicReporting() {
	go myFunc()
}

func myFunc() {
	fmt.Println("Called On")
	//create a panic situation
	names := []string{"kasi", "remya", "nandan"}

	m := make(map[string]map[string]T, len(names))
	for _, name := range names {
		m["uid"][name] = T{cn: "Chaithra", street: "fkmp"}
	}
	fmt.Println(m)
}

func panicHandler(output string) {
	// output contains the full output (including stack traces) of the
	// panic. Put it in a file or something.
	fmt.Printf("The child panicked:\n\n%s\n", output)
	os.Exit(1)
}*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"
)

type panicHandler struct {
	http.Handler
}

var panicFilePath = "./panic/"

func reportPanics(msg string) {
	t := time.Now()
	panicFileSuffix := fmt.Sprintf("%d%02d%02d%02d%02d%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	panicFilename := panicFilePath + "abc_" + panicFileSuffix + ".txt"
	//panicFileSuffix := now.Year() + now.Month() + Now.Day() + now.Hour() + now.Minute() + now.Second()
	file, err := os.OpenFile(panicFilename, os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	if err == nil {
		//log.SetOutput(file)
		file.Write([]byte(msg))
	} else {
		log.Println("Failed to log panic to file, using default stderr", panicFileSuffix)
	}
	//fmt.Printf("Panic Details: %v\n", msg)
	//os.Exit(1)

}
func (h panicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			buf := make([]byte, 1<<20)
			n := runtime.Stack(buf, true)
			//fmt.Fprintf(os.Stderr, "panic: %v\n\n%s", err, buf[:n])
			reportPanics(string(buf[:n]))
			//os.Exit(100)

		}
	}()
	h.Handler.ServeHTTP(w, r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	panic("oh no!")
}

func main() {
	http.Handle("/", panicHandler{http.HandlerFunc(handler)})
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
