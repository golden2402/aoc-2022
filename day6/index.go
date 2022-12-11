package main

import (
	"fmt"
	"io"
	"os"
)

func assert(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file_bytes, e := io.ReadAll(os.Stdin)
	assert(e)

	line := string(file_bytes)

	// part 1:
	for i := 3; i < len(line); i++ {
		set := make(map[rune]bool)

		for _, char := range []rune(line[i-3 : i+1]) {
			set[char] = true
		}

		if len(set) == 4 {
			fmt.Println(i + 1)
			break
		}
	}

	// part 2:
	for i := 13; i < len(line); i++ {
		set := make(map[rune]bool)

		for _, char := range []rune(line[i-13 : i+1]) {
			set[char] = true
		}

		if len(set) == 14 {
			fmt.Println(i + 1)
			break
		}
	}
}
