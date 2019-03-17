package main

import (
	"fmt"
	"github.com/gorilla/mux"

	"net/http"
)

func main() {
	fmt.Println("Hello world")

	r := mux.NewRouter()
	http.Handle("/", r)
}
