package main

import (
	"fmt"
)

// Node Node
type Node struct {
	Data  int
	Point *Node
	Next  *Node
}

// InitLinkNode 初始化列表
func InitLinkNode(n int) []*Node {
	var linkNode []*Node
	for i := 1; i <= n; i++ {
		node := &Node{
			Data:  i,
			Point: nil,
			Next:  nil,
		}
		linkNode = append(linkNode, node)
	}
	for i := range linkNode {
		switch i {
		case n - 1:
			linkNode[i].Next = linkNode[0]
		default:
			linkNode[i].Next = linkNode[i+1]
		}
		linkNode[i].Point = linkNode[i]
	}
	return linkNode
}

//DropNode 删除节点
func DropNode(i int, list []*Node) (result []*Node, err error) {
	var (
		pre    = i
		next   = i + 1
		length = len(list)
	)
	if i > length-1 {
		err = fmt.Errorf("i:%d不能为大于列表下标:%d", i, length)
		return
	}
	if i == 0 {
		result = list[1:]
	} else if i == length-1 {
		result = list[:length-1]
	} else {
		list[pre].Next = list[next]
		result = append(list[next:], list[:pre]...)
	}
	result[len(result)-1].Next = result[0]
	return
}

// Josephus 约瑟夫环形问题
// number 参加人数
// 焦点数
// 剩余人数
func Josephus(number int, focus int, surplus int) (list []*Node, err error) {
	list = InitLinkNode(number)
	for len(list) > surplus {
		length := len(list)
		if length > focus {
			list, err = DropNode(focus-1, list)
			if err != nil {
				panic(err)
			}
		} else {
			i := focus % length
			if i == 0 {
				i = length
			}
			list, err = DropNode(i-1, list)
		}
		fmt.Println()
		for _, i2 := range list {
			fmt.Print(i2)
		}
	}

	return
}

func main() {
	list, err := Josephus(10, 3, 2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println()
	fmt.Println(list)
}
