package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

var (
	input *os.File
	err   error
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	input, err = os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(input)
	fileScanner.Split(bufio.ScanLines)

	max_heap := &IntHeap{}
	heap.Init(max_heap)

	sum := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) == 0 {
			heap.Push(max_heap, sum)
			sum = 0
			continue
		}
		amount, _ := strconv.Atoi(line)

		sum += amount
	}

	result := 0
	for i := 1; i <= 3; i++ {
		max := heap.Pop(max_heap).(int)
		fmt.Println(max)
		result += max
	}
	fmt.Println(result)

	input.Close()
}
