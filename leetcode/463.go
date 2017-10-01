package main

import "fmt"

func main() {
	fmt.Println(islandPerimeter([][]int{{1}, {0}, {0}}))
}

func islandPerimeter(grid [][]int) int {
	perimeter := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				perimeter += 4
				if i-1 >= 0 && grid[i-1][j] == 1 {
					perimeter-=2
				}
				if j-1 >= 0 && grid[i][j-1] == 1 {
					perimeter-=2
				}
			}
		}
	}
	return perimeter
}
