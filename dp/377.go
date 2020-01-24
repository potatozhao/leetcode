package main

func combinationSum4(nums []int, target int) int {
	dpList := make([]int, target+1)
	for i, _ := range dpList {
		dpList[i] = 0
	}
	dpList[0] = 1
	for i := 1; i <= target; i++ {
		for _, v := range nums {
			if v <= i {
				dpList[i] += dpList[i-v]
			}
		}
	}
	return dpList[target]
}

func main() {
	return
}
