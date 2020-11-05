@[TOC](Golang中Slice的append详解)


## 1.描述

切片拥有 长度 和 容量。

切片的长度就是它所包含的元素个数。

切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。

切片 s 的长度和容量可通过表达式 len(s) 和 cap(s) 来获取。

你可以通过重新切片来扩展一个切片，给它提供足够的容量。试着修改示例程序中的切片操作，向外扩展它的容量，看看会发生什么。

```go
package main

import "fmt"

func main() {
	// 以两倍扩容
	s := []int{1, 2, 3, 4}
	printSlice(s)

	s = append(s, 9)
	printSlice(s)

	s = append(s, 11)
	printSlice(s)
	// 2倍扩容也不满足时
	s1 := []int{1}
	s1 = append(s1, 1, 2, 3)
	printSlice(s1)

	s1 = append(s1, 1, 2, 3, 4, 5)
	printSlice(s1)
}

func printSlice(s []int) {
	fmt.Printf("slice:%+v,point:%p,len:%d,cap:%d\n", s, &s, len(s), cap(s))
}

```

运行结果

```
slice:[1 2 3 4],point:0xc00000c060,len:4,cap:4
slice:[1 2 3 4 9],point:0xc00000c0a0,len:5,cap:8
slice:[1 2 3 4 9 11],point:0xc00000c0e0,len:6,cap:8
slice:[1 1 2 3],point:0xc00000c120,len:4,cap:4
slice:[1 1 2 3 1 2 3 4 5],point:0xc00000c160,len:9,cap:10

```

## 2.底层扩容

```go
// grow grows the slice s so that it can hold extra more values, allocating
// more capacity if needed. It also returns the old and new slice lengths.
func grow(s Value, extra int) (Value, int, int) {
    i0 := s.Len()
    i1 := i0 + extra
    if i1 < i0 {
        panic("reflect.Append: slice overflow")
    }
    m := s.Cap()
    if i1 <= m {
        return s.Slice(0, i1), i0, i1
    }
    if m == 0 {
        m = extra
    } else {
        for m < i1 {
            if i0 < 1024 {
                m += m
            } else {
                m += m / 4
            }
        }
    }
    t := MakeSlice(s.Type(), i1, m)
    Copy(t, s)
    return t, i0, i1
}

// Append appends the values x to a slice s and returns the resulting slice.
// As in Go, each x's value must be assignable to the slice's element type.
func Append(s Value, x ...Value) Value {
    s.mustBe(Slice)
    s, i0, i1 := grow(s, len(x))
    for i, j := i0, 0; i < i1; i, j = i+1, j+1 {
        s.Index(i).Set(x[j])
    }
    return s
}
```

## 3.总结

1.常规扩容
使用append对切片进行添加时：
i.如果切片容量足够时，切片在原始数组上追加元素并返回一个新的 slice，底层数组不变
i.如果切片容量不足且切片原先切片容量为0时，切片容量为切片长度
i.如果切片容量不足且切片原先切片容量非0时
    ii.如果切片长度不超1024，容量翻倍
    ii.如果切片长度超过1024，容量按四分之一扩展

2.当追加长度大于切片长度时，容量为偶数，为新切片长度当最小偶数。（至于为什么，为也分不清）

3.GoLang中的切片扩容机制，与切片的数据类型、原本切片的容量、所需要的容量都有关系，比较复杂。
对于常见数据类型，在元素数量较少时，大致可以认为扩容是按照翻倍进行的。但具体情况需要具体分析。

