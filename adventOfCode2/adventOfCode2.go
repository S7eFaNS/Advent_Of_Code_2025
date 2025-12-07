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

			if len(s)%2 != 0 {
				continue
			} else {
				half := len(s) / 2
				first := s[:half]
				last := s[half:]

				if first == last {
					number, err := strconv.ParseInt(s, 10, 64)
					if err != nil {
						panic(err)
					}
					matches += number
				} else {
					continue
				}
			}
		}
	}

	fmt.Println(matches)
}
