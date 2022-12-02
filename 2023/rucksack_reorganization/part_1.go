package part_1

import (
	"bufio"
	"fmt"
	"os"
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

	total_priority := 0
	type void struct{}
	var member void

	for fileScanner.Scan() {
		line := fileScanner.Text()
		line_middle := len(line) / 2
		compartment_set := make(map[rune]void)
		first_compartment, second_compartment := line[:line_middle], line[line_middle:]
		for _, item := range first_compartment {
			compartment_set[item] = member
		}
		for _, item := range second_compartment {
			if _, ok := compartment_set[item]; ok {
				total_priority += priority(item)
				delete(compartment_set, item)
			}
		}
	}

	fmt.Println(total_priority)

	input.Close()
}

func priority(item rune) int {
	if item >= 'a' {
		return int(item - ('a') + 1)
	} else {
		return int(item - ('A') + 27)
	}
}
