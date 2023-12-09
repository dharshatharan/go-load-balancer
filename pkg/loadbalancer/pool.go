package loadbalancer

import (
	"errors"
	"fmt"
	"log"
	"time"
)

type pool struct {
	servers []*server
	current int
}

var nextCounter int

func (p *pool) next() (*server, error) {
	if len(p.servers) == 0 {
		return nil, errors.New("no servers available")
	}

	fmt.Println(len(p.servers), p.current)
	next := p.servers[p.current]
	p.current = (p.current + 1) % len(p.servers)

	if next.isAlive() {
		nextCounter = 0
		return next, nil
	}
	if nextCounter >= len(p.servers) {
		nextCounter = 0
		return nil, errors.New("no alive servers")
	}
	nextCounter++
	return p.next()
}

func (p *pool) initHealthCheck() {
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			log.Println("Health check started")
			for _, s := range p.servers {
				s.healthCheck()
			}
			log.Println("Health check completed")
		}
	}
}
