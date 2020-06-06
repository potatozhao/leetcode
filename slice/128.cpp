/*
给定一个未排序的整数数组，找出最长连续序列的长度。

要求算法的时间复杂度为 O(n)。

示例:

输入: [100, 4, 200, 1, 3, 2]
输出: 4
解释: 最长连续序列是 [1, 2, 3, 4]。它的长度为 4。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-consecutive-sequence
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/


#include<map>
#include<set>
#include<vector>
using namespace std;

class Solution {
public:
    int longestConsecutive(vector<int>& nums) {
        map<int,int> m;
        set<int> filter;
        for(auto v : nums){
            m[v] = 1;
        }
        int max_len= 0;
        for(auto v:nums){
            if(filter.find(v) != filter.end())
            {
                continue;
            }
            int num = v+1;
            while(m.find(num) != m.end()){
                m[v] += m[num];
                num += m[num];
            }
            max_len = max(max_len, m[v]);
            filter.insert(v);
        }
        return max_len;
    }
};