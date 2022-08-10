package main

import (
	"devbookapi/src/config"
	"devbookapi/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadEnv()

	fmt.Printf("Executando API na porta: %d!", config.SERVER_PORT)
	r := router.GenerateRouter()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.SERVER_PORT), r))
}
