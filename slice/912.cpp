#include<swap>

/*
给你一个整数数组 nums，请你将该数组升序排列。

 

示例 1：

输入：nums = [5,2,3,1]
输出：[1,2,3,5]
示例 2：

输入：nums = [5,1,1,2,0,0]
输出：[0,0,1,1,2,5]
 

提示：

1 <= nums.length <= 50000
-50000 <= nums[i] <= 50000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sort-an-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

class Solution {
	public:
		vector<int> sortArray(vector<int>& nums) {
			qsort2(nums, 0, nums.size()-1);
			return nums;
		}
		//正常快排
		int qsort(vector<int>& nums,int start, int end){
			if (start>=end){
				return 0;
			}
			int mid = partition2(nums, start, end);
			qsort(nums, start, mid - 1);
			qsort(nums, mid+1, end);
			return 0;
		}
        // 从左向右方法
		int partition(vector<int>& nums, int start, int end){
			std::swap(nums[start], nums[end]);
			int start_pos = start;
			for (int i = start; i < end; i++){
				if (nums[i] < nums[end]){
					std::swap(nums[i], nums[start_pos]);
					start_pos++;
				}
			}
			std::swap(nums[start_pos], nums[end]);
			return start_pos;
		}
        // 双向法
		int partition2(vector<int>& nums, int start, int end){
			int i = start, j = end+1;
			while(true){
				while(nums[++i] < nums[start]){
					if (i ==end){
						break;
					}
				}
				while(nums[--j] > nums[start]){
					if(j == start){
						break;
					}
				}
				if (i>=j){
					break;
				}
				std:swap(nums[i],nums[j]);
			}
			std::swap(nums[start], nums[j]);
			return j;
		}
        //三向切分
		int qsort2(vector<int>& nums, int start,int end){
			if (start >= end){
				return 0;
			}
			int i = start, j = end, m = start+1, v= nums[start];
			while (m <= j){
				if (nums[m] < v){
					std::swap(nums[m], nums[i]);
					i++;
					m++;
				}else if(nums[m] > v){
					std::swap(nums[m], nums[j]);
					j--;
				}else{
					m++;
				}
			}
			this->qsort2(nums, start, i-1);
			this->qsort2(nums, j+1, end);
			return 0;
		}
	};

    //堆排序
    class Solution {
public:
    vector<int> sortArray(vector<int>& nums) {
        for(int i =1; i<nums.size(); i++){
            swim(nums, i);
        }
        for(int i = nums.size()-1; i >0; i--){
            this->swap(nums[0], nums[i]);
            this->sink(nums, i-1);
        }

        return nums;
    }
    void swim(vector<int>& nums, int k){
        while(k >=1 && nums[k] > nums[(k-1)/2]){
            this->swap(nums[k], nums[(k-1)/2]);
            k = (k-1)/2;
        }
        return;
    }

    void sink(vector<int>& nums, int k){
        for (int i = 0; i <= (k-1)/2&&k > 0;){
            if(i*2 +2 <= k && nums[i*2+1] < nums[i*2+2] && nums[i] < nums[i*2+2]){
                this->swap(nums[i], nums[i*2+2]);
                i = i*2+2;
            }else if (nums[i] < nums[i*2+1]){
                this->swap(nums[i], nums[i*2 +1]);
                i = i*2+1;
            }else{
                break;
            }
        }
        return;
    }

    void swap(int &a, int& b){
        int t = a;
        a = b;
        b = t;
        return;
    }
};