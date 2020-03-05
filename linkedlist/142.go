/*
给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。

为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/linked-list-cycle-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */

 // 快慢指针解法

// 假设第一次相遇时，slow走了 s = y + x步， fast 走了 2s = y + nr + x步
// 那么 2(y+x) = y + nr + x
// 那么 y = nr - x
// 因为slow已经走了x了，所以用一个 normal指针从头开始以1走，那么当normal和slow相遇时，slow一定走了nr圈，那么slow这个点，就是圆起点。

/*
 #include <iostream>
 struct ListNode *detectCycle(struct ListNode *head) {
	 if (head == NULL){
		 return NULL;
	 }
	 struct ListNode *slow = head;
	 struct ListNode *fast = head;
	 while(fast != NULL && fast->next != NULL){
		 slow = slow->next;
		 fast = fast->next->next;
		 if (slow == fast){
			 break;
		 }
	 }
	 if (fast == NULL || fast->next == NULL){
		 return NULL;
	 }
	 fast = head;
	 while(slow != fast){
		 fast = fast->next;
		 slow = slow->next;
	 }
	 return slow;
 }
*/