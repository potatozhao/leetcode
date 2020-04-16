/*
673. 最长递增子序列的个数

给定一个未排序的整数数组，找到最长递增子序列的个数。

示例 1:

输入: [1,3,5,4,7]
输出: 2
解释: 有两个最长递增子序列，分别是 [1, 3, 4, 7] 和[1, 3, 5, 7]。
示例 2:

输入: [2,2,2,2,2]
输出: 5
解释: 最长递增子序列的长度是1，并且存在5个子序列的长度为1，因此输出5。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-longest-increasing-subsequence
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

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