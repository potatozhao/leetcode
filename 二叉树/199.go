/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func rightSideView(root *TreeNode) []int {
    ret := make([]int,0)
    if root == nil{
        return ret
    }
    tmpList := make([]*TreeNode,0, 1)
    tmpList = append(tmpList, root)
    for len(tmpList) != 0{
        tmp := make([]*TreeNode,0)
        for i,_ := range tmpList{
            if tmpList[i].Left != nil{
                tmp = append(tmp, tmpList[i].Left)
            }
            if tmpList[i].Right != nil{
                tmp = append(tmp, tmpList[i].Right)
            }
        }
        ret = append(ret, tmpList[len(tmpList) - 1].Val)
        tmpList = tmp
    }
    return ret
}