package main

import (
	"fmt"
	"sort"
)

var (
	testIn = []int{0, 2, 7, 0}
)

func main() {
	fmt.Println(loop(testIn))

}

func loop(blocks []int) []int {
	n, i := getHighestBlock(blocks)
	// set the highest blocks location value to 0 and increment the index
	blocks[i] = 0
	i++

	// index of the highest block plus 1
	// then looping to the 0 index upon reaching i >= len(blocks)
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
		}
	}
	return blocks[index], index
}

func eqTest(a, b []int) bool {
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
