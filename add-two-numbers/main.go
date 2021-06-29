package main

import (
    "fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var res *ListNode
    r := (l1.Val + l2.Val) % 10
    i := 0
    if l1.Val + l2.Val >= 10 {
        i = 1
    }
    res = &ListNode{Val: r, Next: nil}
    for l1.Next != nil || l2.Next != nil || i != 0 {
        r = i
        if l1.Next != nil {
            l1 = l1.Next
            r += l1.Val
        }
        if l2.Next != nil {
            l2 = l2.Next
            r += l2.Val
        }
        i = 0
        if r >= 10 {
            i = 1
        }
        tt := &ListNode{ Val: r % 10, Next: nil }
        tt.Next = res
        res = tt
    }
    res2 := &ListNode{Val: res.Val, Next: nil}
    for res.Next != nil {
        tt := &ListNode{Val: res.Next.Val, Next: res2}
        res2 = tt
        res = res.Next
    }
    return res2
}

func main() {
    first := ListNode { Val: 2, Next: nil }
    first.Next = &ListNode { Val: 4, Next: nil }
    first.Next.Next = &ListNode { Val: 3, Next: nil }
    second := ListNode { Val: 5, Next: nil }
    second.Next = &ListNode { Val: 6, Next: nil }
    second.Next.Next = &ListNode { Val: 4, Next: nil }
    fmt.Println(*addTwoNumbers(&first, &second));
}
