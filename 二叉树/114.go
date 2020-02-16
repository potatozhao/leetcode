/*
给定一个二叉树，原地将它展开为链表。

例如，给定二叉树

    1
   / \
  2   5
 / \   \
3   4   6
将其展开为：

1
 \
  2
   \
    3
     \
      4
       \
        5
         \
          6

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list
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

// dfs解法
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flattenNode(root)
	return
}

func flattenNode(node *TreeNode) *TreeNode {
	if node.Left == nil && node.Right == nil {
		return node
	}
	if node.Left == nil {
		return flattenNode(node.Right)
	} else {
		tmpNode := flattenNode(node.Left)
		tmpNode.Right = node.Right
		node.Right = node.Left
		node.Left = nil
		if tmpNode.Right == nil {
			return tmpNode
		}
		return flattenNode(tmpNode.Right)
	}
}

// bfs解法
func flatten2(root *TreeNode) {
	if root == nil {
		return
	}
	for root != nil {
		if root.Left != nil {
			tmpLeft := root.Left
			for tmpLeft.Right != nil {
				tmpLeft = tmpLeft.Right
			}
			tmpLeft.Right = root.Right
			root.Right = root.Left
			root.Left = nil
		}
		root = root.Right
	}
	return
}
