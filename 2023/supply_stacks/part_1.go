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
	crate byte
)

type Stack []interface{}

func (stack *Stack) push(el interface{}) {
	*stack = append(*stack, el)
}

func (stack *Stack) pop() interface{} {
	if len(*stack) == 0 {
		return ' '
	}

	var el interface{}
	el, *stack = (*stack)[len(*stack)-1], (*stack)[:len(*stack)-1]
	return el
}

func main() {
	input, err = os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(input)
	fileScanner.Split(bufio.ScanLines)

	crate_loading_stacks := make([]Stack, 10)
	crate_stacks := make([]Stack, 10)
	crates_loaded := false

	for fileScanner.Scan() && !crates_loaded {
		line := fileScanner.Text()
		for i := 0; i*4 < len(line); i++ {
			crate = line[i*4+1]
			if crate == '1' {
				crates_loaded = true
				break
			}
			if crate != ' ' {
				crate_loading_stacks[i].push(crate)
			}
		}
	}

	// print_stacks(crate_loading_stacks)
	for i, loading_stack := range crate_loading_stacks {
		for j := len(loading_stack) - 1; j >= 0; j-- {
			crate_stacks[i].push(loading_stack[j])
		}
		// print_stacks(crate_stacks)
	}

	pattern := regexp.MustCompile(`\d+`)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		instructions := pattern.FindAllString(line, 3)
		times, _ := strconv.Atoi(instructions[0])
		from, _ := strconv.Atoi(instructions[1])
		to, _ := strconv.Atoi(instructions[2])

		fmt.Println(line)
		for i := 0; i < times; i++ {
			crate := crate_stacks[from-1].pop()
			// if crate == ' ' {
			// 	break
			// }
			crate_stacks[to-1].push(crate)
			// print_stacks(crate_stacks)
		}
	}

	print_stacks(crate_stacks)

	input.Close()
}

func print_stacks(crate_stacks []Stack) {
	max_len := 0
	for _, stack := range crate_stacks {
		if len(stack) > max_len {
			max_len = len(stack)
		}
	}

	for height := max_len - 1; height >= 0; height-- {
		for _, stack := range crate_stacks {
			if len(stack) > height {
				fmt.Printf("[%c] ", stack[height])
			} else {
				fmt.Printf("    ")
			}
		}
		fmt.Println()
	}
	for i := 1; i < len(crate_stacks); i++ {
		fmt.Printf(" %d  ", i)
	}
	fmt.Println()
	fmt.Println()
}
