package server

import (
	"fmt"
	"net/http"
)


func RunServer(port int) {
    fs := http.FileServer(http.Dir("assets"))

    mux := http.NewServeMux()

    mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
    mux.HandleFunc("/", indexHandler)

    http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
}
