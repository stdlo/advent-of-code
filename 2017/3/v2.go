package main

import (
	"fmt"
)

/*
	[-1, 1], [0, 1], [1, 1]
	[-1, 0], [0, 0], [1, 0]
	[-1,-1], [0,-1], [1,-1]
*/
/*
	left := d[l{lx - 1, ly}]
	right := d[l{lx + 1, ly}]
	down := d[l{lx, ly - 1}]
	up := d[l{lx, ly + 1}]
*/
var (
	up    = l{0, 1}
	down  = l{0, -1}
	left  = l{-1, 0}
	right = l{1, 0}
)

type data map[l]int
type l struct {
	x int
	y int
}

// take current location and direction, return new direction
func (d data) check(lx, ly int, dir l) l {
	var cx int
	var cy int
	var c int
	switch dir { // match current moving dir, check matching direction
	case up:
		cx = lx + left.x
		cy = ly + left.y
		if d[l{cx, cy}] == 0 {
			return left
		}
	case down:
		cx = lx + right.x
		cy = ly + right.y
		if d[l{cx, cy}] == 0 {
			return right
		}
	case left:
		cx = lx + down.x
		cy = ly + down.y
		if d[l{cx, cy}] == 0 {
			return down
		}
	case right:
		cx = lx + up.x
		cy = ly + up.y
		if d[l{cx, cy}] == 0 {
			return up
		}
	}
}

func (d data) step(i, lx, ly int, dir l) (int, int) {
	d.gen(lx, ly) // generate location value
	var cx int    //:= lx + dir.x
	var cy int    //:= ly + dir.y
	var c int
	var odir = dir
	//var next = dir

	// check switch
	switch dir {
	case up:
		cx = lx + left.x
		cy = ly + left.y
		if d[l{cx, cy}] == 0 {
			dir = left
		}

	case down:
		cx = lx + right.x
		cy = ly + right.y
		if d[l{cx, cy}] == 0 {
			dir = right
		}
	case left:
		cx = lx + down.x
		cy = ly + down.y
		if d[l{cx, cy}] == 0 {
			dir = down
		}
	case right:
		cx = lx + up.x
		cy = ly + up.y
		if d[l{cx, cy}] == 0 {
			dir = up
		}
	}
	c = d[l{cx, cy}]

	// decide on dext direction
	///*
	fmt.Printf("chk (%d,%d) => %d\n", cx, cy, c)
	fmt.Printf("loc (%d,%d) => %d\n", lx, ly, d[l{lx, ly}])
	fmt.Printf("dir %#v\n", dir)
	fmt.Printf("orig_dir %#v\n\n", odir)
	//return 0, 0
	//*/

	ly += dir.y
	lx += dir.x
	i++
	if i > 5 {
		return lx, ly
	}
	if c == 0 {
		d.step(i, lx, ly, dir)
	}

	return lx, ly
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

func main() {
	d := make(data)
	d[l{0, 0}] = 1
	lx := 1
	ly := 0
	//d.gen(lx, ly)
	fmt.Println("Start", l{lx, ly}, "dir", right)
	d.step(0, lx, ly, right)
	fmt.Printf("%#v\n\n", d)
}
