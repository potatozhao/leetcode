/*
113. 路径总和 II
难度
中等
给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。

说明: 叶子节点是指没有子节点的节点。

示例:
给定如下二叉树，以及目标和 sum = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1
返回:

[
   [5,4,11,2],
   [5,8,4,5]
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
    vector<vector<int>> pathSum(TreeNode* root, int sum) {
        vector<int> path;
        vector<vector<int>> paths;
        dfs(root, sum, path, paths);
        return paths;
    }
    int dfs(TreeNode* node, int sum ,vector<int> &path, vector<vector<int>> &paths){
        if (node==nullptr){
            return 0;
        }
        path.push_back(node->val);
        if(node->left == nullptr && node->right == nullptr && sum-node->val ==0){
            paths.push_back(path);
            path.pop_back();
            return 0;
        }
        dfs(node->left, sum-node->val, path, paths);
        dfs(node->right, sum-node->val, path,paths);
        path.pop_back();
        return 0;
    }
};