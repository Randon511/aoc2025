package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Equation struct {
	nums []int
	op   string
}

func part1() {
	file, err := os.Open("input2.txt") // For read access.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := 0
	num_eq := 0
	equations := []Equation{}

	// Find good ids
	for scanner.Scan() {
		line := scanner.Text()

		vals := strings.Split(line, " ")
		if num_eq == 0 {
			num_eq = len(vals)
			for _, i := range vals {
				as_int, err := strconv.Atoi(i)
				if err == nil {
					equations = append(equations, Equation{[]int{as_int}, " "})
				}
			}
			continue
		}

		if vals[0] == "+" || vals[0] == "*" {
			idx := 0
			for _, val := range vals {
				if val == "*" || val == "+" {
					equations[idx].op = val
					idx++
				}
			}
		} else {
			idx := 0
			for _, val := range vals {
				as_int, err := strconv.Atoi(val)
				if err == nil {
					equations[idx].nums = append(equations[idx].nums, as_int)
					idx++
				}
			}
		}
	}

	for _, eq := range equations {
		res := 0
		if eq.op == "+" {
			fmt.Println(eq.nums)
			fmt.Println(eq.op)
			for _, num := range eq.nums {
				res = res + num
			}
		} else {
			fmt.Println(eq.nums)
			fmt.Println(eq.op)
			res = 1
			for _, num := range eq.nums {
				res = res * num
			}
		}
		fmt.Println(res)

		result += res
	}

	fmt.Printf("Result %d", result)
}

func part2() {
	data, err := os.ReadFile("input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")
	rows := len(lines) - 1
	cols := len(lines[0])

	result := 0
	cur_val := 0
	cur_result := 0
	add := false
	for i := 0; i < cols-1; i++ {
		sep := true
		for j := 0; j < rows; j++ {
			if lines[j][i] != ' ' {
				sep = false
				break
			}
		}

		if sep {
			fmt.Printf("%d\n", cur_result)
			result += cur_result
			continue
		}

		// New op
		if lines[rows-1][i] != ' ' {
			cur_val = 0
			if lines[rows-1][i] == '+' {
				fmt.Printf("Add ")
				add = true
				cur_result = 0
			} else {
				fmt.Printf("Mul ")
				add = false
				cur_result = 1
			}
		}

		num_str := ""
		for j := 0; j < rows-1; j++ {
			if lines[j][i] != ' ' {
				num_str += string(lines[j][i])
			}
		}
		cur_val, _ = strconv.Atoi(num_str)
		fmt.Printf("%d ", cur_val)

		if add {
			cur_result += cur_val
		} else {
			cur_result = cur_result * cur_val
		}
	}

	fmt.Printf("%d\n", cur_result)
	result += cur_result
	fmt.Printf("Result %d", result)
}

func main() {
	// part1()
	part2()
}
