/*
给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。

示例:

输入:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
输出: 7
解释: 因为路径 1→3→1→1→1 的总和最小。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-path-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

// 题解： f(i,j) = min(f(i,j-1),f(i-1,j)) + grid[i][j]
func minPathSum(grid [][]int) int {
	if len(grid) <= 0 {
		return 0
	}
	dpList := make([][]int, len(grid))
	for i, _ := range dpList {
		dpList[i] = make([]int, len(grid[0]))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 && j == 0 {
				dpList[i][j] = grid[i][j]
			} else if i == 0 {
				dpList[i][j] = dpList[i][j-1] + grid[i][j]
			} else if j == 0 {
				dpList[i][j] = dpList[i-1][j] + grid[i][j]
			} else {
				dpList[i][j] = min(dpList[i][j-1], dpList[i-1][j]) + grid[i][j]
			}
		}
	}
	return dpList[len(grid)-1][len(grid[0])-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
