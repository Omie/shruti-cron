package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("okay"))
}

func StartHTTPServer(host string, port string) error {
	r := mux.NewRouter()
	r.StrictSlash(true)

	r.HandleFunc("/", rootHandler).Methods("GET", "POST")

	http.Handle("/", r)

	bind := fmt.Sprintf("%s:%s", host, port)
	return http.ListenAndServe(bind, nil)
}
