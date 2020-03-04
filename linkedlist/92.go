/*
反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。

说明:
1 ≤ m ≤ n ≤ 链表长度。

示例:

输入: 1->2->3->4->5->NULL, m = 2, n = 4
输出: 1->4->3->2->5->NULL

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-linked-list-ii
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
 func reverseBetween(head *ListNode, m int, n int) *ListNode {
    if head == nil{
        return head
    }
    thead := new(ListNode)
    thead.Next = head
    start := thead
    tail := thead
    last := thead
    num := 0
    for tail != nil{
        if num == m-1{
            start = tail
        }
        if num == m{
            last = tail 
        }
        if num > m && num <= n{
            temp := tail.Next
            tail.Next =start.Next
            start.Next = tail
            last.Next = temp 
            tail = last
        }
        tail = tail.Next
        num++
    }
    return thead.Next
}