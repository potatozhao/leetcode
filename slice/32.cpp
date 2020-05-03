/*
32. 最长有效括号

给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。

示例 1:

输入: "(()"
输出: 2
解释: 最长有效括号子串为 "()"
示例 2:

输入: ")()())"
输出: 4
解释: 最长有效括号子串为 "()()"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-valid-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/


// dp解法
/*
if s[i] == '(' :
    dp[i] = 0
if s[i] == ')' :
    if s[i - 1] == '(' :
        dp[i] = dp[i - 2] + 2 #要保证i - 2 >= 0

    if s[i - 1] == ')' and s[i - dp[i - 1] - 1] == ')' :
        dp[i] = dp[i - 1] + dp[i - dp[i - 1] - 2] + 2 #要保证i - dp[i - 1] - 2 >= 0

作者：zhanganan0425
链接：https://leetcode-cn.com/problems/longest-valid-parentheses/solution/dong-tai-gui-hua-si-lu-xiang-jie-c-by-zhanganan042/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/



#include<string>
#include<vector>
using namespace std;
class Solution {
public:
    int longestValidParentheses(string s) {
        int ret = 0;
        vector<int> dp_list(s.size() + 1, 0);
        for (int i = 1; i< s.size(); i++){
            if(s[i] == ')'){
                if(s[i-1] == '('){
                    dp_list[i+1] = dp_list[i-1] + 2;
                }else{
                    if (i-1-dp_list[i] >=0 && s[i-1-dp_list[i]] == '('){
                        dp_list[i+1] = dp_list[i]+2 + dp_list[i-1-dp_list[i]];
                    }
                }
            }
            ret = max(dp_list[i+1], ret);
        }
        return ret;
    }
};

