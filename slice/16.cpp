/*
16. 最接近的三数之和

给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。

例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.

与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum-closest
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

#include<vector>
#include<cmath>

using namespace std;
class Solution {
public:
    int threeSumClosest(vector<int>& nums, int target) {
        if (nums.size()<3){
            return 0;
        }
        sort(nums.begin(), nums.end());
        int min_num = INT_MAX, ret = 0;
        for(int i = 0; i < nums.size()-2; i++){
            int a = i+1, b = nums.size()-1;
            while(a < b){
                int tmp_sum = nums[a] + nums[b] + nums[i];
                int tmp = abs(target-tmp_sum);
                if (tmp < min_num){
                    min_num = tmp;
                    ret = tmp_sum;
                }
                if (min_num == 0){
                    return ret;
                }
                if (tmp_sum > target){
                    b--;
                }else{
                    a++;
                }
            }
        }
        return ret;
    }
};

