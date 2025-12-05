package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type FreshRange struct {
	min int64
	max int64
}

func part1() {
	file, err := os.Open("input2.txt") // For read access.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	fresh_ranges := []FreshRange{}

	// Find good ids
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		nums := strings.Split(line, "-")
		min_val, _ := strconv.ParseInt(nums[0], 10, 64)
		max_val, _ := strconv.ParseInt(nums[1], 10, 64)

		fresh_ranges = append(fresh_ranges, FreshRange{min: min_val, max: max_val})
	}

	// Check good ingredients
	for scanner.Scan() {
		line := scanner.Text()

		val, _ := strconv.ParseInt(line, 10, 64)

		for _, fresh_range := range fresh_ranges {
			if (val <= fresh_range.max) && (val >= fresh_range.min) {
				count++
				break
			}
		}
	}

	fmt.Printf("Result %d", count)
}

func part2() {
	file, err := os.Open("input2.txt") // For read access.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := int64(0)
	fresh_ranges := []FreshRange{}

	// Find good ids
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		nums := strings.Split(line, "-")
		min_val, _ := strconv.ParseInt(nums[0], 10, 64)
		max_val, _ := strconv.ParseInt(nums[1], 10, 64)

		fresh_ranges = append(fresh_ranges, FreshRange{min: min_val, max: max_val})
	}

	sort.Slice(fresh_ranges, func(i, j int) bool {
		return fresh_ranges[i].min < fresh_ranges[j].min
	})

	merged_ranges := []FreshRange{}
	current_min := fresh_ranges[0].min
	current_max := fresh_ranges[0].max

	for _, fresh_range := range fresh_ranges {
		if fresh_range.min > current_max {
			merged_ranges = append(merged_ranges, FreshRange{current_min, current_max})
			current_min = fresh_range.min
			current_max = fresh_range.max
		} else if fresh_range.max > current_max {
			current_max = fresh_range.max
		}
	}

	merged_ranges = append(merged_ranges, FreshRange{current_min, current_max})

	for _, merged_range := range merged_ranges {
		count += merged_range.max - merged_range.min + 1
	}

	fmt.Printf("Result %d", count)
}

func main() {
	// part1()
	part2()
}
