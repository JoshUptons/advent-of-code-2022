package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

type tree struct {
	value    int
	visible  bool
	position struct {
		x int
		y int
	}
}

func main() {
	//file, err := os.ReadFile("input.txt")
	file, err := os.ReadFile("input.txt")
	handleErr(err)

	stringLines := strings.Split(string(file), "\n")

	grid := map[int][]tree{}

	// for each tree, calculate the viewing score of each tree
	// view score is calculated by multiplying each distance that the tree can see

	for i := 0; i < len(stringLines); i++ {
		trees := strings.Split(stringLines[i], "")
		for j := 0; j < len(trees); j++ {
			n, err := strconv.Atoi(trees[j])
			handleErr(err)
			grid[i] = append(grid[i], tree{n, false, struct {
				x int
				y int
			}{j, i}})
		}
	}

	highestViewScore := 0
	viewTree := &tree{}

	for i := 1; i <= len(grid)-2; i++ {
		for j := 1; j <= len(grid[i])-2; j++ {

			var (
				t     *tree = &grid[i][j]
				up    int   = 0
				right       = 0
				down        = 0
				left        = 0
			)

			// look up
			for n := i - 1; n >= 0; n-- {
				up++
				if t.value == grid[n][j].value {
					break
				}
			}

			// look right
			for n := j + 1; n <= len(grid[i])-1; n++ {
				right++
				if t.value <= grid[i][n].value {
					break
				}
			}

			// look down
			for n := i + 1; n <= len(grid)-1; n++ {
				down++
				if t.value <= grid[n][j].value {
					break
				}
			}

			// look left
			for n := j - 1; n >= 0; n-- {
				left++
				if t.value <= grid[i][n].value {
					break
				}
			}

			calc := up * down * left * right
			if calc > highestViewScore {
				highestViewScore = calc
				viewTree = &grid[i][j]
			}

		}
	}

	fmt.Printf("highest score tree at grid: [%d, %d], value: %d\nview score: %d\n", viewTree.position.x, viewTree.position.y, viewTree.value, highestViewScore)

}
