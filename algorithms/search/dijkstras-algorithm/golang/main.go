package main

import (
	"fmt"
	"math"
)

func main() {
	graph := make(map[string]map[string]uint32)

	graph["Start"] = map[string]uint32{}
	graph["Start"]["A"] = 6
	graph["Start"]["B"] = 2

	graph["A"] = map[string]uint32{}
	graph["A"]["Dest"] = 1

	graph["B"] = map[string]uint32{}
	graph["B"]["A"] = 3
	graph["B"]["Dest"] = 5

	graph["Dest"] = map[string]uint32{}

	start := "Start"
	Dest := "Dest"
	costs, parents := findShortestPath(graph, start)
	fmt.Println(costs[Dest], parents)
}

func findShortestPath(graph map[string]map[string]uint32, start string) (map[string]uint32, map[string]string) {
	costs := make(map[string]uint32)
	parents := make(map[string]string)
	processed := make(map[string]struct{})

	for node, cost := range graph[start] {
		parents[node] = start
		costs[node] = cost
	}

	lowestCostNode := findLowestCost(costs, processed)
	for len(lowestCostNode) > 0 {
		for node, cost := range graph[lowestCostNode] {
			oldCost, ok := costs[node]
			newCost := cost + costs[lowestCostNode]
			if !ok {
				// Not found direct save
				parents[node] = lowestCostNode
				costs[node] = newCost
			} else if newCost < oldCost {
				// Less that old cost
				parents[node] = lowestCostNode
				costs[node] = newCost
			}
		}

		processed[lowestCostNode] = struct{}{}
		lowestCostNode = findLowestCost(costs, processed)
	}

	return costs, parents
}

func findLowestCost(costs map[string]uint32, processed map[string]struct{}) string {
	lowestNode := ""
	var lowestCost uint32 = math.MaxUint32
	for node, cost := range costs {
		if _, ok := processed[node]; ok {
			continue
		}

		if cost < lowestCost {
			lowestNode = node
			lowestCost = cost
		}
	}

	return lowestNode
}
