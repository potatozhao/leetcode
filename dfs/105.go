/*
根据一棵树的前序遍历与中序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

// 题解：https://www.cnblogs.com/deltadeblog/p/9296469.html 参考了题解。之后要复习

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func buildTree(preorder []int, inorder []int) *TreeNode {
    start := 0
    return createNode(preorder, inorder, &start)
}

func createNode(preorder, inorder []int, pos *int) *TreeNode{
    retNode := new(TreeNode)
    if *pos >= len(preorder){
        return nil
    }
    retNode.Val = preorder[*pos]
    *pos++
    tmpPos := find(inorder, retNode.Val)
    if tmpPos <= 0{
        retNode.Left = nil
    }else{
        retNode.Left = createNode(preorder, inorder[:tmpPos], pos)
    }
    if tmpPos >= len(inorder) - 1{
        retNode.Right = nil
    }else{
        retNode.Right = createNode(preorder, inorder[tmpPos+1:], pos)
    }
    return retNode
}

func find(srcList []int, dst int) int{
    for i := range srcList{
        if srcList[i] == dst{
            return i
        }
    }
    return -1
}