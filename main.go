package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/list/vcn", listVCN)
	mux.HandleFunc("/list/compute", listCompute)
	fmt.Println("Server started on port 8080.")
	http.ListenAndServe(":8080", mux)
}
