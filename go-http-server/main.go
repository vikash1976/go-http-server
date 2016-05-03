package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"log"
	"net/http"
	"path"
	"encoding/json"
)

type Price struct {
	Tick string `json:"tick"`
	Price float64 `json:"price"`
}

type Prices struct {
	SymbolInfo []Price `json:"symbolInfo"`
}


func main() {

	//fs := http.FileServer(http.Dir("static"))
	//http.Handle("/static/", http.StripPrefix("/static/", fs))
	r := httprouter.New()
	r.GET("/", HomeHandler)

	r.GET("/static/*html", StaticFilesHandler)

	r.GET("/dynamic/*html", DynamicFilesHandler)
	// Posts collection
	r.GET("/prices", PricesIndexHandler)
	r.POST("/posts", PostsCreateHandler)

	// Posts singular
	r.GET("/prices/:id", PricesShowHandler)
	r.PUT("/posts/:id", PostUpdateHandler)
	r.GET("/posts/:id/edit", PostEditHandler)

	fmt.Println("Starting server on :9000")
	http.ListenAndServe(":9000", r)
}

func DynamicFilesHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	myPrice := Price{"VRL", 10.23}
	_, file := path.Split(r.URL.Path)
	lp := path.Join("dynamic", "templates", "layout.html")
	fp := path.Join("dynamic", "templates", file)
	fmt.Println("LP: ", lp)

	fmt.Println("FP: ", file)
	//path.Split(fp)

	tmpl, _ := template.ParseFiles(lp, fp)

	if err := tmpl.ExecuteTemplate(rw, "layout", myPrice); err != nil {
		log.Println(err.Error())
		http.Error(rw, http.StatusText(500), 500)
	}
}
func StaticFilesHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println("StaticFilesHandler", r.URL)
	http.FileServer(http.Dir(".")).ServeHTTP(rw, r)
	
}
func HomeHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "Home")
}

func PricesIndexHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	prices := make([]Price, 2)
	
	prices = append(prices, Price{"FIS", 56.87})
	prices = append(prices, Price{"TCS", 46.83})
	
	//priceDetails := Prices {SymbolInfo: []Price{{"FIS", 56.87}, {"TCS", 46.83}}}
	var priceDetails Prices
	
	priceDetails.SymbolInfo = make([]Price, 2)
	priceDetails.SymbolInfo[0] = Price{"INFY", 26.87}
	priceDetails.SymbolInfo = append(priceDetails.SymbolInfo, Price{"FIS", 56.87})
	
	priceDetails.SymbolInfo = append(priceDetails.SymbolInfo, Price{"TCS", 46.83})
	jsonPrices, _ := json.Marshal(priceDetails)
	
	//prices := `[{"tick": "FIS", "price": $68.43},{"tick": "INFY", "price": $38.43}]`
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	rw.Header().Set("Access-Control-Allow-Origin", "*")

	io.WriteString(rw, string(jsonPrices))
}

func PostsCreateHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "posts create")
}

func PricesShowHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	prices := `{"tick": "` + id + `", "price": $68.43}`
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	rw.Header().Set("Access-Control-Allow-Origin", "*")

	io.WriteString(rw, prices)
}

func PostUpdateHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "post update")
}

func PostDeleteHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "post delete")
}

func PostEditHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "post edit")
}

/*package main

import (
	"io"
	"net/http"
	"fmt"
	"log"
	"html/template"
	"path"
	"os"
)

type myHandler int

func (h myHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Println("In API:", req.URL.Path)

	fmt.Println(req.RequestURI)
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	switch req.URL.Path {
	case "/api/cat":
		io.WriteString(res, `<p><strong>kitty kitty kitty<strong></p><img src="https://upload.wikimedia.org/wikipedia/commons/0/06/Kitten_in_Rizal_Park%2C_Manila.jpg">`)
	case "/api/dog":
		io.WriteString(res, `<p><strong>doggy doggy doggy<strong></p><img src="https://upload.wikimedia.org/wikipedia/commons/6/6e/Golde33443.jpg">`)
		case "/api/prices":
		prices := `{"tick": "FIS", "price": $68.43}`
		res.Header().Set("Content-Type", "application/json; charset=utf-8")
		res.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		res.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		res.Header().Set("Access-Control-Allow-Origin", "*")

		io.WriteString(res, prices)
		default:
		fmt.Println("In Default")
		//res.WriteHeader(http.StatusFound)
		http.Redirect(res, req, "http://www.google.co.in", 302)

	}
}

func main() {
	var h myHandler

  fs := http.FileServer(http.Dir("static"))
  http.Handle("/static/", http.StripPrefix("/static/", fs))

  log.Println("Listening...")
	 http.Handle("/api/", h)

	http.HandleFunc("/dynamic/", serveTemplate)

	http.ListenAndServe(":9000", nil)
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
  fmt.Println("In ServeTemplate: ", r.URL)
   _, file := path.Split(r.URL.Path)
  lp := path.Join("dynamic", "templates", "layout.html")
  fp := path.Join("dynamic", "templates", file)
  fmt.Println("LP: ", lp)

  fmt.Println("FP: ", file)
  //path.Split(fp)

  // Return a 404 if the template doesn't exist
  info, err := os.Stat(fp)
  if err != nil {
    if os.IsNotExist(err) {
      http.NotFound(w, r)
      return
    }
  }

  // Return a 404 if the request is for a directory
  if info.IsDir() {
    http.NotFound(w, r)
    return
  }

  tmpl, err := template.ParseFiles(lp, fp)
  if err != nil {
    // Log the detailed error
    log.Println(err.Error())
    // Return a generic "Internal Server Error" message
    http.Error(w, http.StatusText(500), 500)
    return
  }

  if err := tmpl.ExecuteTemplate(w, "layout", nil); err != nil {
    log.Println(err.Error())
    http.Error(w, http.StatusText(500), 500)
  }
}*/
