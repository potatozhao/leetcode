/*
给出整数数组 A，将该数组分隔为长度最多为 K 的几个（连续）子数组。分隔完成后，每个子数组的中的值都会变为该子数组中的最大值。

返回给定数组完成分隔后的最大和。



示例：

输入：A = [1,15,7,9,2,5,10], K = 3
输出：84
解释：A 变为 [15,15,15,9,10,10,10]


提示：

1 <= K <= A.length <= 500
0 <= A[i] <= 10^6

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/partition-array-for-maximum-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 求数组最大值一般用dp。艹！又没想出来解法！窝囊至极！
package main

func maxSumAfterPartitioning(A []int, K int) int {
	dpList := make([]int, len(A)+1)
	for i := 1; i <= len(A); i++ {
		for j := i - 1; i-j <= K && j >= 0; j-- {
			m := 0
			for _, v := range A[j:i] {
				if v > m {
					m = v
				}
			}
			dpList[i] = max(dpList[i], dpList[j]+m*(i-j))
		}
	}
	return dpList[len(A)]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
