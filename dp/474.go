package main

func findMaxForm(strs []string, m int, n int) int {
	zMap := make(map[string]int)
	iMap := make(map[string]int)
	for _, v := range strs {
		zNum := 0
		iNum := 0
		for i, _ := range v {
			if v[i] == '0' {
				zNum += 1
			} else {
				iNum += 1
			}
		}
		zMap[v] = zNum
		iMap[v] = iNum
	}
	MNList := make([][]int, m+1)
	for m, _ := range MNList {
		MNList[m] = make([]int, n+1)
	}
	for _, v := range strs {
		vm := zMap[v]
		vn := iMap[v]
		for i := m; i >= vm; i-- {
			for j := n; j >= vn; j-- {
				if i >= zMap[v] && j >= iMap[v] {
					MNList[i][j] = max(MNList[i][j], MNList[i-vm][j-vn]+1)
				}
			}
		}
	}
	return MNList[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
