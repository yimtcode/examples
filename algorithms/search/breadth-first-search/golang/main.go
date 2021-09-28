package main

import "fmt"

func main() {
	graph := make(map[string][]string)
	graph["A"] = []string{"B", "C"}
	graph["B"] = []string{"D", "E"}
	graph["C"] = []string{"F", "G"}
	graph["D"] = []string{}
	graph["E"] = []string{}
	graph["F"] = []string{}
	graph["G"] = []string{}

	dest, ok := search(graph, "A", "F")
	fmt.Println(dest, ok)
}

func search(graph map[string][]string, start, dest string) (string, bool) {
	// Check start point
	nodes, ok := graph[start]
	if !ok {
		return "", false
	}
	// Check dest point
	if _, ok := graph[dest]; !ok {
		return "", false
	}

	searchQueue := make([]string, 0, len(nodes))
	searchQueue = append(searchQueue, nodes...)
	searched := make(map[string]struct{})

	for len(searchQueue) > 0 {
		current := searchQueue[0]
		searchQueue = searchQueue[1:]

		if _, ok := searched[current]; ok {
			continue
		}

		searched[current] = struct{}{}

		if current == dest {
			return current, true
		}

		searchQueue = append(searchQueue, graph[current]...)
	}

	return "", false
}
