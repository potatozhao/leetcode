/*
给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。

例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [9,20],
  [15,7]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal
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
 func levelOrder(root *TreeNode) [][]int {
    ret := make([][]int, 0)
    if root == nil{
        return ret
    }
    queue := make([]*TreeNode, 0)
    queue = append(queue, root)
    for len(queue) != 0{
        tmpList := make([]int,0, len(queue))
        tmpQueue  := make([]*TreeNode, 0)
        for i := range queue{
            tmpNode := queue[i]
            if tmpNode.Left != nil{
                tmpQueue = append(tmpQueue, tmpNode.Left)
            }
            if tmpNode.Right != nil{
                tmpQueue = append(tmpQueue, tmpNode.Right)
            }
            tmpList = append(tmpList, tmpNode.Val)
        }
        ret = append(ret, tmpList)
        queue = tmpQueue
    }
    return ret
}