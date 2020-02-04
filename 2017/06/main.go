package main

import (
	"fmt"
	"sort"
)

type matrix [][]int

var (
	test  = []int{0, 2, 7, 0}
	input = []int{4, 1, 15, 12, 0, 9, 9, 5, 5, 8, 7, 3, 14, 5, 12, 3}
	match = 0
)

func main() {
	fmt.Println("test ->", solve(test, matrix{}))

	results := solve(input, matrix{})
	fmt.Println("part1 ->", results)
	fmt.Println("part2 ->", results-match)
}

func solve(blocks []int, encountered matrix) int {
	encountered = markAsSeen(blocks, encountered)

	blocks = redistrubute(blocks)

	if previouslySeen(blocks, encountered) {
		return len(encountered)
	}

	return solve(blocks, encountered)
}

func previouslySeen(blocks []int, encountered matrix) bool {
	for i, encounteredBlocks := range encountered {
		if intArrEq(blocks, encounteredBlocks) {
			match = i
			return true
		}
	}
	return false
}

func markAsSeen(blocks []int, encountered matrix) matrix {
	var copyOfBlocks = make([]int, len(blocks))
	copy(copyOfBlocks, blocks)
	return append(encountered, copyOfBlocks)
}

func redistrubute(blocks []int) []int {
	n, i := getHighestBlock(blocks)
	blocks[i] = 0
	i++
	for n > 0 {
		if i >= len(blocks) {
			i = 0
		}
		blocks[i]++ //raise the value of the current location
		n--         // lower the remaining value to be distributed
		i++         // increment the index
	}
	return blocks
}

func getHighestBlock(blocks []int) (int, int) {
	var index int
	var sorted = make([]int, len(blocks))
	copy(sorted, blocks)
	sort.Ints(sorted)
	big := sorted[len(blocks)-1]
	for i, n := range blocks {
		if n == big {
			index = i
			return blocks[index], index
		}
	}
	return blocks[index], index
}

func intArrEq(a, b []int) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
