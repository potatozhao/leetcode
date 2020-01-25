package main

import "fmt"

func change(amount int, coins []int) int {
	dpList := make([]int, amount+1)
	dpList[0] = 1
	for _, v := range coins {
		for i := 1; i <= amount; i++ {
			if v <= i {
				dpList[i] += dpList[i-v]
			}
		}
		fmt.Println(dpList)
	}
	return dpList[amount]
}

func main() {
	a := []int{1, 2, 5}
	fmt.Println(change(5, a))
}
