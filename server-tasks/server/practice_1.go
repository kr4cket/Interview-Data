package main

import (
	"fmt"
	"net/http"
	"time"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!\nRequest method:%s", r.Method)
}

func main() {

	handler := http.NewServeMux()
	handler.HandleFunc("/", homeHandler)

	serv := &http.Server{
		Addr:           ":80",
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   30 * time.Second,
	}

	err := serv.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
}
