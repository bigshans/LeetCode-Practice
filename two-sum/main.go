package main;

import "fmt"

func twoSum(nums []int, target int) []int {
    a := map[int]int{}
    for k, v := range nums {
        if t, ok := a[target - v]; ok {
            return []int{t, k};
        } else {
            a[v] = k
        }
    }
    return []int{0, 0}
}

func main() {
    fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
    fmt.Println(twoSum([]int{3, 2, 4}, 6))
    fmt.Println(twoSum([]int{3, 3}, 6))
}
