package main

import (
    "fmt"
)

func preAppend(s []int, x int) []int {
    s = append(s, 0)
    copy(s[1:], s[0:])
    s[0] = x
    return s
}

func plusOne(digits []int) []int {
    result := []int{}
    length := len(digits)
    r := 1
    for i := length - 1; i >= 0; i -= 1 {
        result = preAppend(result, (digits[i] + r) % 10)
        r = (digits[i] + r) / 10
    }
    if r > 0 {
        result = preAppend(result, 1)
    }
    return result
}

func main() {
    fmt.Println(plusOne([]int{1, 2, 3}));
    fmt.Println(plusOne([]int{4, 3, 2, 1}));
    fmt.Println(plusOne([]int{0}));
    fmt.Println(plusOne([]int{9}));
}
