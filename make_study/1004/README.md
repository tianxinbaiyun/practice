@[TOC]( 组合输出 )
## 4. 下面代码会输出什么？

```text

package main

import "fmt"

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func main() {
	t := Teacher{}
	t.ShowA()
}

```

### 运行结果

```
showA
showB

```

## 总结

Go中没有继承,上面这种写法叫组合。

上面的t.ShowA()等价于t.People.ShowA()