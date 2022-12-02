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
		"X": 1,
		"Y": 2,
		"Z": 3,
	}
	total_score := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		turn := strings.Fields(line)
		enemy_play, my_play := plays[turn[0]], plays[turn[1]]
		switch enemy_play - my_play {
		case 0:
			total_score += my_play + 3 // draw
		case -1, 2:
			total_score += my_play + 6 // win
		case 1, -2:
			total_score += my_play // loss
		}

	}

	fmt.Println(total_score)

	input.Close()
}
