/*
给定一个可能包含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。

示例:

输入: [1,2,2]
输出:
[
  [2],
  [1],
  [1,2,2],
  [2,2],
  [1,2],
  []
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/subsets-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 此题目很惨，没做出来。艹了狗！
// 按照题解，首先是要排序，
// 排序后，根据
package main

import "sort"

type IntSlice []int

func (s IntSlice) Len() int { return len(s) }

func (s IntSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }

func subsetsWithDup(nums []int) [][]int {
	sort.Sort(IntSlice(nums))
	ret := make([][]int, 0)
	sub(nums, 0, []int{}, &ret)
	return ret
}

func sub(nums []int, pos int, array []int, ret *[][]int) {
	*ret = append(*ret, array)
	for i := pos; i < len(nums); i++ {
		if i > pos && nums[i] == nums[i-1] {
			continue
		}
		tmp := make([]int, len(array))
		for i := range array {
			tmp[i] = (array)[i]
		}
		tmp = append(tmp, nums[i])
		sub(nums, i+1, tmp, ret)
	}
}
