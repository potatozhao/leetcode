/*
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。



示例：

给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import "sort"

func threeSum(nums []int) [][]int {
	ret := make([][]int, 0)
	if len(nums) <= 2 {
		return ret
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		for i > 0 && i < len(nums)-2 && nums[i] == nums[i-1] {
			i++ // 去重
		}
		x, y := i+1, len(nums)-1
		for x < y {
			if nums[x]+nums[y] < 0-nums[i] {
				x++
			} else if nums[x]+nums[y] > 0-nums[i] {
				y--
			} else {
				ret = append(ret, []int{nums[i], nums[x], nums[y]})
				x++
				y--
				for x < y && nums[x-1] == nums[x] {
					x++ // 去重
				}
				for y > x && y < len(nums)-1 && nums[y+1] == nums[y] {
					y-- // 去重
				}
			}
		}
	}
	return ret
}
