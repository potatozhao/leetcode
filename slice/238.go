/*
给你一个长度为 n 的整数数组 nums，其中 n > 1，返回输出数组 output ，其中 output[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积。

 

示例:

输入: [1,2,3,4]
输出: [24,12,8,6]
 

提示：题目数据保证数组之中任意元素的全部前缀元素和后缀（甚至是整个数组）的乘积都在 32 位整数范围内。

说明: 请不要使用除法，且在 O(n) 时间复杂度内完成此题。

进阶：
你可以在常数空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组不被视为额外空间。）

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/product-of-array-except-self
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main


// 比较low的方案
func productExceptSelf(nums []int) []int {
    minList := make([]int, len(nums))
    maxList := make([]int, len(nums))
    tmpi := 1
    tmpa :=1
    for i := range nums{
        tmpa *= nums[i]
        tmpi *= nums[len(nums) - i - 1]
        maxList[i] = tmpa
        minList[len(nums) - i - 1] = tmpi
    }
    ret := make([]int, len(nums))
    ret[0] = minList[1]
    ret[len(ret) - 1] = maxList[len(maxList) - 2]
    for i := 1; i < len(ret) - 1; i ++{
        ret[i] = maxList[i-1] * minList[i + 1]
    }
    return ret
}


// 利用返回数组缓存状态，然后反向迭代进行求积
func productExceptSelf(nums []int) []int {
    ret := make([]int, len(nums))
    ret[0] = nums[0]
    for i := 1; i< len(nums); i++{
        ret[i] = nums[i] * ret[i-1]
    }
    tmp := 1
    for i := len(nums) - 1; i >=0; i--{
        if i == 0{
            ret[i] = tmp
        }else{
            ret[i] = ret[i-1] * tmp
        }
        tmp *= nums[i]
    }
    return ret
}