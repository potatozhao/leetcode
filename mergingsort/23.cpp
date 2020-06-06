/*
合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。

示例:

输入:
[
  1->4->5,
  1->3->4,
  2->6
]
输出: 1->1->2->3->4->4->5->6

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-k-sorted-lists
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/


/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */

#include <vector>
#include <queue>
using namespace std;

 struct ListNode {
     int val;
     ListNode *next;
     ListNode(int x) : val(x), next(NULL) {}
 };


struct cmp{
    bool operator()(ListNode* x, ListNode* y){
        return x->val > y->val;
    }
};

class Solution {
public:
    ListNode* mergeKLists(vector<ListNode*>& lists) {
        priority_queue<ListNode*, vector<ListNode*>, cmp> q;
        for(auto l :lists){
            if(l!= NULL){
                q.push(l);
            }
        }
        ListNode* ret = new ListNode(0);
        ListNode* last = ret;
        while(!q.empty()){
            last->next = q.top();
            last = last->next;
            q.pop();
            if (last->next !=NULL){
                q.push(last->next);
            }
        }
        return ret->next;
    }
};