package main

import (
	"fmt"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

type Center struct {
	servers []*Server
}

func NewCenter() *Center {
	return &Center{}
}

func (c *Center) AddServer(server *Server) {
	c.servers = append(c.servers, server)
}

func (c *Center) Servers() []*Server {
	return c.servers
}

func (c *Center) Server(name string) (*Server, error) {
	for _, server := range c.servers {
		if server.name == name {
			return server, nil
		}
	}
	return nil, fmt.Errorf("server %s not found", name)
}

func (c *Center) FindAllPaths(start, end string, strict bool) ([][]string, error) {
	allPaths := [][]string{}
	startingServer, err := c.Server(start)
	if err != nil {
		return nil, err
	}
	err = c.FindPath(&allPaths, []string{}, startingServer.name, end)
	if err != nil {
		return nil, err
	}

	finalPaths := [][]string{}
	if strict {
		for _, path := range allPaths {
			if meetsRequirements(path) {
				finalPaths = append(finalPaths, path)
			}
		}
	} else {
		finalPaths = allPaths
	}

	return finalPaths, nil
}

func (c *Center) FindPath(finalPaths *[][]string, currentPath []string, currentServerName string, finalServerName string) error {
	if currentServerName == finalServerName {
		*finalPaths = append(*finalPaths, currentPath)
		return nil
	}

	server, err := c.Server(currentServerName)
	if err != nil {
		return err
	}

	for _, link := range server.links {
		output := link.output
		newPath := append(currentPath, output)

		err := c.FindPath(finalPaths, newPath, output, finalServerName)
		if err != nil {
			return err
		}
	}

	return nil
}

func markKnownPaths(center *Center, path []string) error {
	if len(path) == 0 {
		return nil
	}

	for i, serverName := range path {
		server, err := center.Server(serverName)
		if err != nil {
			return err
		}
		knownPath := path[i+1:]
		server.AddPath(knownPath)
	}

	return nil
}

func meetsRequirements(path []string) bool {
	hasDac := false
	hasFft := false
	for _, server := range path {
		if server == "dac" {
			hasDac = true
		}
		if server == "fft" {
			hasFft = true
		}
	}

	return hasDac && hasFft
}

func (c *Center) PopulateCenter(data []string) error {
	servers := []*Server{}
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

	// Create out server
	outServer := NewServer("out")
	servers = append(servers, outServer)
	c.servers = servers

	for serverName, outputs := range serverLinks {
		server, err := c.Server(serverName)
		if err != nil {
			spew.Dump(fmt.Errorf("Error while getting server %s, %v", serverName, err))
			return err
		}
		links := []*Link{}
		for _, output := range outputs {
			link := NewLink(serverName, output)
			links = append(links, link)
		}
		server.AddLinks(links)
	}

	return nil
}
