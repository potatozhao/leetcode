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
