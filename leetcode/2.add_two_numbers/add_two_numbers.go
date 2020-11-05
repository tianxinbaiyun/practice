package main

import (
	"fmt"
)

// ListNode ListNode
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l10 := &ListNode{Val: 1}
	l11 := &ListNode{Val: 5}
	l12 := &ListNode{Val: 3}

	l20 := &ListNode{Val: 2}
	l21 := &ListNode{Val: 8}
	l22 := &ListNode{Val: 3}

	l10.Next = l11
	l11.Next = l12
	l20.Next = l21
	l21.Next = l22

	fmt.Println(addTwoNumbers(l10, l20))
	fmt.Println(addTwoNumbers(l11, l21))
	fmt.Println(addTwoNumbers(l12, l22))
}

//func main.go.go() {
//	la1 := make([]*ListNode, 0)
//	la2 := make([]*ListNode, 0)
//	la1 = append(la1, &ListNode{Val: 1})
//	la2 = append(la2, &ListNode{Val: 1})
//	for i := 1; i < 3; i++ {
//		la1 = append(la1, &ListNode{Val: rand.Intn(10)})
//		la2 = append(la2, &ListNode{Val: rand.Intn(9)})
//		la1[i-1].Next = la1[i]
//		la2[i-1].Next = la2[i]
//	}
//	for _, v := range la1 {
//		fmt.Println("v", v.Val)
//	}
//	for _, v := range la2 {
//		fmt.Println("v", v.Val)
//	}
//	res := addTwoNumbers(la1[0], la2[0])
//	fmt.Println(res)
//}

// addTwoNumbers addTwoNumbers
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l = &ListNode{}
	pre := l
	flag := 0
	for l1 != nil || l2 != nil {
		pre.Next = &ListNode{}
		p := pre.Next
		x := 0
		y := 0
		if l1 != nil {
			x = l1.Val
		}
		if l2 != nil {
			y = l2.Val
		}
		p.Val = (x + y + flag) % 10
		flag = (x + y + flag) / 10
		pre = p
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if flag != 0 {
		pre.Next = &ListNode{Val: flag}
	}

	return l.Next
}
