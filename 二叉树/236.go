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