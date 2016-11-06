package main

import (
	"os"
	"net/http"
	"log"
	"fmt"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!") // send data to client side
}

func main() {
	http.HandleFunc("/", sayhelloName) // set router
	err := http.ListenAndServe(":" + os.Getenv("HTTP_PLATFORM_PORT"), nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
