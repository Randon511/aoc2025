package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

type Coor struct {
	x int
	y int
}

func part1() {
	file, err := os.Open("input2.txt") // For read access.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	dirs := []Coor{
		{x: 1, y: 1},
		{x: 1, y: -1},
		{x: 1, y: 0},
		{x: -1, y: 1},
		{x: -1, y: -1},
		{x: -1, y: 0},
		{x: 0, y: 1},
		{x: 0, y: -1},
	}

	scanner := bufio.NewScanner(file)
	count := 0

	input_w := 0
	input_h := 0

	// Build input matrix and create border to make searching easier
	var matrix [][]byte
	for scanner.Scan() {
		line := scanner.Text()

		if len(matrix) == 0 {
			input_w = len(line)
			first_row := bytes.Repeat([]byte{'.'}, input_w+2)
			matrix = append(matrix, first_row)
		}
		row := make([]byte, 0, input_w+2)
		row = append(row, '.')
		row = append(row, []byte(line)...)
		row = append(row, '.')

		matrix = append(matrix, row)
		input_h++
	}

	last_row := bytes.Repeat([]byte{'.'}, input_w+2)
	matrix = append(matrix, last_row)

	for i := range input_w + 2 {
		matrix[0][i] = '.'
		matrix[input_h+1][i] = '.'
	}

	for i := 1; i < input_h+1; i++ {
		for j := 1; j < input_w+1; j++ {
			if matrix[i][j] == '@' {
				num_adj := 0
				for k := range dirs {
					if matrix[i+dirs[k].y][j+dirs[k].x] == '@' {
						num_adj++
					}
				}

				if num_adj < 4 {
					count++
				}
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

	dirs := []Coor{
		{x: 1, y: 1},
		{x: 1, y: -1},
		{x: 1, y: 0},
		{x: -1, y: 1},
		{x: -1, y: -1},
		{x: -1, y: 0},
		{x: 0, y: 1},
		{x: 0, y: -1},
	}

	scanner := bufio.NewScanner(file)
	count := 0

	input_w := 0
	input_h := 0

	// Build input matrix and create border to make searching easier
	var matrix [][]byte
	for scanner.Scan() {
		line := scanner.Text()

		if len(matrix) == 0 {
			input_w = len(line)
			first_row := bytes.Repeat([]byte{'.'}, input_w+2)
			matrix = append(matrix, first_row)
		}
		row := make([]byte, 0, input_w+2)
		row = append(row, '.')
		row = append(row, []byte(line)...)
		row = append(row, '.')

		matrix = append(matrix, row)
		input_h++
	}

	last_row := bytes.Repeat([]byte{'.'}, input_w+2)
	matrix = append(matrix, last_row)

	for i := range input_w + 2 {
		matrix[0][i] = '.'
		matrix[input_h+1][i] = '.'
	}

	for {
		num_removed := 0
		coor_removed := []Coor{}
		for i := 1; i < input_h+1; i++ {
			for j := 1; j < input_w+1; j++ {
				if matrix[i][j] == '@' {
					num_adj := 0
					for k := range dirs {
						if matrix[i+dirs[k].y][j+dirs[k].x] == '@' {
							num_adj++
						}
					}
					if num_adj < 4 {
						num_removed++
						coor_removed = append(coor_removed, Coor{x: j, y: i})
					}
				}
			}
		}

		if num_removed == 0 {
			break
		}

		count += num_removed
		for _, coor := range coor_removed {
			matrix[coor.y][coor.x] = '.'
		}
	}
	fmt.Printf("Result %d\n", count)
}

func main() {
	// part1()
	part2()
}
