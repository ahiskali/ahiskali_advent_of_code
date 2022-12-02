package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	max := 0
	sum := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) == 0 {
			if sum > max {
				max = sum
			}
			sum = 0
			continue
		}
		amount, _ := strconv.Atoi(line)

		sum += amount
	}

	fmt.Println(max)

	input.Close()
}
