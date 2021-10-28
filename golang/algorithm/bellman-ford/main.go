package main

import (
	"fmt"
	"math"
)

func main() {
	graph := make(map[string]map[string]float32)
	graph["A"] = map[string]float32{
		"B": 5,
		"C": 0,
	}
	graph["B"] = map[string]float32{
		"C": -7,
	}
	graph["C"] = map[string]float32{
		"D": 5,
	}
	graph["D"] = map[string]float32{}
	distance, parent := bellmanFord(graph, "A")
	fmt.Println(distance)
	fmt.Println(parent)
}

func bellmanFord(graph map[string]map[string]float32, source string) (distance map[string]float32, parent map[string]string) {
	distance = make(map[string]float32)
	parent = make(map[string]string)
	// 添加当前点到所有节点权重
	// 权重为默认类型最大值
	for fromName := range graph {
		distance[fromName] = math.MaxFloat32
		parent[fromName] = ""
	}
	// 设置自身到自身的权重
	distance[source] = 0

	for i := 0; i < len(graph)-1; i++ {
		for fromName := range graph {
			for toName, weight := range graph[fromName] {
				currentWeight := distance[fromName] + weight
				if distance[toName] > currentWeight {
					// 记录最短距离 > 启始到父节点距离 + 父节点到当前子节点距离
					distance[toName] = currentWeight
					parent[toName] = fromName
				}
			}
		}
	}

	// 负权环检查
	for fromName := range graph {
		for toName := range graph[fromName] {
			if distance[toName] > distance[fromName] + graph[fromName][toName] {
				return nil, nil
			}
		}
	}
	return
}
