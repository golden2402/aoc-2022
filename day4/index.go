package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func assert(e error) {
	if e != nil {
		panic(e)
	}
}

// part 1:
func is_contained(a_min int, a_max int, b_min int, b_max int) bool {
	return b_min >= a_min && b_max <= a_max || a_min >= b_min && a_max <= b_max
}

// part 2:
func is_overlapping(a_min int, a_max int, b_min int, b_max int) bool {
	return a_max >= b_min && b_max >= a_min
}

// for strings.FieldsFunc:
func split(r rune) bool {
	return r == '-' || r == ','
}

func main() {
	file_bytes, e := io.ReadAll(os.Stdin)
	assert(e)

	content := strings.Split(
		strings.ReplaceAll(string(file_bytes), "\r", ""), "\n")

	contain := 0
	overlap := 0

	for _, line := range content {
		var ranges []int

		for _, field := range strings.FieldsFunc(line, split) {
			value, e := strconv.Atoi(field)
			assert(e)

			ranges = append(ranges, value)
		}

		// if only a variadic applied here...
		if is_contained(ranges[0], ranges[1], ranges[2], ranges[3]) {
			contain++
		}

		if is_overlapping(ranges[0], ranges[1], ranges[2], ranges[3]) {
			overlap++
		}
	}

	// output:
	fmt.Println(contain)
	fmt.Println(overlap)
}
