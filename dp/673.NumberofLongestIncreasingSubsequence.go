func findNumberOfLIS(nums []int) int {
    maxNum := 0
    countNum := 0
    countList := make([]int, len(nums))
    maxList := make([]int, len(nums))
    for i := 0;i < len(nums); i++{
        maxList[i] = 1
        countList[i] = 1
        for j := 0;j<=i;j++{
            if nums[j] < nums[i]{
                if maxList[j] + 1 > maxList[i]{
                    maxList[i] = maxList[j] + 1
                    countList[i] = countList[j]
                } else if maxList[j] + 1 == maxList[i]{
                    countList[i] += countList[j]
                }
            }
        }
    }
    for i,_ := range maxList{
        if maxList[i] > maxNum{
            maxNum = maxList[i]
            countNum = countList[i]
        }else if maxList[i] == maxNum{
            countNum += countList[i]
        } 
    }
    return countNum
}