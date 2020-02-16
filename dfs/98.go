/*
给定一个二叉树，判断其是否是一个有效的二叉搜索树。

假设一个二叉搜索树具有如下特征：

节点的左子树只包含小于当前节点的数。
节点的右子树只包含大于当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
示例 1:

输入:
    2
   / \
  1   3
输出: true
示例 2:

输入:
    5
   / \
  1   4
     / \
    3   6
输出: false
解释: 输入为: [5,1,4,null,null,3,6]。
     根节点的值为 5 ，但是其右子节点值为 4 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/validate-binary-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

// 此题提交五次后通过。。。。细节啊！！！！细节啊！！！！！！理解题意啊！！！！！

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	_, _, flag := dfs(root)
	return flag
}

func dfs(node *TreeNode) (int, int, bool) {
	var min, max int
	if node.Left != nil {
		minL, maxL, flag := dfs(node.Left)
		if !flag || node.Val <= maxL {
			return minL, maxL, false
		} else {
			min = minL
		}
	} else {
		min = node.Val
	}

	if node.Right != nil {
		minR, maxR, flag := dfs(node.Right)
		if !flag || node.Val >= minR {
			return minR, maxR, false
		} else {
			max = maxR
		}
	} else {
		max = node.Val
	}

	return min, max, true
}

// 第二种dfs解法，代码简单易懂
func isValidBST2(root *TreeNode) bool {
	intMax := int(^uint(0) >> 1)
	intMin := ^intMax
	return dfs(root, intMin, intMax)
}

func dfs2(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}
	if node.Val >= max || node.Val <= min {
		return false
	}
	return dfs2(node.Left, min, node.Val) && dfs2(node.Right, node.Val, max)
}
