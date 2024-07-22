package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/wadearnold/recurring"
)

func main() {
	// Setup routes
	http.HandleFunc("/recurrings", recurring.RecurringJSON)

	// Start the server
	fmt.Println("Recurring Server is listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
