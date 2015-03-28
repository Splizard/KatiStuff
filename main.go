package main

import "net/http"
import "log"

func main() {
	log.Fatal(http.ListenAndServe(":80", http.FileServer(http.Dir("."))))
}
