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

// Max returns the greater value between x and y.
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
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
