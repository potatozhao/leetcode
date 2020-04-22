/*
给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回锯齿形层次遍历如下：

[
  [3],
  [20,9],
  [15,7]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal
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
 func zigzagLevelOrder(root *TreeNode) [][]int {
    ret := make([][]int,0)
    if root == nil{
        return ret
    }
    stack := make([]*TreeNode,0)
    stack = append(stack, root)
    deep := 0
    for len(stack) > 0{
        subRet := make([]int, len(stack))
        tmpStack := make([]*TreeNode, 0)
        for i := range stack{
            if deep % 2 == 1{
                subRet[i] = stack[len(stack) - 1 - i].Val
            }else{
                subRet[i] = stack[i].Val
            }
            if stack[i].Left != nil{
                tmpStack = append(tmpStack, stack[i].Left)
            }
            if stack[i].Right != nil{
                tmpStack = append(tmpStack, stack[i].Right)
            }
        }
        deep++
        stack = tmpStack
        ret = append(ret, subRet)
    }
    return ret
}