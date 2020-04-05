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

// 染色法
// 迭代方案，color记录左子节点是否已经迭代过。
func inorderTraversal2(root *TreeNode) []int {
	retList := make([]int, 0)
	if root == nil {
		return retList
	}

	type ColorNode struct {
		Color int
		Node  *TreeNode
	}

	stackList := make([]ColorNode, 0)
	stackList = append(stackList, ColorNode{Color: 0, Node: root})
	for len(stackList) > 0 {
		topNode := stackList[len(stackList)-1]
		stackList = stackList[:len(stackList)-1]
		if topNode.Node.Left != nil && topNode.Color < 1 {
			topNode.Color = 1
			stackList = append(stackList, topNode)
			stackList = append(stackList, ColorNode{Color: 0, Node: topNode.Node.Left})
			continue
		}
		retList = append(retList, topNode.Node.Val)
		if topNode.Node.Right != nil {
			stackList = append(stackList, ColorNode{Color: 0, Node: topNode.Node.Right})
		}
	}
	return retList
}
