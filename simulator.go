package CustomerTrafficSimGo

import "net/http"

type Route struct {
	method       string
	path         string
	userFunction func(*http.Request, func() *http.Response)
}

type Simulator struct {
	target string
	routes []Route
}

func NewSimulator(target string) *Simulator {
	return &Simulator{
		target: target,
		routes: []Route{},
	}
}

func (s *Simulator) AddRoute(method string, path string, function func(*http.Request, func() *http.Response)) {
	s.routes = append(s.routes, Route{
		path:         path,
		method:       method,
		userFunction: function,
	})
}

func (s *Simulator) Print() {
	println(s.target)
	for _, route := range s.routes {
		print(route.method)
		print(" - ")
		println(route.path)
	}
}
