package main

import "fmt"

const maxInt = int(^uint(0) >> 1)

type adjacentVertex struct{ vertex, distance int }
type graph [][]adjacentVertex

func createGraph() graph {
	g := make(graph, 9)
	g[0] = []adjacentVertex{{1, 4}, {7, 8}}
	g[1] = []adjacentVertex{{0, 4}, {2, 8}, {7, 11}}
	g[2] = []adjacentVertex{{1, 8}, {3, 7}, {8, 2}}
	g[3] = []adjacentVertex{{2, 7}, {4, 9}, {5, 14}}
	g[4] = []adjacentVertex{{3, 9}, {5, 10}}
	g[5] = []adjacentVertex{{3, 14}, {4, 10}, {6, 2}}
	g[6] = []adjacentVertex{{5, 2}, {7, 1}, {8, 6}}
	g[7] = []adjacentVertex{{0, 8}, {1, 11}, {6, 1}, {8, 7}}
	g[8] = []adjacentVertex{{2, 2}, {6, 6}, {7, 7}}
	return g
}

func (g graph) edgeLength(v1, v2 int) int {
	for _, neighbor := range g[v1] {
		if neighbor.vertex == v2 {
			return neighbor.distance
		}
	}
	return maxInt
}

func findShortestPaths(g graph, startVertex int) {
	distances := make([]int, len(g))
	for i := range distances {
		distances[i] = maxInt
	}
	distances[startVertex] = 0

	visited := make(map[int]bool)
	for len(visited) < len(g) {
		minDistance := maxInt
		index := -1
		for i, distance := range distances {
			if !visited[i] && distance < minDistance {
				minDistance = distance
				index = i
			}
		}
		visited[index] = true

		for _, neighbor := range g[index] {
			newDistance := minDistance + g.edgeLength(index, neighbor.vertex)
			if distances[neighbor.vertex] == maxInt || newDistance < distances[neighbor.vertex] {
				distances[neighbor.vertex] = newDistance
			}
		}
	}

	for i, d := range distances {
		fmt.Printf("Min distance from %d to %d = %d\n", startVertex, i, d)
	}
}

func main() {
	g := createGraph()
	findShortestPaths(g, 0)
}
