package part_1

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

	fully_contains := 0
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
		if sections[0] <= sections[2] && sections[1] >= sections[3] ||
			sections[0] >= sections[2] && sections[1] <= sections[3] {
			fully_contains += 1
		}
	}

	fmt.Println(fully_contains)

	input.Close()
}
