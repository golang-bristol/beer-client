package main

import (
	"fmt"
	"net/http"
)

/*
Supported commands:
GET /beers - return list of beers
POST /beers - create a new beer
GET /beer/[:id] - return a single beer
GET /beer/[:id]/reviews - return a list of reviews for beer
POST /beer/[:id]/review - create a new review for beer
*/

func main() {
	fmt.Println("Hello world!")
}
