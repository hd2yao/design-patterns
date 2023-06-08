package proxy

import "design-patterns/structural-patterns/proxy/servers/server"

func NewNginxServer() *nginx {
	return &nginx{
		Application:       &server.Application{},
		maxAllowedRequest: 2,
		rateLimiter:       make(map[string]int),
	}
}

type nginx struct {
	Application       *server.Application
	maxAllowedRequest int
	rateLimiter       map[string]int
}

func (n *nginx) HandlerRequest(url, method string) (int, string) {
	allowed := n.checkRateLimiting(url)
	if !allowed {
		return 403, "Not Allowed"
	}
	return n.Application.HandlerRequest(url, method)
}

func (n *nginx) checkRateLimiting(url string) bool {
	if n.rateLimiter[url] == 0 {
		n.rateLimiter[url] = 1
	}
	if n.rateLimiter[url] > n.maxAllowedRequest {
		return false
	}
	n.rateLimiter[url] = n.rateLimiter[url] + 1
	return true
}
