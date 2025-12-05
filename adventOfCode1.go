package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	startNumber := 50

	var direction string

	var valStr string

	var lineStr string

	countOfZeros := 0

	old := 0

	matches := 0

	scanner := bufio.NewScanner(os.Stdin)
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for i := 0; i < len(lines); i++ {
		direction = lines[i]

		lineStr = lines[i]

		valStr = lineStr[1:]

		value, err := strconv.Atoi(valStr)
		if err != nil {
			panic(err)
		}

		old = startNumber

		if direction[0] == 'L' {

			matches = (value + (100 - old)) / 100
			countOfZeros += matches

			startNumber = (startNumber - value + 100) % 100

			if matches == 0 && startNumber == 0 {
				countOfZeros++
			}

		} else if direction[0] == 'R' {

			matches = (old + value) / 100
			countOfZeros += matches

			startNumber = (startNumber + value) % 100

			if matches == 0 && startNumber == 0 {
				countOfZeros++
			}

		}

		fmt.Printf(
			"Step %d: %s%d | old=%d, value=%d, matches=%d, new=%d, count=%d\n",
			i+1, string(direction[0]), value, old, value, matches, startNumber, countOfZeros,
		)

	}

	fmt.Println(countOfZeros)
}
