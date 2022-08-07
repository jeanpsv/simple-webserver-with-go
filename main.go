package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	var fileServer http.Handler = http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	fmt.Printf("Starting server at port 8080\n")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
