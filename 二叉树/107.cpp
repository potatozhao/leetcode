/*
107. 二叉树的层次遍历 II
难度
简单

277





给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其自底向上的层次遍历为：

[
  [15,7],
  [9,20],
  [3]
]
*/


/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Solution {
public:
    vector<vector<int>> levelOrderBottom(TreeNode* root) {
        vector<TreeNode*>* a = new vector<TreeNode*>;
        vector<TreeNode*>* b = new vector<TreeNode*>;
        vector<vector<int>> ret;
        if(root == NULL){
            return ret;
        }
        a->push_back(root);
        while(a->size() > 0){
            vector<int> tmp;
            swap(a,b);
            a->clear();
            for(auto x:*b){
                if (x->left!=NULL) a->push_back(x->left);
                if (x->right!=NULL) a->push_back(x->right);
                tmp.push_back(x->val);
            }
            ret.push_back(tmp);
        }
        reverse(ret.begin(), ret.end());
        return ret;
    }
};