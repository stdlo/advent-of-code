package main

import (
	"fmt"
	"sort"
	"strings"
)

type matrix [][]int

var (
	test  = []int{0, 2, 7, 0}
	brad  = []int{11, 11, 13, 7, 0, 15, 5, 5, 4, 4, 1, 1, 7, 1, 15, 11}
	input = []int{4, 1, 15, 12, 0, 9, 9, 5, 5, 8, 7, 3, 14, 5, 12, 3}
)

func main() {
	fmt.Println("test ->", loop(test, matrix{}))
	return
	fmt.Println("part1 ->", loop(input, matrix{}))
}

func loop(blocks []int, encountered matrix) int {
	// mark as previously seen
	encountered = markAsSeen(blocks, encountered)

	//fmt.Println("e", encountered)
	//fmt.Println("b", blocks)

	blocks = redistrubute(blocks)

	// check for previously seen
	if previouslySeen(blocks, encountered) {
		return len(encountered)
	}

	return loop(blocks, encountered)
}

func previouslySeen(blocks []int, encountered matrix) bool {
	for _, encounteredBlocks := range encountered {
		if IntArrEq(blocks, encounteredBlocks) {
			//fmt.Println(blocks, "evolution", i)
			//fmt.Println(encounteredBlocks)
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

func IntArrEq(a, b []int) bool {
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

func arrToS(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}
