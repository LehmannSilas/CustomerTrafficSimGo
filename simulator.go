package CustomerTrafficSimGo

type Route struct {
	method string
	path   string
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

func (s *Simulator) AddRoute(method string, path string) {
	s.routes = append(s.routes, Route{path: path, method: method})
}

func (s *Simulator) Print() {
	println(s.target)
	for _, route := range s.routes {
		print(route.method)
		print(" - ")
		println(route.path)
	}
}
