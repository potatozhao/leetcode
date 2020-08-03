/*
77. 组合
难度
中等

323





给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。

示例:

输入: n = 4, k = 2
输出:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
*/


class Solution {
public:
    vector<vector<int>> combine(int n, int k) {
        k = k>n?n:k;
        vector<int> nums;
        vector<vector<int>> ret;
        if (k<=0){
            return ret;
        }
        dfs(1, n, k,nums, ret);
        return ret;
        
    }
    int dfs(int begin, int end, int k, vector<int> &nums, vector<vector<int>> &ret){
        if(k==0){
            ret.push_back(nums);
            return 0;
        }
        for(int i =begin; i<= end-k+1; i++){
            nums.push_back(i);
            dfs(i+1, end, k-1, nums, ret);
            nums.pop_back();
        }
        return 0;
    }
};