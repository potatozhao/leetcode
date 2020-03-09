/*
给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。

说明：

拆分时可以重复使用字典中的单词。
你可以假设字典中没有重复的单词。
示例 1：

输入: s = "leetcode", wordDict = ["leet", "code"]
输出: true
解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
示例 2：

输入: s = "applepenapple", wordDict = ["apple", "pen"]
输出: true
解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
     注意你可以重复使用字典中的单词。
示例 3：

输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
输出: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/word-break
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

// dp解法， 类似于背包。
func wordBreak(s string, wordDict []string) bool {
	dpList := make([]bool, len(s)+1)
	wordMap := make(map[string]bool)
	for i, _ := range wordDict {
		wordMap[wordDict[i]] = true
	}
	dpList[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if _, ok := wordMap[s[j:i]]; ok && dpList[j] {
				dpList[i] = true
				break
			}
		}
	}
	return dpList[len(s)]
}

//dfs + 剪枝
func wordBreak2(s string, wordDict []string) bool {
	tmpMap := make(map[string]bool)
	for _, v := range wordDict {
		tmpMap[v] = true
	}
	cacheMap := make(map[int]bool)
	return word_break(s, tmpMap, 0, cacheMap)
}

func word_break(s string, wordMap map[string]bool, start int, cacheMap map[int]bool) bool {
	if start >= len(s) {
		return true
	}
	res := false
	for i := start + 1; i <= len(s); i++ {
		if r, ok := cacheMap[i]; ok {
			res = r || res
		} else if _, ok := wordMap[s[start:i]]; ok {
			res = word_break(s, wordMap, i, cacheMap) || res
		}
	}
	cacheMap[start] = res
	return res
}
