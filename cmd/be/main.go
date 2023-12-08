package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/dharshatharan/go-load-balancer/pkg/loadbalancer"
)

func main() {

	port := flag.Int("p", 8081, "port number")
	flag.Parse()

	log.Printf("Starting server at port %d\n", *port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		loadbalancer.PrintRequestDetails(r)

		fmt.Fprintf(w, "Hello from backend server at %s!", strconv.Itoa(*port))
		log.Printf("Replied with a hello message\n")
	})

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*port), nil))

}
