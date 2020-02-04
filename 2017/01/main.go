package main

import "github.com/loganbickmore/advent-of-code/utils"

func arrAtoi(s string) []int {
	var arr = make([]int, len(s))
	for i, r := range s {
		arr[i] = utils.Atoi(string(r))
	}
	return arr
}

func solve(arr []int, x int) int {
	var sum int
	list := append(arr, arr...)
	for i, n := range arr {
		if n == list[i+x] {
			sum += n
		}
	}
	return sum
}

func main() {
	var input = arrAtoi(utils.GetInput())
	println("Part1 ->", solve(input, 1))
	println("Part2 ->", solve(input, len(input)/2))
}
