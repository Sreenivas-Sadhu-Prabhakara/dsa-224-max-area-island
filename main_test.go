package main

import "testing"

func TestMaxAreaOfIsland(t *testing.T) {
	cases := []struct {
		grid [][]int
		want int
	}{
		{[][]int{{0, 1, 1, 0}, {0, 1, 0, 0}, {1, 0, 0, 1}}, 3},
		{[][]int{{0, 0}, {0, 0}}, 0},
		{[][]int{{1}}, 1},
		{[][]int{{1, 1}, {1, 1}}, 4},
		{[][]int{{1, 0, 1}, {0, 1, 0}, {1, 0, 1}}, 1}, // diagonals don't connect
		{nil, 0},
	}
	for _, c := range cases {
		if got := MaxAreaOfIsland(c.grid); got != c.want {
			t.Errorf("MaxAreaOfIsland = %d, want %d", got, c.want)
		}
	}
}
