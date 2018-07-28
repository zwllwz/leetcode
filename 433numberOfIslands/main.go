package main

import "fmt"

// Coordinate is an unexported type
type Coordinate struct {
	X int
	Y int
}

func numIslands(grid [][]bool) int {
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	n := len(grid)
	m := len(grid[0])
	islands := 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] {
				markByBFS(grid, i, j)
				islands++
			}
		}
	}
	return islands
}

func markByBFS(grid [][]bool, x int, y int) {
	directionX := []int{0, 1, -1, 0}
	directionY := []int{1, 0, 0, -1}

	queue := []Coordinate{}

	queue = append(queue, Coordinate{x, y})
	grid[x][y] = false

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for i := 0; i < 4; i++ {
			adj := Coordinate{node.X + directionX[i], node.Y + directionY[i]}
			if !inBound(adj, grid) {
				continue
			}

			if grid[adj.X][adj.Y] {
				grid[adj.X][adj.Y] = false
				queue = append(queue, adj)
			}
		}
	}

}

func inBound(adj Coordinate, grid [][]bool) bool {
	n := len(grid)
	m := len(grid[0])

	return adj.X >= 0 && adj.X < n && adj.Y >= 0 && adj.Y < m
}

func main() {
	grid := [][]bool{
		{true, true, false, false, false},
		{false, true, false, false, true},
		{false, false, false, true, true},
		{false, false, false, false, false},
		{false, false, false, false, true},
	}
	result := numIslands(grid)
	fmt.Println(result)
}
