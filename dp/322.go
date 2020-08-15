package main

/*
322. 零钱兑换
难度
中等

748





给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。

 

示例 1:

输入: coins = [1, 2, 5], amount = 11
输出: 3 
解释: 11 = 5 + 5 + 1
示例 2:

输入: coins = [2], amount = 3
输出: -1
*/


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
