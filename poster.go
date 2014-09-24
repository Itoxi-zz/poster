package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received")
	w.WriteHeader(200)
	w.Write([]byte("success"))
}

func main() {
	pwd, _ := os.Getwd()
	fs := http.FileServer(http.Dir(pwd))
	http.Handle("/", http.StripPrefix("/", fs))
	http.HandleFunc("/form/submitter", handler)
	http.ListenAndServe(":9090", nil)
}
