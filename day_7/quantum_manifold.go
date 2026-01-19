package main

import (
	"strconv"
	"strings"
)

type Node struct {
	id    string
	point Point
	links []Node
}

type Point struct {
	x int
	y int
}

func parseNodes(data []string) ([]Node, map[int][]Node) {
	nodes := []Node{}
	structure := map[int][]Node{}
	for i, row := range data {
		items := strings.Split(row, "")
		for j, item := range items {
			if item == START || item == SPLITTER {
				id := strconv.Itoa(i) + ":" + strconv.Itoa(j)
				point := Point{x: j, y: i}
				node := Node{id: id, point: point}
				if _, ok := structure[j]; ok {
					structure[j] = append(structure[j], node)
				} else {
					structure[j] = []Node{node}
				}
				nodes = append(nodes, node)
			}
		}
	}
	return nodes, structure
}

func buildLinks(nodes []Node, structure map[int][]Node) (Node, map[string]Node) {
	finalNodes := map[string]Node{}
	var startingNode Node
	for _, node := range nodes {
		// Starting node
		if node.point.y == 0 {
			closestNode := findClosestNode(node, node.point.x, structure)
			if closestNode != nil {
				node.links = append(node.links, *closestNode)
				finalNodes[node.id] = node
				startingNode = node
			}
		} else {
			closestLeftNode := findClosestNode(node, node.point.x-1, structure)
			if closestLeftNode != nil {
				node.links = append(node.links, *closestLeftNode)
			}
			closestRightNode := findClosestNode(node, node.point.x+1, structure)
			if closestRightNode != nil {
				node.links = append(node.links, *closestRightNode)
			}
			finalNodes[node.id] = node
		}
	}
	return startingNode, finalNodes
}

func findClosestNode(node Node, xPosition int, structure map[int][]Node) *Node {
	if sameColumn, ok := structure[xPosition]; ok {
		for _, nextNode := range sameColumn {
			if nextNode.point.y > node.point.y {
				return &nextNode
			}
		}
	}
	return nil
}

func findTimelinesWithDFS(currentNode Node, nodes map[string]Node, visited map[string]int) int {
	if len(currentNode.links) == 0 {
		return 2
	}

	if count, ok := visited[currentNode.id]; ok {
		return count
	}

	count := 0
	for _, link := range currentNode.links {
		if node, ok := nodes[link.id]; ok {
			value := findTimelinesWithDFS(node, nodes, visited)
			count += value
		}
	}
	// Need to add a split when there is only one neighbour as every split is always two
	// timelines no matter what.
	// Exception being the starting node as there is no split yet
	if currentNode.point.y != 0 && len(currentNode.links) == 1 {
		count++
	}
	visited[currentNode.id] = count
	return count
}
