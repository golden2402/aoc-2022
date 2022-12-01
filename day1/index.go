package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func assert(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, e := os.Open("./sample.txt")
	assert(e)
	// close file at the end of main execution:
	defer file.Close()

	scanner := bufio.NewScanner(file)

	elves_slice := []int{0}
	elf_i := 0

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) > 0 {
			value, e := strconv.Atoi(line)
			assert(e)

			elves_slice[elf_i] += value
		} else {
			elf_i += 1
			elves_slice = append(elves_slice, 0)
		}
	}

	// sort first:
	sort.Ints(elves_slice)
	elves_count := len(elves_slice)

	// part 1:
	fmt.Println("Most calories:", elves_slice[elves_count-1])

	// part 2:
	sum := 0
	for i := elves_count - 1; i >= elves_count-3; i-- {
		sum += elves_slice[i]
	}

	fmt.Println("Top 3 total calories:", sum)
}
