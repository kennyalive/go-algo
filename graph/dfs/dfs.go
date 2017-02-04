package main

import "fmt"

type graph [][]int

func createGraph() graph {
	g := make([][]int, 4)
	g[0] = []int{1, 2}
	g[1] = []int{2}
	g[2] = []int{0, 3}
	g[3] = []int{3}
	return g
}

func doDfsTraversal(g graph, startVertex int) {
	seen := make(map[int]bool)
	var visit func(int)
	visit = func(vertex int) {
		if !seen[vertex] {
			seen[vertex] = true
			fmt.Println(vertex)
			for _, v := range g[vertex] {
				visit(v)
			}
		}
	}
	visit(startVertex)
}

func main() {
	g := createGraph()
	doDfsTraversal(g, 2)
}
