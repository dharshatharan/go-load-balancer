package main

import (
	"log"
	"net/http"

	"github.com/dharshatharan/go-load-balancer/pkg/loadbalancer"
)

type myTransport struct{}

// Custom Transport to print out request and response details
func (t *myTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	response, err := http.DefaultTransport.RoundTrip(request)
	if err != nil {
		return nil, err
	}

	log.Printf("Response from server: %s %s\n", response.Proto, response.Status)

	return response, nil
}

func main() {

	http.HandleFunc("/", loadbalancer.Balance)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
