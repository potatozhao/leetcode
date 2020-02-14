/*
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

例如，给定如下二叉树:  root = [3,5,1,6,2,0,8,null,null,7,4]



 

示例 1:

输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
输出: 3
解释: 节点 5 和节点 1 的最近公共祖先是节点 3。
示例 2:

输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
输出: 5
解释: 节点 5 和节点 4 的最近公共祖先是节点 5。因为根据定义最近公共祖先节点可以为节点本身。
 

说明:

所有节点的值都是唯一的。
p、q 为不同节点且均存在于给定的二叉树中。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/


// 我自己写的方法太low了。
/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || (q == nil && p == nil){
		return nil
	}
	if q == nil{
	   return p
	}
   if p == nil{
	   return q
	}
	_,n := findFirst(root, p.Val ,q.Val)
	return n
}

func findFirst(node *TreeNode, p int, q int) (bool, *TreeNode) {
   if node == nil{
	   return false, nil
   }
   if node.Val == p{
	   if findSecond(node, q){
		   return true, node
	   }
	   return true, nil
   }
   r, n := findFirst(node.Left, p, q)
   if r {
	   if n != nil{
		   return true, n
	   }
	   if node.Val == q || findSecond(node.Right, q){
		   return true, node
	   }
   }
   ret := r
   r, n = findFirst(node.Right, p, q)
   if r {
	   if n != nil{
		   return true, n
	   }
	   if node.Val == q|| findSecond(node.Left, q){
		   return true, node
	   }
   }
   if !ret{
	   ret = r
   }
   return ret, nil
}

func findSecond(node *TreeNode, q int) bool {
   if node == nil{
	   return false
   }
   if node.Val == q{
	   return true
   }else{
	   if findSecond(node.Left, q) || findSecond(node.Right, q){
		   return true
	   }
	   return false
   }
}



// 题解的，注意看说明！
/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return dfs(root, p, q)
}

func dfs(node, p, q *TreeNode) *TreeNode{
   if node == nil || node == q || node == p {
	   return node
   }
   left := dfs(node.Left, p, q)
   right := dfs(node.Right, p, q)
   if left == nil && right == nil{
	   return nil
   }else if left != nil && right != nil{
	   return node
   }else if left != nil{
	   return left
   }else{
	   return right
   }
}