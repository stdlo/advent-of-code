package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/loganbickmore/advent-of-code/utils"
)

// kinda a mess, reads file and parses it into a 2d slice
func readTo2dSlice(f string) [][]int {
	arr := [][]int{}
	file, _ := ioutil.ReadFile(f)
	re := regexp.MustCompile(`\r?\n`)
	s := re.Split(string(file), -1)
	for _, v := range s {
		if len(v) == 0 {
			// skip blank lines
			continue
		}
		r := strings.Fields(v)
		z := []int{}
		for _, vv := range r {
			z = append(z, utils.Atoi(vv))
		}
		arr = append(arr, z)
	}
	return arr
}

// part1
func diffMaxMin(a, b int, nums []int) int {
	small, big := nums[a], nums[b]
	for _, v := range nums {
		if v > big {
			big = v
		}
		if v < small {
			small = v
		}
	}
	return big - small
}

// part2
func div(a, b int, nums []int) int {
	big, small := sortSize(nums[a], nums[b])
	if big%small == 0 && big != small {
		return big / small
	}
	if b == len(nums)-1 {
		return div(a+1, 0, nums)
	}
	return div(a, b+1, nums)
}

// returns big, small from any two numbers
func sortSize(x, y int) (int, int) {
	if x > y {
		return x, y
	}
	return y, x
}

type fn func(int, int, []int) int

func solve(arr2d [][]int, f fn) int {
	sum := 0
	for _, nums := range arr2d {
		sum += f(1, 1, nums)
	}
	return sum
}

func main() {
	input := readTo2dSlice("input")
	fmt.Println("Part1 ->", solve(input, diffMaxMin))
	fmt.Println("Part2 ->", solve(input, div))
}
