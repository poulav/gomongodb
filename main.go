package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/poulav/gomongodb/router"
)

func main() {
	fmt.Println("API using Mongo DB...")
	r := router.Router()
	fmt.Println("Server is starting...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Server has started at 4000 port...")
}
