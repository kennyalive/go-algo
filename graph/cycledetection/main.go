package main

import "fmt"

type graph [][]int

func createGraph() graph {
	g := make([][]int, 3)
	g[0] = []int{1, 2}
	g[1] = []int{0, 2}
	g[2] = []int{0, 1}
	return g
}

func createGraph2() graph {
	g := make([][]int, 4)
	g[0] = []int{1, 3}
	g[1] = []int{0, 2}
	g[2] = []int{1}
	g[3] = []int{0}
	return g
}

func createGraph3() graph {
	g := make([][]int, 4)
	g[0] = []int{1, 2, 3}
	g[1] = []int{0, 2}
	g[2] = []int{1}
	g[3] = []int{0}
	return g
}

func createGraph4() graph {
	g := make([][]int, 3)
	g[0] = []int{1}
	g[1] = []int{0, 2}
	g[2] = []int{1}
	return g
}

func hasCycle(g graph) bool {
	vertexCount := len(g)
	parents := make([]int, vertexCount)
	for i := range parents {
		parents[i] = -1
	}

	getRoot := func(v int) int {
		for parents[v] != -1 {
			v = parents[v]
		}
		return v
	}

	for v1 := range g {
		for _, v2 := range g[v1] {
			// since we have undirected graph this prevents from visiting
			// each edge twice
			if v1 >= v2 {
				continue
			}

			root1 := getRoot(v1)
			root2 := getRoot(v2)
			if root1 == root2 {
				return true
			}
			parents[root1] = root2
		}
	}
	return false
}

func main() {
	graphs := []graph{createGraph(), createGraph2(), createGraph3(), createGraph4()}
	for i, g := range graphs {
		if hasCycle(g) {
			fmt.Printf("Graph %d has a cycle\n", i)
		} else {
			fmt.Printf("Graph %d has no cycles\n", i)
		}
	}
}
