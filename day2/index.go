package main

import (
	"bufio"
	"fmt"
	"os"
)

var ASCII_A = int('A')
var ASCII_X = int('X')

func assert(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	naive_score := 0 // part 1
	smart_score := 0 // part 2

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		moves := []rune(scanner.Text())

		// input:
		a := int(moves[0]) - ASCII_A
		b := int(moves[2]) - ASCII_X

		// part 1:
		switch (3 + b - a) % 3 {
		case 0: // tie:
			naive_score += 3
		case 1: // won:
			naive_score += 6
		}

		naive_score += b + 1

		// part 2:
		// X == lose, Y == draw, Z == win
		offset := b - 1
		your_move := (3 + a + offset) % 3

		smart_score += (offset+1)*3 + your_move + 1
	}

	fmt.Println("Naive score:", naive_score)
	fmt.Println("Smart score:", smart_score)
}
