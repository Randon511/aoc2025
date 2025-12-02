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
	current_val := 50
	count := 0

	for scanner.Scan() {
		line := scanner.Text()
		dir := line[:1]
		val_str := line[1:]
		val, err := strconv.Atoi(val_str)
		if err != nil {
			fmt.Printf("Atoi failed: %s", val_str)
			return
		}

		// Limit to as big as a single rotation
		converted_val := val % 100

		// Let a negative number represent rotating left
		if dir == "L" {
			converted_val = -1 * converted_val
		}

		sum := current_val + converted_val

		if sum > 99 {
			current_val = sum - 100
		} else if sum < 0 {
			current_val = sum + 100
		} else {
			current_val = sum
		}

		if current_val == 0 {
			count++
		}

		// fmt.Printf("Rotation: %s%d, new val: %d\n", dir, val, current_val)
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
	current_val := 50
	init_val := current_val
	count := 0

	for scanner.Scan() {
		line := scanner.Text()
		dir := line[:1]
		val_str := line[1:]
		val, err := strconv.Atoi(val_str)
		if err != nil {
			fmt.Printf("Atoi failed: %s", val_str)
			return
		}

		num_multiples := (int)(val / 100)
		count += num_multiples

		// Limit to as big as a single rotation
		converted_val := val % 100

		// Let a negative number represent rotating left
		if dir == "L" {
			converted_val = -1 * converted_val
		}

		init_val = current_val
		sum := current_val + converted_val

		fmt.Printf("Init: %d ", current_val)

		if sum > 99 {
			current_val = sum - 100
		} else if sum < 0 {
			current_val = sum + 100
		} else {
			current_val = sum
		}

		// Just landed on zero, or just passed through zero
		if (current_val == 0) || ((init_val != 0) && ((sum > 99) || (sum < 0))) {
			count++
		}

		fmt.Printf("\n")
	}

	fmt.Printf("Result %d", count)
}

func main() {
	// part1()
	part2()
}
