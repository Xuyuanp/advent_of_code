package main

import (
	"fmt"
	"log"

	"github.com/Xuyuanp/advent_of_code/pkg/utils"
)

const inputPath = "2020/01/puzzle1/input.txt"

func twoSum(vals []int, magic int) (int, int) {
	m := make(map[int]bool)
	for _, v := range vals {
		if m[v] {
			return v, magic - v
		}
		m[magic-v] = true
	}
	return -1, -1
}

// I known this is not the best solution but the most straightforward
func threeSum(vals []int, magic int) (int, int, int) {
	for i := 0; i < len(vals)-2; i++ {
		x, y := twoSum(vals[i+1:], magic-vals[i])
		if x != -1 {
			return vals[i], x, y
		}
	}
	return -1, -1, -1
}

func main() {
	lines, err := utils.InputInts(inputPath)
	if err != nil {
		log.Fatalf("get input failed: %v", err)
	}

	x, y, z := threeSum(lines, 2020)
	if x != -1 {
		fmt.Println(x, y, z)
		fmt.Println(x * y * z)
	}
}
