/*
给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字）。

 

示例 1:

输入: [2,3,-2,4]
输出: 6
解释: 子数组 [2,3] 有最大乘积 6。
示例 2:

输入: [-2,0,-1]
输出: 0
解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-product-subarray
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main


// 超时

func maxProduct(nums []int) int {
    if len(nums) == 1{
        return nums[0]
    }
    max := 0
    for i := 1; i <= len(nums); i++{
        for j := 0; j <= len(nums) - i; j++{
            tmp := 1
            for m := j;  m < j + i; m++{
                tmp *= nums[m]
            }
            if tmp > max{
                max = tmp
            }
        }
    }
    return max
}

// 很糟。。。只能说能用吧。。。
func maxProduct(nums []int) int {
    if len(nums) == 1{
        return nums[0]
    }
    ret := 0
    dpList := make([]int, len(nums))
    for i := 0; i < len(nums);i++{
        dpList[i] = nums[i]
        ret = max(dpList[i], ret)
        for j := i+1; j < len(nums); j++{
            dpList[j] = dpList[j-1] * nums[j]
            ret = max(dpList[j], ret)
        }
    }
    return ret
}


// 动规解法。
func maxProduct(nums []int) int {
    if len(nums) == 1{
        return nums[0]
    }
    ret := 0
    mi := 1
    ma := 1
    for i := range nums{
        if nums[i] >= 0{
            mi = min(mi * nums[i], nums[i])
            ma = max(ma * nums[i], nums[i])
        }else{
            tmpi := mi
            mi = min(ma * nums[i], nums[i])
            ma = max(tmpi * nums[i], nums[i])
        }
        if ma > ret{
            ret = ma
        }
    }
    return ret
}

func max(a,b int) int {
    if a >b {
        return a
    }
    return b
}

func min(a,b int) int{
    if a < b{
        return a
    }
    return b
}