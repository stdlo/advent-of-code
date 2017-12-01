package main

import "github.com/loganbickmore/advent-of-code/utils"

func toIntArr(s string) []int {
	var arr = make([]int, len(s))
	for i, r := range s {
		arr[i] = utils.Atoi(string(r))
	}
	return arr
}

func part1(arr []int) int {
	var sum int
	list := append(arr, arr...)
	for i, n := range arr {
		if n == list[i+1] {
			sum += n
		}
	}
	return sum
}

func part2(arr []int) int {
	var sum int
	list := append(arr, arr...)
	for i, n := range arr {
		if n == list[i+len(arr)/2] {
			sum += n
		}
	}
	return sum
}

func main() {
	var input = toIntArr(utils.GetInput())
	println("Part1 ->", part1(input))
	println("Part2 ->", part2(input))
}
