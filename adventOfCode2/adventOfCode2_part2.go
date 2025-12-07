package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var start []int64
	var end []int64
	var matches int64

	var input string

	if scanner.Scan() {
		input = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	ranges := strings.Split(input, ",")

	for i := 0; i < len(ranges); i++ {

		if ranges[i] == "" {
			continue
		}

		bounds := strings.Split(ranges[i], "-")

		startNum, err := strconv.ParseInt(bounds[0], 10, 64)
		if err != nil {
			panic(err)
		}

		endNum, err := strconv.ParseInt(bounds[1], 10, 64)
		if err != nil {
			panic(err)
		}

		start = append(start, startNum)
		end = append(end, endNum)

	}

	for i := 0; i < len(start); i++ {
		for j := start[i]; j <= end[i]; j++ {

			s := strconv.FormatInt(j, 10)
			lenS := len(s)

			for k := 1; k <= lenS/2; k++ {

				if lenS%k != 0 {
					continue
				}

				t := s[:k]
				repeated := strings.Repeat(t, lenS/k)

				if repeated == s {
					matches += j
					break
				}
			}
		}
	}

	fmt.Println(matches)
}
