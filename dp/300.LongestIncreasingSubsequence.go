func lengthOfLIS(nums []int) int {
    maxLen := 0
    dpList := make([]int, len(nums))
    for i := 0; i < len(nums); i++{
        dpList[i] = 1
        for j := 0;j<i;j++{
            if nums[i] > nums[j]{
                if dpList[j] + 1 > dpList[i]{
                    dpList[i] = dpList[j] + 1
                }
            }
        }
        if dpList[i] > maxLen{
            maxLen = dpList[i]
        }
    }
    return maxLen
}
// DP解法