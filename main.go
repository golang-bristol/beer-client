package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

/*
Supported commands:
GET /beers - return list of beers
POST /beers - create a new beer
GET /beer/[:id] - return a single beer
GET /beer/[:id]/reviews - return a list of reviews for beer
POST /beer/[:id]/review - create a new review for beer
*/

var requestURL = flag.String("url", "http://localhost:8080/beers", " webserver ")
var requestMethod = flag.String("method", "GET", "GET/PUT")
var id = flag.String("id", "", "id")
var requestData = flag.String("data", "", "post data")

func main() {
	flag.Parse()
	switch strings.ToUpper(*requestMethod) {
	case "GET":
		res, err := http.Get(*requestURL)
		if err != nil {
			// Request failed.
			fmt.Println("Error executing HTTP request :", err)
			return
		}

		//defer res.Body.Close()

		responseData, err := ioutil.ReadAll(res.Body)
		if err != nil {
			// Could not read data.
			fmt.Println("Error reading response data :", err)
			return
		}
		fmt.Println(string(responseData))

	}
	// Print received data.

}
