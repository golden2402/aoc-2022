package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

func assert(e error) {
	if e != nil {
		panic(e)
	}
}

var ascii_a = int('a')

func main() {
	scanner, e := io.ReadAll(os.Stdin)
	assert(e)

	// why the fuck is there a CR?:
	content := strings.Split(strings.ReplaceAll(string(scanner), "\r", ""), "\n")

	// part 1:
	item_sum := 0

	for _, line := range content {
		rucksack := []rune(line)
		item_count := len(rucksack)

		set_a := make(map[rune]bool)
		set_b := make(map[rune]bool)

		for i := 0; i < item_count/2; i++ {
			set_a[rucksack[i]] = true
		}

		for i := item_count / 2; i < item_count; i++ {
			set_b[rucksack[i]] = true
		}

		for k := range set_a {
			_, ok := set_b[k]

			if !ok {
				continue
			}

			// start with uppercase, work down:
			offset := 26
			// get priorities:
			item := unicode.ToLower(k)
			if item == k {
				offset = 0
			}

			item_sum += (int(item) - ascii_a + 1) + offset
		}
	}

	fmt.Println(item_sum)

	// part 2:
	badge_sum := 0

	// static array of sets:
	sets := make([]map[rune]bool, 3)

	for chunk := 0; chunk < len(content); chunk += 3 {
		// create a fresh group of sets:
		for i := 0; i < len(sets); i++ {
			sets[i] = make(map[rune]bool)
		}

		for i, line := range content[chunk : chunk+3] {
			for _, char := range []rune(line) {
				sets[i][char] = true
			}
		}

		// pick character from first map
		for k := range sets[0] {
			// check other two maps to see if they have them:
			is_common := true

			for i := 1; i < 3; i++ {
				_, ok := sets[i][k]

				// if another set does not contain the chosen character,
				// raise flag and break:
				if !ok {
					is_common = false
					break
				}
			}

			// if another set does not contain the chosen character,
			// select next:
			if !is_common {
				continue
			}

			// start with uppercase, work down:
			offset := 26
			// get priorities:
			item := unicode.ToLower(k)
			if item == k {
				offset = 0
			}

			badge_sum += (int(item) - ascii_a + 1) + offset
		}
	}

	fmt.Println(badge_sum)
}
