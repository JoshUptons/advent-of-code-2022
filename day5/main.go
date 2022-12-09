package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Move struct {
	Amount int
	From   int
	To     int
}

func (m *Move) log(i int) {
	fmt.Printf("move #%d: moving %d crates from %d to %d\n", i, m.Amount, m.From, m.To)
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func parseMoves(lines []string) []Move {
	moves := []Move{}
	for _, line := range lines {
		if len(line) > 0 {
			if line[:4] == "move" {
				re := regexp.MustCompile(`\w+ (\d+)`)
				found := re.FindAllStringSubmatch(line, -1)
				commands := []int{}
				for _, n := range found {
					i, err := strconv.Atoi(n[1])
					handleErr(err)
					commands = append(commands, i)
				}
				newMove := Move{commands[0], commands[1], commands[2]}
				moves = append(moves, newMove)
			}
		}
	}
	return moves
}

func parseStacks(lines []string) map[int][]string {
	stacks := map[int][]string{}
	for _, line := range lines {
		if line[:4] == " 1  " {
			break
		}

		numStacks, remainder := len(line)/4, len(line)%4
		if remainder > 0 {
			numStacks++
		}
		pos := 0
		for i := 0; i < numStacks; i++ {
			if line[pos:pos+3] != "   " {
				stacks[i+1] = append(stacks[i+1], strings.Trim(line[pos:pos+3], "[] "))
			}
			pos += 4
		}
	}
	return stacks
}

func moveCrates(s *map[int][]string, move Move) bool {
	if len((*s)[move.From]) < move.Amount {
		return false
	}
	for i := 0; i < move.Amount; i++ {
		x := (*s)[move.From][0]
		(*s)[move.From] = append([]string{}, (*s)[move.From][1:]...)
		(*s)[move.To] = append([]string{x}, (*s)[move.To]...)
	}
	return true
}

func getTopCrates(s map[int][]string) string {
	r := ""
	for i := 1; i <= len(s); i++ {
		r += s[i][0]
	}
	return r
}

func main() {

	file, err := os.ReadFile("input.txt")
	handleErr(err)

	lines := strings.Split(string(file), "\n")

	// parse the stacks
	stacks := parseStacks(lines)
	moves := parseMoves(lines)

	for i, move := range moves {
		move.log(i)
		ok := moveCrates(&stacks, move)
		if !ok {
			fmt.Println("Error completing move")
			os.Exit(1)
		}
	}

	fmt.Println(getTopCrates(stacks))

}
