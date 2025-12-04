package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func split_string(s string, n int) map[string]struct{} {
	seg_size := len(s) / n
	result := make(map[string]struct{}, n)

	for i := 0; i < n; i++ {
		start := i * seg_size
		end := start + seg_size
		part := s[start:end]
		result[part] = struct{}{}
	}

	// fmt.Printf("Val %s, Chunk %d, len: %d, Vals: ", s, n, len(result))
	// for part := range result {
	// 	fmt.Print(part + " ")
	// }
	// fmt.Println()

	return result
}

func part1() {
	file, err := os.Open("input2.txt") // For read access.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := ""

	if scanner.Scan() {
		input = scanner.Text()
	}

	sum := 0
	ranges := strings.SplitSeq(input, ",")

	for r := range ranges {
		vals := strings.Split(r, "-")
		min_val, _ := strconv.Atoi(vals[0])
		max_val, _ := strconv.Atoi(vals[1])

		for i := min_val; i <= max_val; i++ {
			id := strconv.Itoa(i)
			if (len(id) % 2) == 0 {
				mid := len(id) / 2
				if id[:mid] == id[mid:] {
					sum += i
				}
			}
		}
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
	input := ""

	if scanner.Scan() {
		input = scanner.Text()
	}

	result_set := make(map[int]struct{})
	ranges := strings.SplitSeq(input, ",")

	for r := range ranges {
		vals := strings.Split(r, "-")
		min_val, _ := strconv.Atoi(vals[0])
		max_val, _ := strconv.Atoi(vals[1])

		for i := min_val; i <= max_val; i++ {
			id := strconv.Itoa(i)

			if len(id) == 1 {
				continue
			}

			num_digits := len(id)
			divisors := []int{}

			for j := 2; j < num_digits; j++ {
				division := float32(num_digits) / float32(j)
				if division == float32(int32(division)) {
					divisors = append(divisors, j)
				}
			}

			divisors = append(divisors, num_digits)

			for _, divisor := range divisors {
				substrings := split_string(id, divisor)
				if len(substrings) == 1 {
					fmt.Printf("Invalid ID: %d\n", i)
					result_set[i] = struct{}{}
					break
				}
			}
		}
	}

	sum := 0
	for unique_id := range result_set {
		sum += unique_id
	}

	fmt.Printf("Result %d", sum)
}

func main() {
	// part1()
	part2()
}
