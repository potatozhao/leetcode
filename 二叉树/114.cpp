/*
114. 二叉树展开为链表
难度
中等

435





给定一个二叉树，原地将它展开为一个单链表。

 

例如，给定二叉树

    1
   / \
  2   5
 / \   \
3   4   6
将其展开为：

1
 \
  2
   \
    3
     \
      4
       \
        5
         \
          6
*/

/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    void flatten(TreeNode* root) {
        if(root == NULL) return;
        TreeNode* tmp = root;
        while(tmp !=NULL){
            if (tmp->left != NULL){
                TreeNode* l = tmp->left;
                TreeNode* r = tmp->right;
                while(l->right!=NULL){
                    l = l->right;
                }
                l->right = r;
                tmp->right = tmp->left;
                tmp->left = NULL;
            }
            tmp = tmp->right;
        }
        return;
    }
};