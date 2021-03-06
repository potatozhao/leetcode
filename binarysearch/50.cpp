/*

50. Pow(x, n)
难度
中等

471





实现 pow(x, n) ，即计算 x 的 n 次幂函数。

示例 1:

输入: 2.00000, 10
输出: 1024.00000
示例 2:

输入: 2.10000, 3
输出: 9.26100
示例 3:

输入: 2.00000, -2
输出: 0.25000
解释: 2-2 = 1/22 = 1/4 = 0.25
说明:

-100.0 < x < 100.0
n 是 32 位有符号整数，其数值范围是 [−231, 231 − 1] 。

*/



// 用long是为了解决int的边界问题。
class Solution {
public:
    double dfs(double x, long n){
        if(n == 0){
            return 1.0;
        }
        double y = dfs(x, n/2);
        return n%2==0?y*y:y*y*x;
    }

    double bfs(double x, long n){
        double ans = 1.0;
        double x_last = x;
        while(n>0){
            if(n%2==1){
                 ans *= x_last;
            }
            x_last *= x_last;
            n /= 2;
        }
        return ans;
    }

    double myPow(double x, int n) {
        double ret = 1;
        long t = n;
        if(x==0) return 0;
        return n>0?dfs(x,t):dfs(x,-t);
    }
};

