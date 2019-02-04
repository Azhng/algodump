package main

import (
	"fmt"
	"math"
)

type pair struct {
	node   string
	weight int
}

func dfs(graph map[string]map[string]int,
	vtx string,
	dist *map[string]int,
	visited map[string]bool,
	parents map[string]string) {

	visited[vtx] = true

	for adj, weight := range graph[vtx] {
		if parents[adj] == vtx {
			continue
		}

		// haven't visited
		if _, ok := visited[adj]; !ok {
			visited[adj] = true
			if newDist := weight + (*dist)[vtx]; newDist < (*dist)[adj] {
				(*dist)[adj] = newDist
			}
			dfs(graph, adj, dist, visited, parents)
		} else {
			if newDist := weight + (*dist)[vtx]; newDist < (*dist)[adj] {
				(*dist)[adj] = newDist
			}
		}
	}

	for adj, weight := range graph[vtx] {
		if newDist := weight + (*dist)[vtx]; newDist < (*dist)[adj] {
			panic("Found negative weight cycles")
		}
	}
}

func shortestPath(graph map[string]map[string]int, node string) map[string]int {
	dist := make(map[string]int)
	visited := make(map[string]bool)
	parents := make(map[string]string)

	fmt.Println(graph)

	for k, _ := range graph {
		dist[k] = math.MaxInt32
	}

	dist[node] = 0

	dfs(graph, node, &dist, visited, parents)

	return dist
}

func main() {
	//graph := map[string]map[string]int{
	//	"A": map[string]int{
	//		"C": -2,
	//	},
	//	"C": map[string]int{
	//		"B": -3,
	//	},
	//	"B": map[string]int{
	//		"A": -1,
	//	},
	//}
	graph := map[string]map[string]int{
		"A": map[string]int{
			"B": -1,
			"C": 4,
		},
		"B": map[string]int{
			"C": 3,
			"D": 2,
			"E": 2,
		},
		"C": map[string]int{},
		"D": map[string]int{
			"C": 5,
			"B": 1,
		},
		"E": map[string]int{
			"D": -3,
		},
	}
	fmt.Println(shortestPath(graph, "A"))
}
