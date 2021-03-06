/*
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

 

示例:

给定 1->2->3->4, 你应该返回 2->1->4->3.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/swap-nodes-in-pairs
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
 // 加表头法
 func swapPairs(head *ListNode) *ListNode {
    thead := &ListNode{
        Val:0,
        Next:head,
    }
    s := thead
    a := thead.Next
    if a == nil{
        return thead.Next
    }
    b := a.Next
    for a != nil && b != nil{
        s.Next = b
        a.Next = b.Next
        s.Next.Next = a
        s = a
        if s.Next == nil{
            break
        }
        a = s.Next
        b = s.Next.Next
    }
    return thead.Next
}

// 递归解法
func dfs(head *ListNode) *ListNode{
    if head == nil || head.Next == nil{
        return head
    }
    next := head.Next
    head.Next = dfs(next.Next
    next.Next = head
    return next
}