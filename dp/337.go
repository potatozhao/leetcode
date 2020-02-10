/*
在上次打劫完一条街道之后和一圈房屋后，小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为“根”。 除了“根”之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。 如果两个直接相连的房子在同一天晚上被打劫，房屋将自动报警。

计算在不触动警报的情况下，小偷一晚能够盗取的最高金额。

示例 1:

输入: [3,2,3,null,3,null,1]

     3
    / \
   2   3
    \   \
     3   1

输出: 7
解释: 小偷一晚能够盗取的最高金额 = 3 + 3 + 1 = 7.
示例 2:

输入: [3,4,5,1,3,null,1]

     3
    / \
   4   5
  / \   \
 1   3   1

输出: 9
解释: 小偷一晚能够盗取的最高金额 = 4 + 5 = 9.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/house-robber-iii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

//需要重新复习下。没有好好做。思路没理清楚。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
	cache := make(map[*TreeNode]int)
	return subRob(root, cache)

}

func subRob(node *TreeNode, cache map[*TreeNode]int) int {
	if node == nil {
		return 0
	}
	if ret, ok := cache[node]; ok {
		return ret
	}
	var leftDo, rigthDo int
	if node.Left == nil {
		leftDo = 0
	} else {
		leftDo = subRob(node.Left.Left, cache) + subRob(node.Left.Right, cache)
	}
	if node.Right == nil {
		rigthDo = 0
	} else {
		rigthDo = subRob(node.Right.Left, cache) + subRob(node.Right.Right, cache)
	}

	nDo := subRob(node.Left, cache) + subRob(node.Right, cache)
	ret := max(leftDo+rigthDo+node.Val, nDo)
	cache[node] = ret
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//此处解答，及其巧妙！妙哉！
func rob2(root *TreeNode) int {
	a, b := subRob(root)
	return max(a, b)

}

func subRob2(node *TreeNode) (int, int) {
	if node == nil {
		return 0, 0
	}
	leftA, leftB := subRob2(node.Left)
	rightA, rightB := subRob2(node.Right)

	a := max(leftA, leftB) + max(rightA, rightB)
	b := leftA + rightA + node.Val
	return a, b
}
