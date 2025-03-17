package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)


func main() {
	config.ToLoad()



	fmt.Println(config.StringConnectionDataBase)

	fmt.Println("Running API!")

	r := router.ToGenerate()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))

}