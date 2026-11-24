package main

import "fmt"

// MaxAreaOfIsland returns the cell count of the largest 4-connected group
// of 1s. Number-of-islands with the DFS returning its region's size.
// The grid is consumed (islands are sunk); pass a copy to preserve it.
func MaxAreaOfIsland(grid [][]int) int {
	best := 0
	var sink func(r, c int) int
	sink = func(r, c int) int {
		if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[r]) || grid[r][c] != 1 {
			return 0
		}
		grid[r][c] = 0
		return 1 + sink(r+1, c) + sink(r-1, c) + sink(r, c+1) + sink(r, c-1)
	}
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == 1 {
				if area := sink(r, c); area > best {
					best = area
				}
			}
		}
	}
	return best
}

func main() {
	grid := [][]int{
		{0, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 0, 0, 1},
	}
	fmt.Println(MaxAreaOfIsland(grid))
}
