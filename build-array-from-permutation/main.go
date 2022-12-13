package main

import "fmt"

func buildArray(nums []int) []int {
/*
    b := make([]int, len(nums))
    for i := range nums {
        b[i] = nums[nums[i]]
    }
    return b;
    length := len(nums)
*/
    length := len(nums)
    b := make([]int, length)
    for i := 0; i < length; i += 1 {
        b[i] = nums[nums[i]]
    }
    return b;
}

func main() {
    fmt.Println(buildArray([]int{0, 2, 1, 5, 3, 4}))
}
