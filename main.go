package main

import (
	"fmt"
	"net/http"
)

const Message = "This is the new another yet test"

func main() {

	http.HandleFunc("/", baseEndpoint)
	fmt.Println(http.ListenAndServe(":8081", nil))

}

func baseEndpoint(w http.ResponseWriter, r *http.Request) {
	response := []byte(printMessage())
	w.Write(response)
	return
}

func printMessage() (message string) {
	message = Message
	return
}
