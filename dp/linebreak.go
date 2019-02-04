package main

import (
	"fmt"
	"math"
)

func main() {
	input := []string{
		"aaa",
		"bb",
		"cc",
		"ddddd",
	}

	lineSize := 6

	lc := make([][]int, len(input) + 1)
	for rowIdx, _ := range lc {
		lc[rowIdx] = make([]int, len(input) + 1)
	}

	for rowIdx := 0; rowIdx <= len(input); rowIdx++ {
		for colIdx := rowIdx + 1; colIdx <= len(input); colIdx++ {
			subsequence := input[rowIdx:colIdx]
			seqLen := 0
			for _, str := range subsequence {
				seqLen += len(str) + 1
			}
			if seqLen > 0 { // get rid of the extra length
				seqLen--
			}
			if seqLen > lineSize { // unfitable
				lc[rowIdx][colIdx] = math.MaxInt32
			} else {
				diff := lineSize - seqLen
				lc[rowIdx][colIdx] = diff * diff * diff
			}
		}
	}

	costs := []int{0}

	for j := 1; j <=len(input); j++ {
		min := math.MaxInt32
		for i := 1; i <= j; i++ {
			fmt.Println(j, i, costs[i - 1] + lc[i - 1][j])
			if lc[i][j] == math.MaxInt32 {
				continue
			}
			if c := costs[i - 1] + lc[i - 1][j]; c < min {
				min = c
			}
		}
		costs = append(costs, min)
	}

	fmt.Println(input)
	fmt.Println()
	for _, row := range lc {
		fmt.Println(row)
	}
	fmt.Println()
	fmt.Println(costs)
}
