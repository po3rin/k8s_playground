package main

import (
	"fmt"
	"log"
	"net/http"
)

func handlerRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello k8s!!!")
}

func handlerHealthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "health!!!")
}

func main() {
	http.HandleFunc("/", handlerRoot)
	http.HandleFunc("/healthz", handlerHealthz)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
