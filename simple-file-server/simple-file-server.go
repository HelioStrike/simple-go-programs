package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	port := flag.String("p", "8080", "port")
	directory := flag.String("d", ".", "directory to serve")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*directory)))

	log.Printf("Serving %s on port %s...\n", *directory, *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
