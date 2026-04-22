package main

import (
	"log"
	"net/http"

	"github.com/aportela/gotask/internal/router"
)

func main() {

	log.Println("starting GOTask v0.1alpha...")

	r := router.NewRouter()

	http.ListenAndServe(":3000", r)

}
