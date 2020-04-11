/*
41. 缺失的第一个正数

给你一个未排序的整数数组，请你找出其中没有出现的最小的正整数。

 

示例 1:

输入: [1,2,0]
输出: 3
示例 2:

输入: [3,4,-1,1]
输出: 2
示例 3:

输入: [7,8,9,11,12]
输出: 1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/first-missing-positive
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/



// 利用下标index的方法，注意重复。
package main

func firstMissingPositive(nums []int) int {
    for i := 0;i < len(nums);{
        for nums[i] < len(nums) && nums[i] != i + 1 && nums[i] > 0{
            key := nums[i] - 1
            if nums[i] == nums[key]{
                break
            }
            nums[i], nums[key] = nums[key], nums[i] 
        }
        i++
    }
    for i := range nums{
        if nums[i] != i+1{
            return i + 1
        }
    }
    return len(nums)+1
}