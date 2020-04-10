/*
17. 电话号码的字母组合
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。



示例:

输入："23"
输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/


package main

func letterCombinations(digits string) []string {
    wordList := []string{
        "",
        "",
        "abc",
        "def",
        "ghi",
        "jkl",
        "mno",
        "pqrs",
        "tuv",
        "wxyz",
    }
    return dfs(wordList, digits, 0)
}

func dfs(wordList []string, digits string, pos int)[]string{
    ret := make([]string, 0)
    if pos >= len(digits){
        return ret
    }
    key := int(digits[pos] - '0')
    if pos == len(digits)- 1{
        for i := range wordList[key]{
            ret = append(ret, string(wordList[key][i]))
        }
        return ret
    }
    nextret := dfs(wordList, digits, pos + 1)
    for i := range nextret{
        for j := range wordList[key]{
            ret = append(ret, string(wordList[key][j]) + nextret[i])
        }
    }
    return ret
}