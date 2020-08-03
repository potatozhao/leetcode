/*
63. 不同路径 II
难度
中等

387





一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。

现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
*/


class Solution {
public:
    int uniquePathsWithObstacles(vector<vector<int>>& obstacleGrid) {
        vector<int> * a = new vector<int>(obstacleGrid[0].size()+1, 0);
        vector<int> * b = new vector<int>(obstacleGrid[0].size()+1, 0);
        for(int i = 1; i<=obstacleGrid.size(); i++){
            for(int j =1; j<=obstacleGrid[0].size(); j++){
                if(j == 1 && i == 1&& obstacleGrid[i-1][j-1] != 1){
                    (*a)[j] = 1;
                }else if(obstacleGrid[i-1][j-1] == 1){
                    (*a)[j] = 0;
                }else{
                    (*a)[j] = (*b)[j] + (*a)[j-1];
                }
            }
            swap(a, b);
        }
        return b->at(a->size()-1);
    }
};