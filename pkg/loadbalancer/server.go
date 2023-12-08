package loadbalancer

import (
	"net/http/httputil"
	"net/url"
)

type server struct {
	url   *url.URL
	proxy *httputil.ReverseProxy
}
