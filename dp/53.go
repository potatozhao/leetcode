/*
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:

输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
进阶:

如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-subarray
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/


package main


// f(x) = max(f(x-1) + a(x), a(x)) f(x)的意思是以x结尾的最长子序列。
func maxSubArray(nums []int) int {
    if len(nums) == 0{
        return 0
    }
    intMin := -2147483647
    maxNum := int64(intMin)
    sum := int64(intMin)
    for  i := range nums{
        sum = max(sum + int64(nums[i]), int64(nums[i]))
        maxNum = max(sum, maxNum)
    }
    return int(maxNum)
}

func max(a,b int64) int64 {
    if a > b{
        return a
    }
    return b
}