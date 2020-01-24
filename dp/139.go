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
