/*
74. 搜索二维矩阵
难度
中等

216





编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：

每行中的整数从左到右按升序排列。
每行的第一个整数大于前一行的最后一个整数。
示例 1:

输入:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 3
输出: true
示例 2:

输入:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 13
输出: false
*/

/*
二维矩阵拉平成一维去搜索。
*/

class Solution {
public:
    bool searchMatrix(vector<vector<int>>& matrix, int target) {
        auto m = matrix.size();
        if (m<=0){
            return false;
        }
        auto n = matrix[0].size();
        if(n<=0){
            return false;
        }
        int start = 0, end = m*n-1;   
        while(start <= end){
            auto mid = (start+end)/2;
            if(matrix[mid/n][mid%n] == target){
                return true;
            }else if(matrix[mid/n][mid%n] > target){
                end = mid -1;
            }else{
                start = mid + 1;
            }
        }
        return false;
    }
};