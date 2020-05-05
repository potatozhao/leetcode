/*
22. 括号生成
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

 

示例：

输入：n = 3
输出：[
       "((()))",
       "(()())",
       "(())()",
       "()(())",
       "()()()"
     ]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/generate-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

class Solution {
	public:
		vector<string> generateParenthesis(int n) {
			vector<string> ret;
			if(n <=0){
				return ret;
			}
			string cur;
			dfs(ret, cur, n,n);
			return ret;
		}
	
		void dfs(vector<string> &ret, string &cur, int i, int j){
			if (i==0 && j==0){
				ret.push_back(cur);
				return;
			}
			if (j>i){
				cur.push_back(')');
				dfs(ret, cur, i, j-1);
				cur.pop_back();
			}
			if(i >0){
				cur.push_back('(');
				dfs(ret,cur,i-1,j);
				cur.pop_back();
			}
			return;
		}
	};