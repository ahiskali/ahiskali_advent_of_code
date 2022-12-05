package main

import (
	"bufio"
	"fmt"
	"os"
)

type (
	void     struct{}
	rune_set map[rune]void
)

var (
	input  *os.File
	err    error
	member void
)

func (rset rune_set) add(key rune) {
	rset[key] = member
}

func (rset *rune_set) init(str string) {
	*rset = make(rune_set)

	for _, key := range str {
		rset.add(key)
	}
}

func (rset rune_set) intersection(other_set rune_set) rune_set {
	var iterated_set, checked_set rune_set
	intersection := make(rune_set)

	if len(rset) < len(other_set) {
		iterated_set, checked_set = rset, other_set
	} else {
		iterated_set, checked_set = other_set, rset
	}

	for key := range iterated_set {
		if _, ok := checked_set[key]; ok {
			intersection.add(key)
		}
	}

	return intersection
}

func main() {
	input, err = os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(input)
	fileScanner.Split(bufio.ScanLines)

	total_priority := 0

	for fileScanner.Scan() {
		var first_group, second_group, third_group rune_set

		first_group.init(fileScanner.Text())
		fileScanner.Scan()
		second_group.init(fileScanner.Text())
		fileScanner.Scan()
		third_group.init(fileScanner.Text())

		for badge := range first_group.intersection(second_group).intersection(third_group) {
			total_priority += priority(badge)
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
