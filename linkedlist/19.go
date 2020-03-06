/*
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：

给定的 n 保证是有效的。

进阶：

你能尝试使用一趟扫描实现吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list
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
 func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if head == nil{
        return head
    }
    thead := new(ListNode)
    thead.Next = head
    slow := thead
    fast := thead
    for i := 0; i < n; i++{
        fast = fast.Next
        if fast == nil || fast.Next == nil{
            break
        }
    }
    for fast.Next != nil{
        slow = slow.Next
        fast = fast.Next
    }
    slow.Next = slow.Next.Next
    return thead.Next
}