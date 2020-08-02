/*
47. 全排列 II
难度
中等

366





给定一个可包含重复数字的序列，返回所有不重复的全排列。

示例:

输入: [1,1,2]
输出:
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]

*/


class Solution {
public:
    vector<vector<int>> permuteUnique(vector<int>& nums) {
        sort(nums.begin(), nums.end());
        vector<vector<int>> ret;
        ret.push_back(nums);
        while(next(nums)){
            ret.push_back(nums);
        }
        return ret;
    }
    bool next(vector<int>& nums){
        int cur = nums.size()-1;
        int pre = cur-1;
        while(cur>0 && nums[pre] >= nums[cur]) cur--,pre--;
        if(cur <=0 ) return false;
        for(cur =nums.size() -1;nums[cur]<=nums[pre];cur--){
        }
        swap(nums[cur], nums[pre]);
        reverse(nums.begin() + pre + 1, nums.end());
        return true;
    }
};