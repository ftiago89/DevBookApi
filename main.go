package main

import (
	"devbookapi/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Print("Executando API!")
	r := router.GenerateRouter()

	log.Fatal(http.ListenAndServe(":5000", r))
}
