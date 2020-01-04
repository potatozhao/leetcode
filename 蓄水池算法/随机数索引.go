import (
    "time"
    "math/rand"
)

type Solution struct {
    nums []int
}


func Constructor(nums []int) Solution {
    ret := Solution{nums:nums}
    return ret
}


func (this *Solution) Pick(target int) int {
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    idx := -1
    n := 0
    for i,_ := range(this.nums){
        if this.nums[i] == target{
            n += 1
            if r.Intn(n) == n -1{
                idx = i
            }
        }
    }
    return idx
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */