/**
给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。

示例:

输入: nums = [1,2,3]
输出:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/subsets
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

//简单的dfs，内存需要优化。

func subsets(nums []int) [][]int {
	return dfs(nums, 0)
}

func dfs(nums []int, pos int) [][]int {
	if pos >= len(nums) {
		return [][]int{{}}
	}
	sonRet := dfs(nums, pos+1)
	return copyNums(sonRet, nums[pos])
}

func copyNums(in [][]int, num int) [][]int {
	ret := make([][]int, 0, len(in)*2)
	for _, v := range in {
		tmpList1 := make([]int, len(v))
		tmpList2 := make([]int, len(v)+1)
		tmpList2[0] = num
		for i, _ := range v {
			tmpList1[i] = v[i]
			tmpList2[i+1] = v[i]
		}
		ret = append(ret, tmpList1)
		ret = append(ret, tmpList2)
	}
	return ret
}
