package main

import "fmt"

const maxInt = int(^uint(0) >> 1)

type graph [][]int

func createGraph() graph {
	g := [][]int{
		[]int{0, 5, maxInt, 10},
		[]int{maxInt, 0, 3, maxInt},
		[]int{maxInt, maxInt, 0, 1},
		[]int{maxInt, maxInt, maxInt, 0},
	}
	return g
}

func findShortestPaths(g graph) {
	// make a copy of input graph, since we are going to modify it
	d := make([][]int, len(g))
	for i, row := range g {
		d[i] = append([]int(nil), row...)
	}
	// run floyd-warshall aglorithm
	n := len(d)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if d[i][k] != maxInt && d[k][j] != maxInt {
					newDist := d[i][k] + d[k][j]
					if newDist < d[i][j] {
						d[i][j] = newDist
					}
				}
			}
		}
	}
	// print output
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if d[i][j] != maxInt {
				fmt.Printf("dist[%d][%d] = %d\n", i, j, d[i][j])
			}
		}
	}
}

func main() {
	g := createGraph()
	findShortestPaths(g)
}
