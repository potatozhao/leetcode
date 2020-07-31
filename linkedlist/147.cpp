/*
147. 对链表进行插入排序
*/

/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    ListNode* insertionSortList(ListNode* head) {
        ListNode* thead = new ListNode;
        ListNode* t = head;
        while(t != NULL){
            ListNode* tmp = t;
            t = t->next;
            tmp->next = NULL;
            insert(thead, tmp);
        }
        return thead->next;
    }
    int insert(ListNode* head, ListNode* node){
        ListNode* tmp = head;
        while(tmp->next != NULL){
            if(node->val < tmp->next->val){
                break;
            }
            tmp = tmp->next;
        }
        node->next = tmp->next;
        tmp->next = node;
        return 0;
    }
};