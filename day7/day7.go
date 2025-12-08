package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func part1() {
	data, err := os.ReadFile("input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")

	count := 0
	t_indexes := make([]bool, len(lines[0]))
	for i := range t_indexes {
		t_indexes[i] = false
	}

	for idx, ch := range lines[0] {
		if ch == 'S' {
			t_indexes[idx] = true
			break
		}
	}

	for i := 2; i < len(lines)-1; i++ {
		new_t_indexes := make([]bool, len(lines[0]))
		for j := range new_t_indexes {
			new_t_indexes[j] = false
		}

		for j := 0; j < len(t_indexes); j++ {
			if t_indexes[j] == true {
				if lines[i][j] == '^' {
					new_t_indexes[j+1] = true
					new_t_indexes[j-1] = true
					count++
				} else {
					new_t_indexes[j] = true
				}
			}
		}

		t_indexes = new_t_indexes
	}
	fmt.Printf("Result: %d", count)
}

func part2() {
	data, err := os.ReadFile("input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")
	beams := make([]int, len(lines[0]))

	for idx, ch := range lines[0] {
		if ch == 'S' {
			beams[idx]++
			break
		}
	}

	for i := 2; i < len(lines)-1; i += 2 {
		old_beams := beams
		new_beams := make([]int, len(lines[0]))

		for j := range len(lines[0]) {
			if lines[i][j] == '^' {
				new_beams[j+1] += old_beams[j]
				new_beams[j-1] += old_beams[j]
			} else {
				new_beams[j] += old_beams[j]
			}
		}

		beams = new_beams
	}

	result := 0
	for _, val := range beams {
		result += val
	}
	fmt.Printf("Result: %d", result)
}

func main() {
	// part1()
	part2()
}
