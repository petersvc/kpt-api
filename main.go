package main

import (
	"fmt"
	"kpt_api/router"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Server is getting started...")
	port := ":4000"

	log.Fatal(http.ListenAndServe(port, router.Router(port)))

}
