/*
有 N 个网络节点，标记为 1 到 N。

给定一个列表 times，表示信号经过有向边的传递时间。 times[i] = (u, v, w)，其中 u 是源节点，v 是目标节点， w 是一个信号从源节点传递到目标节点的时间。

现在，我们向当前的节点 K 发送了一个信号。需要多久才能使所有节点都收到信号？如果不能使所有节点收到信号，返回 -1。

注意:

N 的范围在 [1, 100] 之间。
K 的范围在 [1, N] 之间。
times 的长度在 [1, 6000] 之间。
所有的边 times[i] = (u, v, w) 都有 1 <= u, v <= N 且 0 <= w <= 100。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/network-delay-time
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

// dfs解法
func networkDelayTime(times [][]int, N int, K int) int {
    timeList := make([]int, N+1)
    for i := range timeList{
        timeList[i] = -1
    }
    graphMap := make(map[int][]CostNode)
    for i := range times{
        s := times[i][0]
        e := times[i][1]
        c := times[i][2]
        if _, ok:= graphMap[s]; !ok{
            graphMap[s] = make([]CostNode, 0)
        }
        graphMap[s] = append(graphMap[s], CostNode{Next:e,Cost:c})
    }
    timeList[K] = 0
    dfs(graphMap, K, 0, timeList)
    maxTime := 0
    for i := range timeList{
        if i == 0{
            continue
        }
        if timeList[i] <0{
            return -1
        }
        if maxTime < timeList[i]{
            maxTime = timeList[i]
        }
    }
    return maxTime
}

func dfs(graphMap map[int][]CostNode, s , c int, timeList []int){
    v := graphMap[s]
    for i := range v{
        if timeList[v[i].Next] < 0 || timeList[v[i].Next] > v[i].Cost + c{
            timeList[v[i].Next] = v[i].Cost + c
            dfs(graphMap, v[i].Next, v[i].Cost + c, timeList)
        }
    }
}

type CostNode struct{
    Next int
    Cost int
}



//Dijkstra 算法
func networkDelayTimeD(times [][]int, N int, K int) int {
    distinceList := make([]int, N+1)
    for i := range distinceList{
        distinceList[i] = -1
    }
    graphMap := make([][]int,N+1)
    for i := range graphMap{
        graphMap[i] = make([]int, N+1)
        for j := range graphMap[i]{
            graphMap[i][j] = -1
        }
    }

    for i := range times{
        graphMap[times[i][0]][times[i][1]] = times[i][2]
    }


    for i := range graphMap[K]{
        distinceList[i] = graphMap[K][i]
    }
    distinceList[K] = 0

    visitedList := make([]bool, N+1)
    visitedList[K] = true

    for i := 0; i < N-1; i++{
        minDistince := int(^uint(0) >> 1)
        minNum := 0
        for j := 1; j <= N; j++{
            if !visitedList[j] && distinceList[j] >= 0 && distinceList[j] < minDistince{
                minNum = j
                minDistince = distinceList[j]
            }
        }

        for j :=1; j <= N; j++{
            if !visitedList[j] && graphMap[minNum][j] >=0{
                if distinceList[j] < 0{
                    distinceList[j] = minDistince + graphMap[minNum][j]
                }else if distinceList[j] > minDistince + graphMap[minNum][j]{
                    distinceList[j] = minDistince + graphMap[minNum][j]
                }
            }
        }
        visitedList[minNum] = true
    }
    maxNum := 0
    for i := 1; i < len(distinceList); i++{
        if distinceList[i] == -1{
            return -1
        }
        if distinceList[i] > maxNum{
            maxNum = distinceList[i]
        }
    }
    return maxNum
}

//floyd 算法 时间复杂度n的三次方
func networkDelayTimeF(times [][]int, N int, K int) int {
	graphMap := make([][]int,N+1)
	// 初始化邻接矩阵
    for i := range graphMap{
        graphMap[i] = make([]int, N+1)
        for j := range graphMap[i]{
            if i == j{
                graphMap[i][j] = 0
            }else{
                graphMap[i][j] = -1
            }
        }
    }

    for i := range times{
        graphMap[times[i][0]][times[i][1]] = times[i][2]
    }

    for m := 1; m<=N; m++{ // 依次使用1-N的节点对全部节点进行进行松弛。
        for i := 1; i <= N; i++{
            for j := 1; j<= N; j++{
                if graphMap[i][m] != - 1 && graphMap[m][j] != -1 && (graphMap[i][j] == -1 || graphMap[i][j] > graphMap[i][m] + graphMap[m][j]){
                graphMap[i][j] = graphMap[i][m] + graphMap[m][j]
                }
            }
        }
    }
    maxNum := 0
    for i := 1; i <= N; i++{
        if graphMap[K][i] == -1{
            return -1
        }
        if graphMap[K][i] > maxNum{
            maxNum = graphMap[K][i]
        }
    }
    return maxNum
}

// bellmen-ford算法，缺少负权环判断。
func networkDelayTimeB(times [][]int, N int, K int) int {

    distanceList := make([]int, N+1)
    for i := range distanceList{
        distanceList[i] = -1
    }
    distanceList[K] = 0
    for i := 0; i < N -1;i++{
        flag := true
        for j := range times{
            s := times[j][0]
            e := times[j][1]
            c := times[j][2]
            if distanceList[s] != -1 &&(distanceList[e] == -1 || distanceList[s] + c < distanceList[e]){
                distanceList[e] = distanceList[s] + c
                flag = false
            }
        }
        if flag{
            break
        }
    }
    maxNum := 0
    for i :=1; i < len(distanceList); i++{
        if distanceList[i] == -1{
            return -1
        }
        if distanceList[i] > maxNum{
            maxNum = distanceList[i]
        }
    }
    return maxNum
}