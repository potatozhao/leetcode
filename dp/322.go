package main

import (
	"fmt"
)

func coinChange(coins []int, amount int) int {
	dpList := make([]int, amount+1)
	for i, _ := range dpList {
		dpList[i] = 999999999
	}
	dpList[0] = 0
	for i := 0; i < len(coins); i++ {
		for j := 0; j <= amount; j++ {
			if coins[i] <= j {
				dpList[j] = min(dpList[j], dpList[j-coins[i]]+1)
			}
		}
	}
	if dpList[amount] > amount {
		return -1
	} else {
		return dpList[amount]
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	a := []int{1, 2, 5}
	fmt.Println(coinChange(a, 11))
	return
}
