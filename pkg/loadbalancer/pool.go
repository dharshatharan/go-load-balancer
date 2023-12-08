package loadbalancer

import "fmt"

type pool struct {
	servers []*server
	current int
}

func (p *pool) next() *server {
	if len(p.servers) == 0 {
		return nil
	}

	fmt.Println(len(p.servers), p.current)
	next := p.servers[p.current]
	p.current = (p.current + 1) % len(p.servers)
	return next
}
