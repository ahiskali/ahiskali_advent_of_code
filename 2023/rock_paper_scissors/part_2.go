package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	input *os.File
	err   error
)

func main() {
	input, err = os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(input)
	fileScanner.Split(bufio.ScanLines)

	plays := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}
	total_score := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		turn := strings.Fields(line)
		enemy_play, outcome := plays[turn[0]], turn[1]
		switch outcome[0] {
		case 'X':
			my_play := enemy_play - 1
			if my_play == 0 {
				my_play = 3
			}

			total_score += my_play + 0 // loss
		case 'Y':
			total_score += enemy_play + 3 // draw
		case 'Z':
			total_score += enemy_play%3 + 1 + 6 // win
		}

	}

	fmt.Println(total_score)

	input.Close()
}
