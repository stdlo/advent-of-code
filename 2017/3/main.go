package main

import (
	"fmt"

	"github.com/loganbickmore/advent-of-code/utils"
)

var input = 277678

// PART 1 //

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
// END PART 1 //

// PART 2 //
var (
	up          = l{0, 1}
	down        = l{0, -1}
	left        = l{-1, 0}
	right       = l{1, 0}
)

type data map[l]int
type l struct {
	x int
	y int
}

// take current location and direction, return new direction
func (d data) check(lx, ly int, dir l) (l, int) {
	var cx int
	var cy int
	var chk int
	switch dir { // match current moving dir, check matching direction
	case up:
		cx = lx + left.x
		cy = ly + left.y
		chk = d[l{cx, cy}]
		if chk == 0 {
			return left, chk
		}
	case down:
		cx = lx + right.x
		cy = ly + right.y
		chk = d[l{cx, cy}]
		if chk == 0 {
			return right, chk
		}
	case left:
		cx = lx + down.x
		cy = ly + down.y
		chk = d[l{cx, cy}]
		if chk == 0 {
			return down, chk
		}
	case right:
		cx = lx + up.x
		cy = ly + up.y
		chk = d[l{cx, cy}]
		if chk == 0 {
			return up, chk
		}
	}
	return dir, chk
}

func (d data) step(lx, ly int, dir l) int {
	n := d.gen(lx, ly)            // generate location value
	dir, _ = d.check(lx, ly, dir) // set new dir
	/* Debugging prints
		fmt.Printf("chk (%d,%d) => %d\n", cx, cy, c)
		fmt.Printf("loc (%d,%d) => %d\n", lx, ly, d[l{lx, ly}])
		fmt.Printf("dir %#v\n", dir)
		fmt.Printf("orig_dir %#v\n\n", odir)
	//*/
	ly += dir.y
	lx += dir.x
	if n > input {
		return n
	}
	return d.step(lx, ly, dir)
}

func sum(input ...int) int {
	sum := 0
	for i := range input {
		sum += input[i]
	}
	return sum
}

func (d data) gen(lx, ly int) int {
	n := sum(
		d[l{lx - 1, ly + 1}], // [-1, 1]
		d[l{lx, ly + 1}],     // [ 0, 1]
		d[l{lx + 1, ly + 1}], // [ 1, 1]
		d[l{lx - 1, ly}],     // [-1, 0]
		d[l{lx + 1, ly}],     // [ 1, 0]
		d[l{lx - 1, ly - 1}], // [-1,-1]
		d[l{lx, ly - 1}],     // [ 0,-1]
		d[l{lx + 1, ly - 1}], // [ 1,-1]
	)
	d[l{lx, ly}] = n
	return n
}
// END PART 2 //

func main() {
	m := genMatrix(530) //530x530 contains the target of input
	tar, cen := getCoords(m, input)
	fmt.Println("Part1 ->", utils.Dist(tar, cen))

	d := make(data) // setup data structure
	d[l{0, 0}] = 1  // genesis coordinates
	fmt.Println("Part2 ->", d.step(1, 0, right))
}
