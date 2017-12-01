package main

import (
	"fmt"

	"github.com/loganbickmore/advent-of-code/utils"
)

var inputStr string
var length int

func main() {
	inputStr = utils.GetInput()
	length = len(inputStr)

	var inputArr = make([]int, length)
	var sumSlice = []int{}
	var sum int

	// convert string to array
	for i, r := range inputStr {
		n := utils.Atoi(string(r))
		inputArr[i] = n
	}

	steps := length / 2

	// loop array  to get matches into a slice
	for i, n := range inputArr {

		// setup index of array using steps
		matchIndex := i + steps
		if matchIndex > length-1 {
			matchIndex = matchIndex - length
		}

		match := inputArr[matchIndex]

		// if good add to sumSlice
		if n == match {
			sumSlice = append(sumSlice, n)
		}
	}

	// generate final sum by adding everything in the sumSlice
	for _, n := range sumSlice {
		sum += n
	}
	fmt.Println("sum =>", sum)
}
