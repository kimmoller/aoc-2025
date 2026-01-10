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

func (c *Center) FindAllPaths(start, end string) ([][]string, error) {
	allPaths := [][]string{}
	startingServer, err := c.Server(start)
	if err != nil {
		return nil, err
	}
	for _, link := range startingServer.links {
		path := []string{startingServer.name}
		output := link.output
		err := c.FindPath(&allPaths, path, output, end)
		if err != nil {
			return nil, err
		}
	}

	return allPaths, nil
}

func (c *Center) FindPath(finalPaths *[][]string, currentPath []string, nextServerName string, finalServerName string) error {
	path := currentPath
	path = append(path, nextServerName)

	server, err := c.Server(nextServerName)
	if err != nil {
		return err
	}

	// This needs to somehow handle the branching paths
	// We cannot just return here. Instead it has to create a new path per new link and somehow return all paths once they are done
	for _, link := range server.links {
		output := link.output
		if output == finalServerName {
			path = append(path, output)
			*finalPaths = append(*finalPaths, path)
		} else {
			err := c.FindPath(finalPaths, path, output, finalServerName)
			if err != nil {
				return err
			}
		}
	}

	return nil
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
