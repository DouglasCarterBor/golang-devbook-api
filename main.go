package main

import (
    "api/src/router"
    "fmt"
	"log"
	"net/http"
)


func main() {
	fmt.Println("Running API!")

	r := router.ToGenerate()

	log.Fatal(http.ListenAndServe("0.0.0.0:5000", r))

}