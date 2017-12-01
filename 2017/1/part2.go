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
var matchSlice = []int{}
var sum int

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
	for i, r := range inputStr {
		n, _ := strconv.Atoi(string(r))
		inputArr[i] = n
	}

	steps := length / 2

	for i, n := range inputArr {
		matchIndex := i + steps
		if matchIndex > length-1 {
			matchIndex = matchIndex - length
		}
		match := inputArr[matchIndex]

		/*
			fmt.Println("---->")
			fmt.Println("i", i)
			fmt.Println("n", n)
			fmt.Println("matchIndex", matchIndex)
			fmt.Printf("inputArr[%d] %d\n", matchIndex, inputArr[matchIndex])
			fmt.Println("<----")
		*/

		if n == match {
			matchSlice = append(matchSlice, n)
		}
	}

	for _, n := range matchSlice {
		sum += n
	}
	fmt.Println("steps =>", steps)

	fmt.Println("sum =>", sum)
}
