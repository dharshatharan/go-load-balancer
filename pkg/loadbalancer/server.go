package loadbalancer

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
)

type server struct {
	url   *url.URL
	proxy *httputil.ReverseProxy
	alive bool
	mux   sync.RWMutex
}

func (s *server) isAlive() bool {
	s.mux.RLock()
	defer s.mux.RUnlock()
	return s.alive
}

func (s *server) setAlive(alive bool) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.alive = alive
}

func (s *server) healthCheck() {
	status := "up"

	res, err := http.Get(s.url.String() + "/health")
	if err != nil || res.StatusCode != 200 {
		s.setAlive(false)
		status = "down"
	} else {
		s.setAlive(true)
	}
	log.Printf("%s health check status: %s\n", s.url.Host, status)
}
