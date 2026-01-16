package main

type Server struct {
	name  string
	links []Link
}

func NewServer(name string) Server {
	return Server{name: name}
}

func (s *Server) AddLinks(links []Link) {
	s.links = append(s.links, links...)
}
