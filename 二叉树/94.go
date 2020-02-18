/*
给定一个二叉树，返回它的中序 遍历。

示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,3,2]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-inorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//
func inorderTraversal(root *TreeNode) []int {
	retList := make([]int, 0)
	dfs(root, &retList)
	return retList
}

// 递归解法。十分简单。
func dfs(node *TreeNode, list *[]int) {
	if node == nil {
		return
	}
	dfs(node.Left, list)
	*list = append(*list, node.Val)
	dfs(node.Right, list)
	return
}
