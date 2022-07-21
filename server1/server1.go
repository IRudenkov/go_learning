package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "URLPath = %q\n", r.URL.Path)
	if err != nil {
		return
	}
}
