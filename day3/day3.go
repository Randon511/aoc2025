package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func part1() {
	file, err := os.Open("input2.txt") // For read access.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()

		nums := make([]int, len(line))
		for i := 0; i < len(line); i++ {
			nums[i], _ = strconv.Atoi(string(line[i]))
		}

		length := len(nums)
		first := nums[length-2]
		second := nums[length-1]

		for i := (length - 3); i >= 0; i-- {
			if nums[i] >= first {
				if first >= second {
					second = first
				}

				first = nums[i]
			}
		}

		fmt.Printf("Line: %s, Final num: %d%d\n", line, first, second)
		sum += (first * 10) + second
	}

	fmt.Printf("Result %d", sum)
}

func part2() {
	file, err := os.Open("input2.txt") // For read access.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sum int64

	for scanner.Scan() {
		line := scanner.Text()

		input := make([]int, len(line))
		for i := 0; i < len(line); i++ {
			input[i], _ = strconv.Atoi(string(line[i]))
		}

		length := len(input)
		digits := make([]int, 12)
		for i := 0; i < 12; i++ {
			digits[i] = input[length-(12-i)]
		}

		fmt.Printf("Line: %s, Initial num:", line)
		for i := 0; i < 12; i++ {
			fmt.Printf("%d", digits[i])
		}

		for i := (length - 13); i >= 0; i-- {
			if input[i] >= digits[0] {
				temp1 := digits[0]
				for j := 1; j < 12; j++ {
					if temp1 >= digits[j] {
						temp2 := digits[j]
						digits[j] = temp1
						temp1 = temp2
					} else {
						break
					}
				}
				digits[0] = input[i]
			}
		}

		var result int64
		for _, d := range digits {
			result = result*10 + int64(d)
		}

		fmt.Printf(", Final num:%d\n", result)
		sum += result
	}

	fmt.Printf("Result %d", sum)
}

func main() {
	// part1()
	part2()
}
