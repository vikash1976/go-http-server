package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"encoding/json"
)
//Price structure defined
type Price struct {
	Tick string `json:"tick"`
	Price float64 `json:"price"`
}

//Prices Structure
type Prices struct {
	SymbolInfo []Price `json:"price"`
}


func main() {

	
	r := httprouter.New()
	r.GET("/", HomeHandler)

	r.GET("/static/*html", StaticFilesHandler)

	r.GET("/dynamic/*html", DynamicFilesHandler)
	// Prices - Get Prices
	r.GET("/prices", PricesIndexHandler)
	

	// Get a single Price
	r.GET("/prices/:id", PricesShowHandler)
	// Post a single Price
	r.PUT("/prices/:id", PostUpdateHandler)
	

	log.Println("Starting server on :9000")
	http.ListenAndServe(":9000", r)
}
// Handles templated content
func DynamicFilesHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	myPrice := Price{"VRL", 10.23}
	_, file := path.Split(r.URL.Path)
	lp := path.Join("dynamic", "templates", "layout.html")
	fp := path.Join("dynamic", "templates", file)
	log.Println("LP: ", lp)

	log.Println("FP: ", file)
	
	tmpl, _ := template.ParseFiles(lp, fp)

	if err := tmpl.ExecuteTemplate(rw, "layout", myPrice); err != nil {
		log.Println(err.Error())
		http.Error(rw, http.StatusText(500), 500)
	}
}
// Handles static contents
func StaticFilesHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Println("StaticFilesHandler", r.URL)
	http.FileServer(http.Dir(".")).ServeHTTP(rw, r)
	
}
// Handler function for / route
func HomeHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "Home")
}

// Handler function for /prices route
func PricesIndexHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Println("In Prices", r.URL)
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
	log.Println("Done Prices")
}

// Handler function for /prices:id route
func PricesShowHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	prices := `{"tick": "` + id + `", "price": $68.43}`
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	rw.Header().Set("Access-Control-Allow-Origin", "*")

	io.WriteString(rw, prices)
}
// Handler function for /prices/:id PUT route
func PostUpdateHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "post update")
	log.Println("Param is: ", p)
	body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        log.Fatalln("some error")
    }
    log.Println(string(body))
	 //var t map[string]interface{}
	 
	 var t Price
    err = json.Unmarshal(body, &t)
    if err != nil {
        log.Fatalln("some other error: ", err)
    }
    log.Println(t.Price)
	
}


