package main

import (
	"fmt"

	"github.com/loganbickmore/advent-of-code/utils"
)

var input = 277678

// genMatrix creates a matrix n*n in size and accepts a target number
// returns a 2d slice in a spiral format, target coords and center coords
func genMatrix(n, t int) ([][]int, []int, []int) {
	// lots of help from  http://www.geeksforgeeks.org/print-n-x-n-spiral-matrix-using-o1-extra-space/
	m := [][]int{}
	cen := []int{}
	tar := []int{}
	for i := 0; i < n; i++ {
		mN := []int{}
		for j := 0; j < n; j++ {
			var x int
			x = utils.Min(utils.Min(i, j), utils.Min(n-1-i, n-1-j))

			if i <= j { // upper right
				num := (n-2*x)*(n-2*x) - (i - x) - (j - x)
				//fmt.Printf("%2d ", num)
				if num == t {
					tar = []int{i, j}
				}
				if num == 1 {
					cen = []int{i, j}
				}
				mN = append(mN, num)

			} else { // lower left
				num := (n-2*x-2)*(n-2*x-2) + (i - x) + (j - x)
				//fmt.Printf("%2d ", num)
				if num == input {
					tar = []int{i, j}
				}
				if num == 1 {
					cen = []int{i, j}
				}
				mN = append(mN, num)
			}
		}
		m = append(m, mN)
		//fmt.Println()
	}
	return m, tar, cen
}

func dist(a, b []int) int {
	return (a[0] - b[0]) + (a[1] - b[1])
}

func main() {
	_, tar, cen := genMatrix(530, 12)

	fmt.Println("Part1 ->", dist(tar, cen))
}
