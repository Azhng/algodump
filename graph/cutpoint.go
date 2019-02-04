package main

import (
	"fmt"
)

func dfs(graph map[int][]int,
	visited map[int]bool,
	discovered map[int]int,
	parent map[int]int,
	lowTime map[int]int,
	result *[]int,
	depth int,
	vtx int) {

	discovered[vtx] = depth
	lowTime[vtx] = depth
	visited[vtx] = true

	isCutPoint := false
	childCount := 0

	for _, adj := range graph[vtx] {
		// skip parent
		if p, ok := parent[vtx]; ok && p == adj {
			continue
		}

		// if haven't visited this place
		if _, ok := visited[adj]; !ok {
			childCount++
			parent[adj] = vtx
			dfs(graph, visited, discovered, parent, lowTime, result, depth+1, adj)

			if discovered[vtx] <= lowTime[adj] {
				isCutPoint = true
			} else {
				if lowTime[adj] < lowTime[vtx] {
					lowTime[vtx] = lowTime[adj]
				}
			}
		} else {
			if lowTime[adj] < lowTime[vtx] {
				lowTime[vtx] = lowTime[adj]
			}
		}
	}

	if _, ok := parent[vtx]; !ok && childCount >= 2 {
		*result = append(*result, vtx)
		return
	}

	if _, ok := parent[vtx]; ok && isCutPoint {
		*result = append(*result, vtx)
		return
	}
}

func cutPoint(graph map[int][]int) []int {
	visited := make(map[int]bool)
	discovered := make(map[int]int)
	parent := make(map[int]int)
	lowTime := make(map[int]int)
	result := make([]int, 0)

	var start int
	for k, _ := range graph {
		start = k
		break
	}

	dfs(graph, visited, discovered, parent, lowTime, &result, 0, start)

	fmt.Println("start: ", start)
	fmt.Println("visited: ", visited)
	fmt.Println("discovered: ", discovered)
	fmt.Println("parent: ", parent)
	fmt.Println("lowTime: ", lowTime)
	fmt.Println("result: ", result)

	return result
}

func foo(a *([]int)) {
	*a = append(*a, 1243)
}

func main() {
	//graph := map[int][]int{
	//	1: []int{2},
	//	2: []int{1, 3},
	//	3: []int{2},
	//}
	graph := map[int][]int{
		1: []int{2, 6},
		2: []int{1, 3, 6},
		3: []int{2, 4, 5},
		4: []int{3, 5},
		5: []int{3, 4},
		6: []int{1, 2},
	}

	fmt.Println(graph)
	fmt.Println(cutPoint(graph)) // should be [2, 3]
}
