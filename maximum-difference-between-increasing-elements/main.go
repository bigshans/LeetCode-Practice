package main

import "fmt"

func min(a int, b int) int {
  if a < b {
    return a
  }
  return b
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}

func maximumDifference(nums []int) int {
  length := len(nums)
  m := -1
  for i := 1; i < length; i += 1 {
    for j := i - 1; j >= 0; j -= 1 {
      if nums[i] > nums[j] {
        m = max(m, nums[i] - nums[j])
      }
    }
  }
  return m
}

func main() {
  fmt.Println(maximumDifference([]int{ 7, 1, 5, 4 }))
  fmt.Println(maximumDifference([]int{ 9, 4, 3, 2 }))
}
