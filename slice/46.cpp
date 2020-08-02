/*
46. 全排列
难度
中等

819





给定一个 没有重复 数字的序列，返回其所有可能的全排列。

示例:

输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
*/


class Solution {
public:
    vector<vector<int>> permute(vector<int>& nums) {
        vector<vector<int>> ret;
        dfs(0, nums.size()-1, nums, ret);
        return ret;
    }

    int dfs(int begin, int end, vector<int> &nums, vector<vector<int>> &ret){
        if(begin == end){
            ret.push_back(nums);
            return 0;
        }
        for(int i = begin; i<= end; i++){
            swap(nums[begin], nums[i]);
            dfs(begin+1, end, nums, ret);
            swap(nums[begin],nums[i]);
        }
        return 0;
    }
};