/*
145. 二叉树的后序遍历

给定一个二叉树，返回它的 后序 遍历。

示例:

输入: [1,null,2,3]  
   1
    \
     2
    /
   3 

输出: [3,2,1]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-postorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


// 解法： 染色法。跟中序遍历差不多的方法。
type ColorNode struct {
    Node *TreeNode
    Color int
}
func postorderTraversal(root *TreeNode) []int {
    ret := make([]int, 0)
    if root == nil{
        return ret
    }
    stack := make([]ColorNode, 0)
    stack = append(stack, ColorNode{Node:root, Color:0})
    for len(stack) >0{
        top := stack[len(stack) - 1]
        if top.Node.Left != nil && top.Color <1 {
            stack[len(stack) - 1].Color = 1
            stack = append(stack, ColorNode{Node:top.Node.Left, Color:0})
        }else if top.Node.Right != nil && top.Color < 2{
            stack[len(stack) - 1].Color = 2
            stack = append(stack, ColorNode{Node:top.Node.Right, Color:0})
        }else{
            ret = append(ret, top.Node.Val)
            stack = stack[:len(stack) - 1]
        }
    }
    return ret
}

// hash 方案
func postorderTraversal(root *TreeNode) []int {
    ret := make([]int, 0)
    if root == nil{
        return ret
    }
    usedMap := make(map[*TreeNode]struct{})
    stack := make([]*TreeNode, 0)
    stack = append(stack, root)
    for len(stack) >0{
        top := stack[len(stack) - 1]
        if _,ok := usedMap[top.Left];top.Left != nil && !ok  {
            usedMap[top.Left] = struct{}{}
            stack = append(stack, top.Left)
        }else if _,ok := usedMap[top.Right];top.Right != nil && !ok{
            usedMap[top.Right] = struct{}{}
            stack = append(stack, top.Right)
        }else{
            ret = append(ret, top.Val)
            stack = stack[:len(stack) - 1]
        }
    }
    return ret
}