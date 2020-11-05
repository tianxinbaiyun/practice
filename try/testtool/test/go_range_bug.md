@[TOC]Range 循环取地址BUG以及解决方法

## 1.BUG 重现

代码

```go
package main

import "fmt"

type Record struct {
	ID int
}

func main() {
	records := make([]Record, 0)
	list := make([]*Record, 0)
	for i := 0; i < 5; i++ {
		records = append(records, Record{
			ID: i,
		})
	}
	for i, record := range records {
		fmt.Printf("record[%d](v:%v,p:%p),ID:(v:%v,p:%p)\n", i, record, &record, record.ID, &record.ID)
		fmt.Printf("list[%d](v:%v,p:%p),ID:(v:%v,p:%p)\n", i, records[i], &records[i], records[i].ID, &records[i].ID)
		list = append(list, &record)
	}
	return
}
```

输出情况
```text
record[0](v:{0},p:0xc00000a0f0),ID:(v:0,p:0xc00000a0f0)
list[0](v:{0},p:0xc00000e1c0),ID:(v:0,p:0xc00000e1c0)
record[1](v:{1},p:0xc00000a0f0),ID:(v:1,p:0xc00000a0f0)
list[1](v:{1},p:0xc00000e1c8),ID:(v:1,p:0xc00000e1c8)
record[2](v:{2},p:0xc00000a0f0),ID:(v:2,p:0xc00000a0f0)
list[2](v:{2},p:0xc00000e1d0),ID:(v:2,p:0xc00000e1d0)
record[3](v:{3},p:0xc00000a0f0),ID:(v:3,p:0xc00000a0f0)
list[3](v:{3},p:0xc00000e1d8),ID:(v:3,p:0xc00000e1d8)
record[4](v:{4},p:0xc00000a0f0),ID:(v:4,p:0xc00000a0f0)
list[4](v:{4},p:0xc00000e1e0),ID:(v:4,p:0xc00000e1e0)
[0xc00000a0f0 0xc00000a0f0 0xc00000a0f0 0xc00000a0f0 0xc00000a0f0]

Process finished with exit code 0

```

## 2.解决方法

### 2.1 循环体内使用原数据下标对象操作

### 2.2 在写入原数据时使用地址

### 2.3 循环体内部重新分配一个变量

## 总结
循环range的变量是一个临时的变量，其地址不随数据变化而变化，而其值跟随数据变化。
所以对循环条件上的的变量，避免对其取地址，因为他们的地址不变
