package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {

	port := flag.Int("p", 8081, "port number")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request form %s\n", r.RemoteAddr)
		log.Printf("%s %s %s\n", r.Method, r.URL.Path, r.Proto)
		log.Printf("Host: %s\n", r.Host)
		log.Printf("User-Agent: %s\n", r.UserAgent())
		log.Printf("Accept: %s\n", r.Header["Accept"])

		fmt.Fprintf(w, "Hello from backend server at %s!", strconv.Itoa(*port))
		log.Printf("Replied with a hello message\n")
	})

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*port), nil))

}
