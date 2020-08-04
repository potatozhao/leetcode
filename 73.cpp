/*
73. 矩阵置零
难度
中等

260





给定一个 m x n 的矩阵，如果一个元素为 0，则将其所在行和列的所有元素都设为 0。请使用原地算法。

示例 1:

输入: 
[
  [1,1,1],
  [1,0,1],
  [1,1,1]
]
输出: 
[
  [1,0,1],
  [0,0,0],
  [1,0,1]
]
示例 2:

输入: 
[
  [0,1,2,0],
  [3,4,5,2],
  [1,3,1,5]
]
输出: 
[
  [0,0,0,0],
  [0,4,5,0],
  [0,3,1,0]
]
进阶:

一个直接的解决方案是使用  O(mn) 的额外空间，但这并不是一个好的解决方案。
一个简单的改进方案是使用 O(m + n) 的额外空间，但这仍然不是最好的解决方案。
你能想出一个常数空间的解决方案吗？
*/


/*
利用边界进行标记，首先先记录边界是否需要置空，再对边界外的进行标记，再回头置空边界。
*/

class Solution {
public:
    void setZeroes(vector<vector<int>>& matrix) {
        bool x = false, y = false;
        for(int i = 0; i<matrix.size();i++){
            for(int j = 0; j<matrix[0].size();j++){
                // 标记边界是否置空
                if(i ==0 || j == 0){
                    if (matrix[i][j] == 0 && i==0){
                        x = true;
                    }
                    if (matrix[i][j] == 0 && j==0){
                        y = true;
                    }
                }else if(matrix[i][j] == 0){
                    //对边界外，利用边界进行标记
                    matrix[i][0] = 0;
                    matrix[0][j] = 0;
                }
            }
        }
        // 利用边界，对边界外的元素进行置零
        for(int i = 1; i<matrix.size();i++){
            if(matrix[i][0] == 0){
                for(int j = 1; j<matrix[0].size();j++){
                    matrix[i][j] = 0;
                }
            }
        }        
        for(int j = 1; j<matrix[0].size();j++){
            if(matrix[0][j] == 0){
                for(int i = 1; i<matrix.size();i++){
                    matrix[i][j] = 0;
                }
            }
        }
        // 置空边界
        if(x){
            for(int i = 0; i<matrix[0].size();i++){
                matrix[0][i] = 0;
            }
        }
        if(y){
            for(int i = 0; i<matrix.size();i++){
                matrix[i][0] = 0;
            }
        }

        return ;
    }
};