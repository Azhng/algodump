package main

import (
	"fmt"
)

func main() {
	//dict := map[string]bool{
	//	"pea":    true,
	//	"nut":    true,
	//	"butter": true,
	//}

	//input := "peanutbutter"

	dict := map[string]bool{
		"I":    true,
		"like": true,
		"to":   true,
		"play": true,
	}

	input := "Iliketoplay"

	dp := make([][]int, len(input))
	for rowIdx, _ := range dp {
		dp[rowIdx] = make([]int, len(input))
	}

	for colIdx := 0; colIdx < len(dp[0]); colIdx++ {
		for rowIdx := 0; rowIdx < len(dp)-colIdx; rowIdx++ {
			beginIdx := rowIdx
			endIdx := rowIdx + colIdx
			if _, ok := dict[input[beginIdx:endIdx+1]]; ok {
				dp[beginIdx][endIdx] = beginIdx
			} else {
				set := false
				for k := beginIdx; k < endIdx; k++ {
					fmt.Println(input[beginIdx:k+1], input[k+1:endIdx+1])
					if (endIdx-beginIdx) > 1 && dp[beginIdx][k] != -1 && dp[k+1][endIdx] != -1 {
						dp[beginIdx][endIdx] = k
						set = true
						break
					}
				}
				if !set {
					dp[beginIdx][endIdx] = -1
				}
			}
		}
	}

	fmt.Println(dict)
	for _, row := range dp {
		for _, col := range row {
			fmt.Printf("%2d ", col)
		}
		fmt.Println()
	}

	// just check the top right element in the dp table
}
