/*
这里有一幅服务器分布图，服务器的位置标识在 m * n 的整数矩阵网格 grid 中，1 表示单元格上有服务器，0 表示没有。

如果两台服务器位于同一行或者同一列，我们就认为它们之间可以进行通信。

请你统计并返回能够与至少一台其他服务器进行通信的服务器的数量。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/count-servers-that-communicate
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// \

package main

// 题解 首先，先统计每行每列的数量，然后遍历，如果这行，行列的数量都是1，那么这个肯定是无联通的。pass

func countServers(grid [][]int) int {
	if len(grid) <= 0 {
		return 0
	}
	if len(grid[0]) <= 0 {
		return 0
	}
	maxNum := 0
	mList := make([]int, len(grid))
	nList := make([]int, len(grid[0]))
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				mList[i]++
				nList[j]++
				maxNum++
			}
		}
	}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				if mList[i] <= 1 && nList[j] <= 1 {
					maxNum--
				}
			}
		}
	}
	return maxNum
}
