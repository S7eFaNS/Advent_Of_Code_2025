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

		if direction[0] == 'L' {

			startNumber = (startNumber - value + 100) % 100

			if startNumber == 0 {
				countOfZeros++
			}

		} else if direction[0] == 'R' {

			startNumber = (startNumber + value) % 100

			if startNumber == 0 {
				countOfZeros++
			}

		}

	}

	fmt.Println(countOfZeros)
}
