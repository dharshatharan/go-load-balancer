package loadbalancer

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
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

func printRequestDetails(r *http.Request) {
	log.Printf("Received request form %s\n", r.RemoteAddr)
	log.Printf("%s %s %s\n", r.Method, r.URL.Path, r.Proto)
	log.Printf("Host: %s\n", r.Host)
	log.Printf("User-Agent: %s\n", r.UserAgent())
	log.Printf("Accept: %s\n", r.Header["Accept"])
}

func balance(w http.ResponseWriter, r *http.Request) {
	printRequestDetails(r)

	url, err := url.Parse("http://localhost:8081")
	if err != nil {
		log.Fatal(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.Transport = &myTransport{}
	proxy.ServeHTTP(w, r)
}
