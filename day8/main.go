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
	value   int
	visible bool
}

func main() {
	file, err := os.ReadFile("input.txt")
	handleErr(err)

	stringLines := strings.Split(string(file), "\n")

	grid := map[int][]tree{}

	for i := 0; i < len(stringLines)-1; i++ {
		line := strings.Split(stringLines[i], "")
		for j := 0; j < len(line); j++ {
			n, err := strconv.Atoi(line[j])
			handleErr(err)
			grid[i] = append(grid[i], tree{n, false})
		}

	}

	totalTreesVisible := 0

	var (
		maxLeft   int = -1
		maxRight      = -1
		maxBottom     = -1
		maxTop        = -1
	)

	for i := 0; i < len(grid); i++ {

		maxRight = -1
		maxLeft = -1

		/*
			left and right
			unfortunately cannot utilize additional pointers for top and bottom
			here, as in order to reset the maxBottom and maxTop, we would need
			to do it per column checked, which would not hold value correctly
			per iteration.
		*/
		for j := 0; j < len(grid[i]); j++ {

			var (
				fromLeft  *tree = &grid[i][j]
				fromRight       = &grid[i][len(grid[i])-j-1]
			)

			if fromLeft.value > maxLeft {
				maxLeft = fromLeft.value
				if !fromLeft.visible {
					fromLeft.visible = true
					totalTreesVisible++
				}
			}

			if fromRight.value > maxRight {
				maxRight = fromRight.value
				if !fromRight.visible {
					fromRight.visible = true
					totalTreesVisible++
				}
			}

		}
		/*
			top and bottom
		*/
		for col := 0; col < len(grid[0]); col++ {
			maxTop = -1
			maxBottom = -1
			for row := 0; row < len(grid); row++ {
				var (
					fromTop    *tree = &grid[row][col]
					fromBottom       = &grid[len(grid)-1-row][col]
				)

				if fromTop.value > maxTop {
					maxTop = fromTop.value
					if !fromTop.visible {
						fromTop.visible = true
						totalTreesVisible++
					}
				}

				if fromBottom.value > maxBottom {
					maxBottom = fromBottom.value
					if !fromBottom.visible {
						fromBottom.visible = true
						totalTreesVisible++
					}
				}
			}
		}

	}

	for i := 0; i < len(grid); i++ {
		fmt.Println(grid[i])
	}

	fmt.Printf("total trees visible: %d\n", totalTreesVisible)

}
