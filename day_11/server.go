package main

type Server struct {
	name  string
	links []*Link
	// These should probably be split into leads to out or not
	// Somehow make sure that every path is included
	paths [][]string
}

func NewServer(name string) *Server {
	return &Server{name: name}
}

func (s *Server) AddLinks(links []*Link) {
	s.links = append(s.links, links...)
}

func (s *Server) AddPath(path []string) {
	s.paths = append(s.paths, path)
}
