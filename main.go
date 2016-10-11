package main

import (
	"fmt"
	// "net/http"
	"os"
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

	cmd := os.Args[0]

	fmt.Printf("Program Name: %s\n", cmd)

	argCount := len(os.Args[1:])

	fmt.Printf("Total Arguments (excluding program name): %d\n", argCount)

}
