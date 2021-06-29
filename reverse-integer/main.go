package main

import (
    "fmt"
)

func reverse(x int) int {
    if x == 0 {
        return 0
    }
    a := x
    base := 1
    for a != 0 {
        a /= 10
        base *= 10
    }
    r := 0
    for base != 0 {
        r += x % 10 * base
        x /= 10
        base /= 10
    }
    ans := r / 10
    if ans > 2147483647 || ans < -2147483648 {
        return 0
    }
    return ans
}

func main() {
    fmt.Println(reverse(123))
    fmt.Println(reverse(-123))
    fmt.Println(reverse(120))
    fmt.Println(reverse(0))
    fmt.Println(reverse(9000000))
}
