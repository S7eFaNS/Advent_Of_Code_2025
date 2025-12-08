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

		freq := make(map[int]int)
		for _, d := range array[i] {
			freq[d]++
		}

		for j := 0; j < len(array[i]); j++ {

			if freq[array[i][j]] != 1 {
				continue
			}

			if biggestNum < array[i][j] {

				secondBiggestNumb = biggestNum
				secondBiggestIndex = biggestIndex

				biggestNum = array[i][j]
				biggestIndex = j

				fmt.Println("first:", biggestNum)
				fmt.Println("first second:", secondBiggestNumb)

			} else if secondBiggestNumb < array[i][j] {

				secondBiggestNumb = array[i][j]
				secondBiggestIndex = j

				fmt.Println("second:", secondBiggestNumb)

			}
		}

		if biggestNum == -1 || secondBiggestNumb == -1 {
			continue
		}

		if biggestIndex < secondBiggestIndex {
			combinedStr = fmt.Sprintf("%d%d", biggestNum, secondBiggestNumb)
			fmt.Println("string", combinedStr)
		} else {
			combinedStr = fmt.Sprintf("%d%d", secondBiggestNumb, biggestNum)
		}

		combinedNum, err := strconv.Atoi(combinedStr)
		if err != nil {
			panic(err)
		}

		fmt.Println("int:", combinedNum)

		sum += combinedNum

	}

	fmt.Println(sum)
}
