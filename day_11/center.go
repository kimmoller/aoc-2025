package main

import (
	"fmt"
	"slices"
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
	startingServer, err := c.Server(start)
	if err != nil {
		return nil, err
	}
	paths, err := c.FindPath([]string{}, startingServer.name, end)
	if err != nil {
		return nil, err
	}

	finalPaths := [][]string{}
	if strict {
		for _, path := range paths {
			if meetsRequirements(path) {
				finalPaths = append(finalPaths, path)
			}
		}
	} else {
		finalPaths = paths
	}

	return finalPaths, nil
}

func (c *Center) FindPath(currentPath []string, currentServerName string, finalServerName string) ([][]string, error) {
	if currentServerName == finalServerName {
		return [][]string{currentPath}, nil
	}

	server, err := c.Server(currentServerName)
	if err != nil {
		return nil, err
	}

	knownPaths := [][]string{}
	for _, link := range server.links {
		output := link.output
		newPath := append(currentPath, output)

		outputServer, err := c.Server(output)
		if err != nil {
			return nil, err
		}

		paths := [][]string{}
		if len(outputServer.paths) > 0 {
			// Construct new paths with known paths
			startOfPath := slices.Clone(newPath)
			finalPaths := [][]string{}
			for _, path := range outputServer.paths {
				finalPath := append(startOfPath, path...)
				finalPaths = append(finalPaths, finalPath)
			}
			paths = finalPaths
		} else {
			// Unknown path
			paths, err = c.FindPath(newPath, output, finalServerName)
			if err != nil {
				return nil, err
			}
		}

		for _, path := range paths {
			knownPath := slices.Clone(path)
			knownPaths = append(knownPaths, knownPath)
		}
	}

	if len(server.links) > 1 {
		addKnownPathsToServer(server, knownPaths)
	}

	return knownPaths, nil
}

func addKnownPathsToServer(server *Server, knownPaths [][]string) {
	for _, knownPath := range knownPaths {
		copy := slices.Clone(knownPath)
		serverIndex := 0
		for i, serverName := range copy {
			if serverName == server.name {
				serverIndex = i
			}
		}
		subPath := copy[serverIndex+1:]
		server.AddPath(subPath)
	}
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
			link := NewLink(output)
			links = append(links, link)
		}
		server.AddLinks(links)
	}

	return nil
}
