package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputStr string
var match = []int{}
var last int
var first int
var prev int
var sum int

func init() {
	data, _ := ioutil.ReadFile("input")
	inputStr = strings.TrimSuffix(string(data), "\n")
}

func main() {
	for i, r := range inputStr {
		n, _ := strconv.Atoi(string(r))
		if prev == n {
			match = append(match, n)
		}

		switch i {
		case 0:
			first = n
		case len(inputStr) - 1:
			last = n
		}
		prev = n
	}

	// calc first/last
	if first == last {
		match = append(match, first)
	}

	for _, n := range match {
		sum += n
	}

	fmt.Println("sum =>", sum)
}
