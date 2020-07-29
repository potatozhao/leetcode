/*
给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。

返回被除数 dividend 除以除数 divisor 得到的商。

整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2

 

示例 1:

输入: dividend = 10, divisor = 3
输出: 3
解释: 10/3 = truncate(3.33333..) = truncate(3) = 3
示例 2:

输入: dividend = 7, divisor = -3
输出: -2
解释: 7/-3 = truncate(-2.33333..) = -2

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/divide-two-integers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

class Solution {
public:
    int divide(int dividend, int divisor) {
        bool p = true;
        unsigned int di;
        unsigned int ds;
        if (dividend > 0){
            p =!p;
            di = -dividend;
        }

        if(divisor>0){
            p=!p;
            divisor = -divisor;
        }
        int ret =0;
        while(dividend <= divisor){
            ret++;
            dividend-=divisor;
        }
        if(p) return ret;
        return -ret;
    }
};


/* 第二种方式

*/

class Solution {
public:
    int divide(int dividend, int divisor) {
        bool p = true;
        long a = dividend;
        long b = divisor;
        p=b>0?p:!p;
        p=a>0?p:!p;
        a=a>0?a:-a;
        b=b>0?b:-b;
        long ret = dfs(a,b);
        if (!p) ret = -ret;
        if (ret < INT_MIN) return INT_MIN;
        if (ret > INT_MAX) return INT_MAX;
        return ret;
    }
    // 递归解决。
    long dfs(long a, long b){  // 似乎精髓和难点就在于下面这几句
        if(a<b) return 0;
        long count = 1;
        long tb = b; // 在后面的代码中不更新b
        while((tb+tb)<=a){
            count = count + count; // 最小解翻倍
            tb = tb+tb; // 当前测试的值也翻倍
        }
        return count + dfs(a-tb,b);
    }
};