package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var array [][]int

	sum := 0

	var combinedStr string

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

	for _, line := range lines {
		var digits []int
		for _, ch := range line {
			digits = append(digits, int(ch-'0'))
		}
		array = append(array, digits)
	}

	for i := 0; i < len(array); i++ {

		biggestNum := -1

		secondBiggestNumb := -1

		biggestIndex := -1

		secondBiggestIndex := -1

		for j := 0; j < len(array[i])-1; j++ {

			current := array[i][j]

			if current > biggestNum {
				biggestNum = current
				biggestIndex = j
			}
		}

		for k := biggestIndex + 1; k < len(array[i]); k++ {

			current := array[i][k]

			if secondBiggestNumb < current {

				secondBiggestNumb = current
				secondBiggestIndex = k
			}

		}

		if biggestIndex < secondBiggestIndex {
			combinedStr = fmt.Sprintf("%d%d", biggestNum, secondBiggestNumb)
		} else {
			combinedStr = fmt.Sprintf("%d%d", secondBiggestNumb, biggestNum)
		}

		combinedNum, err := strconv.Atoi(combinedStr)
		if err != nil {
			panic(err)
		}

		sum += combinedNum

	}

	fmt.Println(sum)
}
