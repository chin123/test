package main

import "os"
import "net/http"
import "log"
import "fmt"

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello astaxie!") // send data to client side
}

func main() {
	http.HandleFunc("/", sayhelloName) // set router
	err := http.ListenAndServe(":" + os.Getenv("HTTP_PLATFORM_PORT"), nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
