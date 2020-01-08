/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func flatten(root *TreeNode)  {
    if root == nil{
        return
    }
    flattenNode(root)
    return
}

func flattenNode(node *TreeNode) *TreeNode{
    if node.Left == nil && node.Right == nil{
        return node
    }
    if node.Left == nil{
        return flattenNode(node.Right)
    }else{
        tmpNode := flattenNode(node.Left)
        tmpNode.Right = node.Right
        node.Right = node.Left
        node.Left = nil
        if tmpNode.Right == nil{
            return tmpNode
        }
        return flattenNode(tmpNode.Right)
    }
}