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

func flatten2(root *TreeNode)  {
    if root == nil{
        return
    }
    for root != nil{
        if root.Left != nil{
            tmpLeft := root.Left
            for tmpLeft.Right != nil{
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