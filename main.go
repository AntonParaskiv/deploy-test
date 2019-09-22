package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", baseEndpoint)
	fmt.Println(http.ListenAndServe(":8081", nil))

}

func baseEndpoint(w http.ResponseWriter, r *http.Request) {
	response := []byte(`Alright`)
	w.Write(response)
	return
}