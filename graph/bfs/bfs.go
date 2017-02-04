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

func doBfsTraversal(g graph, vertices []int) {
	seen := make(map[int]bool)
	for len(vertices) > 0 {
		list := vertices
		vertices = nil
		for _, v := range list {
			if !seen[v] {
				seen[v] = true
				fmt.Println(v)
				vertices = append(vertices, g[v]...)
			}
		}
	}
}

func main() {
	g := createGraph()
	doBfsTraversal(g, []int{2})
}
