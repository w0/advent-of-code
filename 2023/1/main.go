package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

func sum(s []int) int {
	sum := 0

	for _, value := range s {
		sum += value
	}
	return sum
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var numbers []int

	for scanner.Scan() {
		stringBytes := []byte(scanner.Text())

		var digits []byte

		for _, rune := range stringBytes {
			if rune > 47 && rune < 58 {
				digits = append(digits, rune)
				break
			}
		}
		slices.Reverse(stringBytes)
		for _, rune := range stringBytes {
			if rune > 47 && rune < 58 {
				digits = append(digits, rune)
				break
			}
		}
		val, err := strconv.Atoi(string(digits))

		if err != nil {
			log.Fatal(err)
		}

		numbers = append(numbers, val)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Sum: %d\n", sum(numbers))
}
