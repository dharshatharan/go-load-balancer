package loadbalancer

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var serverPool pool

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

func PrintRequestDetails(r *http.Request) {
	log.Printf("Received request form %s\n", r.RemoteAddr)
	log.Printf("%s %s %s\n", r.Method, r.URL.Path, r.Proto)
	log.Printf("Host: %s\n", r.Host)
	log.Printf("User-Agent: %s\n", r.UserAgent())
	log.Printf("Accept: %s\n", r.Header["Accept"])
}

func createServers() {
	ports := []string{"8081", "8082"}
	for _, port := range ports {
		url, err := url.Parse("http://localhost:" + port)
		if err != nil {
			log.Fatal(err)
		}
		proxy := httputil.NewSingleHostReverseProxy(url)
		proxy.Transport = &myTransport{}
		serverPool.servers = append(serverPool.servers, &server{url: url, proxy: proxy})
	}
}

func Init() {
	createServers()
}

func Balance(w http.ResponseWriter, r *http.Request) {
	PrintRequestDetails(r)

	server := serverPool.next()
	if server == nil {
		http.Error(w, "Service not available", http.StatusServiceUnavailable)
		return
	}
	server.proxy.ServeHTTP(w, r)
}
