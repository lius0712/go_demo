package main

import (
	"fmt"
	"log"
	"net/http"
)

//curl http://localhost:9999/
//URL.path="/"
//curl http://localhost:9999/hello
//Header["User-Agent"]=["curl/7.85.0"]
//Header["Accept"]=["*/*"]

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":9999", nil))
}
func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.path=%q\n", req.URL.Path)
}
func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q]=%q\n", k, v)
	}

}
