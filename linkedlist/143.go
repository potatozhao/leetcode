/*
给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例 1:

给定链表 1->2->3->4, 重新排列为 1->4->2->3.
示例 2:

给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reorder-list
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

// 快慢指针找到中间点
func reorderList(head *ListNode) {
	// 1,2,3时的特殊处理。
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}
	// 快慢指针定位中间
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 翻转链表
	lastHead := slow.Next
	slow.Next = nil // 置空前半截链表
	next := lastHead.Next
	lastHead.Next = nil // 表尾置空
	for next != nil {
		tmp := next
		next = next.Next
		tmp.Next = lastHead
		lastHead = tmp
	}
	// 合并
	tmpF := head
	tmpL := lastHead
	for tmpF.Next != nil && tmpL != nil {
		tmp := tmpL.Next
		tmpL.Next = tmpF.Next
		tmpF.Next = tmpL
		tmpL = tmp
		tmpF = tmpF.Next.Next
	}
	// 这个看中间点位置。如果中间点在正中间，则用，如果不在，则不用。
	if tmpF.Next == nil {
		tmpF.Next = tmpL
	}
	return
}
