package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var inputStr string
var length int

func init() {
	if len(os.Args) > 1 {
		inputStr = os.Args[1]
	} else {
		data, _ := ioutil.ReadFile("input")
		inputStr = strings.TrimSuffix(string(data), "\n")
	}
	length = len(inputStr)
}

func main() {
	var inputArr = make([]int, length)
	var sumSlice = []int{}
	var sum int

	// convert string to array
	for i, r := range inputStr {
		n, _ := strconv.Atoi(string(r))
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
