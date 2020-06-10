package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	x, err := httputil.DumpRequest(r, true)
	log.Println(fmt.Sprintf("%q", x))
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	switch r.Method {
	case "GET":
		w.Write([]byte("Hello there!"))
	case "POST":
		name := r.FormValue("name")
		w.Write([]byte(fmt.Sprintf("General, %s!", name)))
	default:
		w.Write([]byte("Unknown Request."))
	}
}

func main() {
	var port = 8080

	http.HandleFunc("/", handler)

	log.Printf("Starting server on port %d...\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal(err)
	}
}
