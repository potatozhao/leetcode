/*
给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。

你应当保留两个分区中每个节点的初始相对位置。

示例:

输入: head = 1->4->3->2->5->2, x = 3
输出: 1->2->2->4->3->5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/partition-list
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
 func partition(head *ListNode, x int) *ListNode {
    if head == nil{
        return head
    }
    var left *ListNode
    var rignt *ListNode
    nextL := &left
    nextR := &rignt
    tmp := head
    for tmp != nil{
        temp := tmp.Next
        tmp.Next = nil
        if tmp.Val < x{
            *nextL = tmp
            nextL = &((*nextL).Next)
        }else if tmp.Val >= x{
            *nextR = tmp
            nextR = &((*nextR).Next)
        }
        tmp = temp
    }
    if left != nil{
        *nextL = rignt
        return left
    }
    return rignt
}