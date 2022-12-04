package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	overlaps := 0
	split_regex := regexp.MustCompile(`[,-]`)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		unsplit_ranges := split_regex.Split(line, 4)
		sections := make([]int, 4)
		for i, section := range unsplit_ranges {
			sections[i], err = strconv.Atoi(section)
			if err != nil {
				fmt.Println(err)
			}
		}
		var x1, x2, y1, y2 int
		if sections[0] <= sections[2] {
			x1, x2, y1, y2 = sections[0], sections[1], sections[2], sections[3]
		} else {
			y1, y2, x1, x2 = sections[0], sections[1], sections[2], sections[3]
		}

		if (x1 <= y2) && (x2 >= y1) {
			overlaps += 1
		}
	}

	fmt.Println(overlaps)

	input.Close()
}
