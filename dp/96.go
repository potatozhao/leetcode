/*
给定一个整数 n，求以 1 ... n 为节点组成的二叉搜索树有多少种？

示例:

输入: 3
输出: 5
解释:
给定 n = 3, 一共有 5 种不同结构的二叉搜索树:

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/unique-binary-search-trees
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

// 其实就是从1-i，每一个数做父节点，然后数量就是左子树的数量*右子树，因为是一个连续的数字，那么123的数量也就等于456的数量。
// 所以状态方程是f(x) := for{num += dpList[j-1] * dpList[i-j-1

func numTrees(n int) int {
	dpList := make([]int, n)
	dpList[0] = 1
	for i := 1; i < len(dpList); i++ {
		num := 0
		// 计算左右子树都有的情况
		for j := 1; j < i; j++ {
			num += dpList[j-1] * dpList[i-j-1]
		}
		// 计算只有左子树或者右子树的特殊情况。
		num += 2 * dpList[i-1]
		dpList[i] = num
	}
	return dpList[n-1]
}
