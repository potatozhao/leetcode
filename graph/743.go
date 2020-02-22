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