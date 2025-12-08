package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Node struct {
	x   int
	y   int
	z   int
	cid int
}

type NodePair struct {
	n1   *Node
	n2   *Node
	dist float64
}

type kv struct {
	k int
	v int
}

func calc_dist(n1 Node, n2 Node) float64 {
	dx := float64(n2.x - n1.x)
	dy := float64(n2.y - n1.y)
	dz := float64(n2.z - n1.z)
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func part1() {
	data, err := os.ReadFile("input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")
	nodes := []Node{}
	pairs := []NodePair{}
	cur_num_circuits := 0

	for _, coords := range lines {
		if coords == "" {
			break
		}

		coords = strings.TrimSpace(coords)
		parts := strings.Split(coords, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])

		nodes = append(nodes, Node{x: x, y: y, z: z, cid: 0})
	}

	for i := 0; i < len(nodes); i++ {
		for j := i + 1; j < len(nodes); j++ {
			dist := calc_dist(nodes[i], nodes[j])
			pairs = append(pairs, NodePair{n1: &nodes[i], n2: &nodes[j], dist: dist})
		}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].dist < pairs[j].dist
	})

	for i := 0; i < 1000; i++ {
		pair := pairs[i]
		if pair.n1.cid == 0 && pair.n2.cid == 0 {
			cur_num_circuits++
			pair.n1.cid = cur_num_circuits
			pair.n2.cid = cur_num_circuits
		} else if pair.n1.cid != 0 && pair.n2.cid == 0 {
			pair.n2.cid = pair.n1.cid
		} else if pair.n1.cid == 0 && pair.n2.cid != 0 {
			pair.n1.cid = pair.n2.cid
		} else {
			old_id := pair.n2.cid
			new_id := pair.n1.cid
			for idx, _ := range nodes {
				if nodes[idx].cid == old_id {
					nodes[idx].cid = new_id
				}
			}
		}
	}

	c_size := make(map[int]int)
	for _, n := range nodes {
		c_size[n.cid]++
	}

	vals := make([]kv, 0, len(c_size))
	for cid, cnt := range c_size {
		if cid != 0 {
			vals = append(vals, kv{cid, cnt})
		}
	}

	sort.Slice(vals, func(i, j int) bool {
		return vals[i].v > vals[j].v
	})

	result := 1
	for i := 0; i < len(nodes); i++ {
		if i > 2 {
			break
		}

		result = result * vals[i].v
	}

	fmt.Printf("Res :%d", result)
}

func part2() {
	data, err := os.ReadFile("input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")
	nodes := []Node{}
	pairs := []NodePair{}
	cur_num_circuits := 0

	for _, coords := range lines {
		if coords == "" {
			break
		}

		coords = strings.TrimSpace(coords)
		parts := strings.Split(coords, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])

		nodes = append(nodes, Node{x: x, y: y, z: z, cid: 0})
	}

	for i := 0; i < len(nodes); i++ {
		for j := i + 1; j < len(nodes); j++ {
			dist := calc_dist(nodes[i], nodes[j])
			pairs = append(pairs, NodePair{n1: &nodes[i], n2: &nodes[j], dist: dist})
		}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].dist < pairs[j].dist
	})

	for i := 0; i < len(pairs); i++ {
		pair := pairs[i]
		if pair.n1.cid == 0 && pair.n2.cid == 0 {
			cur_num_circuits++
			pair.n1.cid = cur_num_circuits
			pair.n2.cid = cur_num_circuits
		} else if pair.n1.cid != 0 && pair.n2.cid == 0 {
			pair.n2.cid = pair.n1.cid
		} else if pair.n1.cid == 0 && pair.n2.cid != 0 {
			pair.n1.cid = pair.n2.cid
		} else {
			old_id := pair.n2.cid
			new_id := pair.n1.cid
			for idx, _ := range nodes {
				if nodes[idx].cid == old_id {
					nodes[idx].cid = new_id
				}
			}
		}

		done := true
		first_id := nodes[0].cid
		for _, node := range nodes {
			if node.cid != first_id {
				done = false
				break
			}
		}

		if done {
			fmt.Printf("Res :%d\n", pair.n1.x*pair.n2.x)
			break
		}
	}
}

func main() {
	// part1()
	part2()
}
