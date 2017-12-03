package main

import (
	"fmt"
	"math"
)

var input = 277678
var targetCoords = []int{}
var centerCoords = []int{}

/*
37, 36, 35, 34, 33, 32, 31,
38, 17, 16, 15, 14, 13, 30,
39, 18,  5,  4,  3, 12, 29,
40, 19,  6,  1,  2, 11, 28,
41, 20,  7,  8,  9, 10, 27,  ^
42, 21, 22, 23, 24, 25, 26,  |
43, 44, 45, 46, 47, 48, 49, 50,
*/
var test = [][]int{
	[]int{1},
	[]int{2, 3, 4, 5, 6, 7, 8, 9},
	[]int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25},
	makeRange(26, 49),
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func genMatrix(n int) [][]int {
	m := [][]int{}
	for i := 0; i < n; i++ {
		mN := []int{}
		for j := 0; j < n; j++ {
			// x stores the layer in which (i,j)th element lies
			var x int

			// find minimum of four inputs (i, j) and (n-1-i, n-1-j)
			x = min(min(i, j), min(n-1-i, n-1-j))

			if i <= j { // upper right
				num := (n-2*x)*(n-2*x) - (i - x) - (j - x)
				//fmt.Printf("%2d ", num)
				if num == input {
					targetCoords = []int{i, j}
				}
				if num == 1 {
					centerCoords = []int{i, j}
				}
				mN = append(mN, num)

			} else { // lower left
				num := (n-2*x-2)*(n-2*x-2) + (i - x) + (j - x)
				if num == input {
					targetCoords = []int{i, j}
				}
				if num == 1 {
					centerCoords = []int{i, j}
				}
				//fmt.Printf("%2d ", num)
				mN = append(mN, num)
			}
		}
		m = append(m, mN)
		//fmt.Println()
	}
	return m
}
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
func calcSteps(n int) int {
	sqrt := int(math.Sqrt(float64(n)))

	return sqrt % n
}

func main() {
	genMatrix(530)
	//fmt.Println(m[coords[0]][coords[1]])

	distance := (targetCoords[0] - centerCoords[0]) + (targetCoords[1] - centerCoords[1])
	fmt.Println(distance)
	return
	num := 5
	l := 2
	for i, draw := range spiral(num) {
		fmt.Printf("%*d ", l, draw)
		if i%num == num-1 {
			fmt.Println("")
		}
	}

	return
	for _, v := range test {
		fmt.Println(len(v))
	}
	fmt.Println(20%8, 20/8)

	return
	testnum := 8
	total := 0
	for i := 1; i <= testnum; i++ {
		total += i
	}
	fmt.Println("total", total)
	fmt.Println("results", total%testnum/2)

	fmt.Println("test 1 expect 0 ->", calcSteps(1))
	fmt.Println("test 12 expect 3 ->", calcSteps(12))
	fmt.Println("test 23 expect 2 ->", calcSteps(23))
	fmt.Println("test 1024 expect 31 ->", calcSteps(1024))

	//input := readTo2dSlice("input")
	//fmt.Println("Part1 ->", solve(input, diffMaxMin))
	//fmt.Println("Part2 ->", solve(input, div))
}

/// example from http://www.golangprograms.com/golang-program-to-print-a-matrix-in-spiral-form.html

// Draw Spiral in Golang
func spiral(n int) []int {
	left, top, right, bottom := 0, 0, n-1, n-1
	sz := n * n
	s := make([]int, sz)
	i := 0
	for left < right {
		// work right, along top
		for c := left; c <= right; c++ {
			s[top*n+c] = i
			i++
		}
		top++
		// work down right side
		for r := top; r <= bottom; r++ {
			s[r*n+right] = i
			i++
		}
		right--
		if top == bottom {
			break
		}
		// work left, along bottom
		for c := right; c >= left; c-- {
			s[bottom*n+c] = i
			i++
		}
		bottom--
		// work up left side
		for r := bottom; r >= top; r-- {
			s[r*n+left] = i
			i++
		}
		left++
	}
	// center (last) element
	s[top*n+left] = i

	return s
}
