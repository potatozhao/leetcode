/*
给定两个非空链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储单个数字。将这两数相加会返回一个新的链表。



你可以假设除了数字 0 之外，这两个数字都不会以零开头。

进阶:

如果输入链表不能修改该如何处理？换句话说，你不能对列表中的节点进行翻转。

示例:

输入: (7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
输出: 7 -> 8 -> 0 -> 7

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 补齐递归法
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	addzero(&l1, &l2)
	num := dfs(l1, l2)
	if num > 0 {
		tmpNode := &ListNode{
			Val: num,
		}
		tmpNode.Next = l1
		l1 = tmpNode
	}
	return l1
}

func addzero(l1 **ListNode, l2 **ListNode) {
	tmp1 := *l1
	tmp2 := *l2
	num1 := 0
	num2 := 0
	for tmp1 != nil || tmp2 != nil {
		if tmp1 != nil {
			tmp1 = tmp1.Next
			num1++
		}
		if tmp2 != nil {
			tmp2 = tmp2.Next
			num2++
		}
	}
	if num2 > num1 {
		for i := 0; i < num2-num1; i++ {
			tmpNode := &ListNode{}
			tmpNode.Next = *l1
			*l1 = tmpNode
		}
	} else if num2 < num1 {
		for i := 0; i < num1-num2; i++ {
			tmpNode := &ListNode{}
			tmpNode.Next = *l2
			*l2 = tmpNode
		}
	}
	return
}

func dfs(l1 *ListNode, l2 *ListNode) int {
	if l1 == nil || l2 == nil {
		return 0
	}
	num := dfs(l1.Next, l2.Next)
	num += l1.Val
	num += l2.Val
	l1.Val = num % 10
	return num / 10
}
