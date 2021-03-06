/*
给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。

注意:
不能使用代码库中的排序函数来解决这道题。

示例:

输入: [2,0,2,1,1,0]
输出: [0,0,1,1,2,2]
进阶：

一个直观的解决方案是使用计数排序的两趟扫描算法。
首先，迭代计算出0、1 和 2 元素的个数，然后按照0、1、2的排序，重写当前数组。
你能想出一个仅使用常数空间的一趟扫描算法吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sort-colors
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 解法，两次互换。快排原理。
func sortColors(nums []int)  {
    if len(nums) <= 1{
        return 
    }
    start := sort(nums, 0)
    if start >= len(nums){
        return
    }
    sort(nums[start:], 1)
}

func sort(nums []int, dst int) int {
    start := 0
    end := len(nums) -1
    for start < end{
        if nums[start] == dst{
            start++
            continue
        }
        if nums[end] != dst{
            end--
            continue
        }
        nums[start], nums[end] = nums[end], nums[start]
        start++
        end--
    }
    if nums[start] == dst{
        return start + 1
    }
    return start
}
// 荷兰国旗问题。解法明早再看吧。
// 三指针方案。
// base case start := -1 end = len(nums) mid = 0
// corner case start < mid < end
// if nums[mid] == dst mid++
// if nums[mid] < dst [mid] <=> [start+1] start++ mid++
// if nums[mid] > dst [mid] <=> [end-1] end --

/// 注意步骤！！！！！

func sortColors(nums []int)  {
    start := -1
    end := len(nums)
    mid := 0
    for start< mid && mid < end{
        if nums[mid] == 1{
            mid++
        }else if nums[mid] < 1{
            nums[mid], nums[start+1] = nums[start + 1], nums[mid]
            start++
            mid++
        }else{
            nums[mid], nums[end - 1] = nums[end - 1], nums[mid]
            end--
        }
    }
}