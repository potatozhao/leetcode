/*给两个整数数组 A 和 B ，返回两个数组中公共的、长度最长的子数组的长度。

示例 1:

输入:
A: [1,2,3,2,1]
B: [3,2,1,4,7]
输出: 3
解释:
长度最长的公共子数组是 [3, 2, 1]。
说明:

1 <= len(A), len(B) <= 1000
0 <= A[i], B[i] < 100
*/

package main

func findLength(A []int, B []int) int {
	dpList := make([][]int, len(A)+1)
	for i, _ := range dpList {
		dpList[i] = make([]int, len(B)+1)
	}
	maxNum := 0
	for i := 1; i <= len(A); i++ {
		for j := 1; j <= len(B); j++ {
			if A[i-1] == B[j-1] {
				dpList[i][j] = dpList[i-1][j-1] + 1
			}
			if dpList[i][j] > maxNum {
				maxNum = dpList[i][j]
			}
		}
	}
	return maxNum
}

// 此答案很糟糕 公式为f(i,j) = f(i-1,j-1)
