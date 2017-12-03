package main

import (
	"fmt"

	"github.com/loganbickmore/advent-of-code/utils"
)

var input = 277678

// genMatrix creates a matrix n*n in size and accepts a target number
// returns a 2d slice in a spiral format, target coords and center coords
func genMatrix(n int) [][]int {
	// lots of help from  http://www.geeksforgeeks.org/print-n-x-n-spiral-matrix-using-o1-extra-space/
	m := [][]int{}
	for i := 0; i < n; i++ {
		mN := []int{}
		for j := 0; j < n; j++ {
			var x int
			x = utils.Min(utils.Min(i, j), utils.Min(n-1-i, n-1-j))

			if i <= j { // upper right
				num := (n-2*x)*(n-2*x) - (i - x) - (j - x)
				//fmt.Printf("%2d ", num)
				mN = append(mN, num)

			} else { // lower left
				num := (n-2*x-2)*(n-2*x-2) + (i - x) + (j - x)
				//fmt.Printf("%2d ", num)
				mN = append(mN, num)
			}
		}
		m = append(m, mN)
		//fmt.Println()
	}
	return m
}

func getCoords(m [][]int, t int) ([]int, []int) {
	var tar, cen []int
	for i, row := range m {
		for j, n := range row {
			if n == t {
				tar = []int{i, j}
			}
			if n == 1 {
				cen = []int{i, j}
			}
		}
	}
	return tar, cen
}

func main() {
	m := genMatrix(530) //530x530 contains the target of input
	tar, cen := getCoords(m, input)
	fmt.Println("Testcases")
	fmt.Println("1: expect 0 ->", utils.Dist(getCoords(m, 1)))
	fmt.Println("12: expect 3 ->", utils.Dist(getCoords(m, 12)))
	fmt.Println("23: expect 2 ->", utils.Dist(getCoords(m, 23)))
	fmt.Println("1024: expect 31 ->", utils.Dist(getCoords(m, 1024)))

	fmt.Println("Part1 ->", utils.Dist(tar, cen))
}
