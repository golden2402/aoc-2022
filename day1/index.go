package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func assert(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, e := os.ReadFile("./sample.txt")
	assert(e)
	// close file at the end of main execution:

	elves := strings.Split(string(file), "\n\n")
	elves_count := len(elves)

	elves_calories := make([]int, elves_count)

	for i, elf := range elves {
		sum := 0

		for _, calories := range strings.Split(elf, "\n") {
			value, e := strconv.Atoi(calories)
			assert(e)

			sum += value
		}

		elves_calories[i] = sum
	}

	// sort first:
	sort.Ints(elves_calories)

	// part 1:
	fmt.Println("Most calories:", elves_calories[elves_count-1])

	// part 2:
	sum := 0
	for i := elves_count - 1; i >= elves_count-3; i-- {
		sum += elves_calories[i]
	}

	fmt.Println("Top 3 total calories:", sum)
}
