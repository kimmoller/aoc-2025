package main

import (
	"fmt"
	"strings"
)

type Center struct {
	servers []Server
}

func NewCenter() *Center {
	return &Center{}
}

func (c *Center) AddServer(server Server) {
	c.servers = append(c.servers, server)
}

func (c *Center) Servers() []Server {
	return c.servers
}

func (c *Center) Server(name string) (*Server, error) {
	for _, server := range c.servers {
		if server.name == name {
			return &server, nil
		}
	}
	return nil, fmt.Errorf("server %s not found", name)
}

func (c *Center) Paths(start, end string) (int, error) {
	return c.getPathsUsingDFS(start, make(map[string]int), end)
}

func (c *Center) PathsWithMiddleSteps(start, end string) (int, error) {
	svrToFft, err := c.getPathsUsingDFS(start, make(map[string]int), "fft")
	if err != nil {
		return 0, err
	}
	fftToDac, err := c.getPathsUsingDFS("fft", make(map[string]int), "dac")
	if err != nil {
		return 0, err
	}
	dacToOut, err := c.getPathsUsingDFS("dac", make(map[string]int), end)
	if err != nil {
		return 0, err
	}
	svrToDac, err := c.getPathsUsingDFS(start, make(map[string]int), "dac")
	if err != nil {
		return 0, err
	}
	dacToFft, err := c.getPathsUsingDFS("dac", make(map[string]int), "fft")
	if err != nil {
		return 0, err
	}
	fftToOut, err := c.getPathsUsingDFS("fft", make(map[string]int), end)
	if err != nil {
		return 0, err
	}
	return (svrToFft * fftToDac * dacToOut) + (svrToDac * dacToFft * fftToOut), nil
}

func (c *Center) getPathsUsingDFS(currentServer string, visited map[string]int, end string) (int, error) {
	if currentServer == end {
		return 1, nil
	}

	if count, ok := visited[currentServer]; ok {
		return count, nil
	}

	server, err := c.Server(currentServer)
	if err != nil {
		return 0, err
	}

	count := 0
	for _, link := range server.links {
		value, err := c.getPathsUsingDFS(link.output, visited, end)
		if err != nil {
			return 0, err
		}
		count += value
	}

	visited[currentServer] = count
	return count, nil
}

func (c *Center) PopulateCenter(data []string) error {
	servers := []Server{}
	serverLinks := map[string][]string{}
	for _, input := range data {
		parts := strings.Split(input, ":")

		serverName := parts[0]
		server := NewServer(serverName)
		servers = append(servers, server)

		output := parts[1]
		outputs := strings.Split(output, " ")

		filteredOutputs := []string{}
		for _, output := range outputs {
			if output != "" {
				filteredOutputs = append(filteredOutputs, output)
			}
		}

		serverLinks[serverName] = filteredOutputs
	}

	serversWithLinks := []Server{}
	for _, server := range servers {
		links := []Link{}
		if outputs, ok := serverLinks[server.name]; ok {
			for _, output := range outputs {
				link := NewLink(output)
				links = append(links, link)
			}
			server.AddLinks(links)
		}
		serversWithLinks = append(serversWithLinks, server)
	}

	// Create out server
	outServer := NewServer("out")
	serversWithLinks = append(serversWithLinks, outServer)

	c.servers = serversWithLinks

	return nil
}
