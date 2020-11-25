@[TOC](约瑟夫问题)

## 约瑟夫问题

约瑟夫问题（有时也称为约瑟夫斯置换，是一个出现在计算机科学和数学中的问题。在计算机编程的算法中，类似问题又称为约瑟夫环。又称“丢手绢问题”.）

## 实现代码

```text
package main

import (
	"errors"
	"fmt"
)

type Node struct {
	Data int
	Point *Node
	Next *Node
}

//初始化列表
func InitLinkNode(n int)[]*Node  {
	var linkNode []*Node
	for i:=1;i<=n ;i++  {
		node := &Node{
			Data: i,
			Point:nil,
			Next: nil,
		}
		linkNode = append(linkNode,node)
	}
	for i, _ := range linkNode {
		switch i {
		case n-1:
			linkNode[i].Next = linkNode[0]
		default:
			linkNode[i].Next = linkNode[i+1]
		}
		linkNode[i].Point = linkNode[i]
	}
	return linkNode
}

//删除节点
func DropNode(i int,list []*Node)(result []*Node,err error)  {
	var (
		pre int = i
		next int = i+1
		length int = len(list)
	)
	if i>length-1{
		err = errors.New(fmt.Sprintf("i:%d不能为大于列表下标:%d",i,length))
		return
	}
	if i==0{
		result = list[1:]
	}else if i==length-1{
		result = list[:length-1]
	} else{
		list[pre].Next = list[next]
		result = append(list[next:],list[:pre]...)
	}
	result[len(result)-1].Next = result[0]
	return
}

// Josephus 约瑟夫环形问题
// number 参加人数
// 焦点数
// 剩余人数
func Josephus(number int,focus int,surplus int)  (list []*Node,err error){
	list = InitLinkNode(number)
	for len(list)>surplus{
		length := len(list)
		if length > focus{
			list,err = DropNode(focus-1,list)
			if err != nil{
				panic(err)
			}
		}else{
			i:=focus%length
			if i==0{
				i = length
			}
			list,err =  DropNode(i-1,list)
		}
		fmt.Println()
		for _, i2 := range list {
			fmt.Print(i2)
		}
	}

	return
}

func main()  {
	list,err := Josephus(10,3,2)
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println()
	fmt.Println(list)
}
```

## 运行结果


```
&{4 0xc00000c0e0 0xc00000c100}&{5 0xc00000c100 0xc00000c120}&{6 0xc00000c120 0xc00000c140}&{7 0xc00000c140 0xc00000c160}&{8 0xc00000c160 0xc00000c180}&{9 0xc00000c180 0xc00000c1a0}&{10 0xc00000c1a0 0xc00000c060}&{1 0xc00000c060 0xc00000c080}&{2 0xc00000c080 0xc00000c0e0}
&{7 0xc00000c140 0xc00000c160}&{8 0xc00000c160 0xc00000c180}&{9 0xc00000c180 0xc00000c1a0}&{10 0xc00000c1a0 0xc00000c060}&{1 0xc00000c060 0xc00000c080}&{2 0xc00000c080 0xc00000c0e0}&{4 0xc00000c0e0 0xc00000c100}&{5 0xc00000c100 0xc00000c140}
&{10 0xc00000c1a0 0xc00000c060}&{1 0xc00000c060 0xc00000c080}&{2 0xc00000c080 0xc00000c0e0}&{4 0xc00000c0e0 0xc00000c100}&{5 0xc00000c100 0xc00000c140}&{7 0xc00000c140 0xc00000c160}&{8 0xc00000c160 0xc00000c1a0}
&{4 0xc00000c0e0 0xc00000c100}&{5 0xc00000c100 0xc00000c140}&{7 0xc00000c140 0xc00000c160}&{8 0xc00000c160 0xc00000c1a0}&{10 0xc00000c1a0 0xc00000c060}&{1 0xc00000c060 0xc00000c0e0}
&{8 0xc00000c160 0xc00000c1a0}&{10 0xc00000c1a0 0xc00000c060}&{1 0xc00000c060 0xc00000c0e0}&{4 0xc00000c0e0 0xc00000c100}&{5 0xc00000c100 0xc00000c160}
&{4 0xc00000c0e0 0xc00000c100}&{5 0xc00000c100 0xc00000c160}&{8 0xc00000c160 0xc00000c1a0}&{10 0xc00000c1a0 0xc00000c0e0}
&{10 0xc00000c1a0 0xc00000c0e0}&{4 0xc00000c0e0 0xc00000c100}&{5 0xc00000c100 0xc00000c1a0}
&{10 0xc00000c1a0 0xc00000c0e0}&{4 0xc00000c0e0 0xc00000c1a0}
[0xc00000c1a0 0xc00000c0e0]


```

## 总结

约瑟夫问题延伸

| 总数N  | 直线(1始终排在首位) | 环形 |
|---------|---------|---------|
| 留杀型 | 留下1 | N=2^n,留下1  N≠2^n,留下(N-2^n)*2+1 |
| 杀留型  | 留下2^n(2^n<=N,且最接近N) | N=2^n, 留2^n  N≠2^n,留(N-2^n)*2| 
| 留杀杀型  | 留下1  | N=3^n 或者N=2*3^n,留1  N≠3^n或2*3^n, N为偶数,(N-2*3n)/2*3+1 N为奇数,(N-3n)/2*3+1 | 
| 杀杀留型  | 留下3^n或者2*3^n(3^n<=N或者2*3^n看哪个,且最接近N)  | N=3^n 或者N=2*3^n,留1  N≠3^n或2*3^n, N为偶数,(N-2*3n)/2*3 N为奇数,(N-3n)/2*3 | 



