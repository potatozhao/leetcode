/*
在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。

示例 1:

输入: 4->2->1->3
输出: 1->2->3->4
示例 2:

输入: -1->5->3->4->0
输出: -1->0->3->4->5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sort-list
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
func sortList(head *ListNode) *ListNode {
	ret := sort(head)
	return ret
}

func sort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	right := cutLinked(head)
	l := sort(head)
	r := sort(right)
	var ret *ListNode
	next := &ret
	tmpl := l
	tmpr := r
	for tmpl != nil && tmpr != nil {
		var temp *ListNode
		if tmpl.Val < tmpr.Val {
			temp = tmpl
			tmpl = tmpl.Next
		} else {
			temp = tmpr
			tmpr = tmpr.Next
		}
		temp.Next = nil
		*next = temp
		next = &((*next).Next)
	}
	*next = tmpr
	if tmpl != nil {
		*next = tmpl
	}
	return ret
}

func cutLinked(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow := head
	fast := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	ret := slow.Next
	slow.Next = nil
	return ret
}
