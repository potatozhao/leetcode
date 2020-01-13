func searchRange(nums []int, target int) []int {
    if len(nums) <= 0{
        return []int{-1,-1}
    }
    s := binarySearch(nums, target, true)
    e := binarySearch(nums, target, false)
    return []int{s,e}
}


func binarySearch(list []int,key int, flag bool) int{
    start := 0
    end := len(list) - 1
    mid := -1
    for start <= end{
        mid = (start + end) >> 1
        if flag{
            if list[mid] >= key{
                end = mid -1
            }else{
                start = mid + 1
            }
        }else{
            if list[mid] > key{
                end = mid -1
            }else{
                start = mid +1
            }
        }
    }
    if list[mid] < key{
        mid += 1
    }else if list[mid] > key{
        mid -= 1
    }
    if mid < 0 || mid >= len(list) || list[mid] != key{
        return -1
    }
    return mid
}