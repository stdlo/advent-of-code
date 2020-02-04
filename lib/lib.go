// helper package
// first based on github.com/lukechampine/advent/utils
package utils

import (
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
)

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Dist returns the distance between two points on a matrix
// https://stackoverflow.com/a/15179572
func Dist(a, b []int) int {
	return Abs(a[0]-b[0]) + Abs(a[1]-b[1])
}

// Max returns the greater value between x and y.
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Min returns the lesser value between x and y.
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// GetInput will return a string from the first
// argument passed, or from the file "./input"
func GetInput() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	}
	data, _ := ioutil.ReadFile("input")
	re := regexp.MustCompile(`\r?\n`)
	return re.ReplaceAllString(string(data), "")
}

// Itoa is a passthrough for strconv.Itoa
func Itoa(i int) string {
	return strconv.Itoa(i)
}

// Atoi is a passthrough for strconv.Atoi that panics upon failure.
func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
