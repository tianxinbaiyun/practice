@[TOC]( for 取地址)
## 以下代码有什么问题，说明原因


```text
package main
import (
    "fmt"
)
type student struct {
    Name string
    Age  int
}
func pase_student() map[string]*student {
    m := make(map[string]*student)
    stus := []student{
        {Name: "zhou", Age: 24},
        {Name: "li", Age: 23},
        {Name: "wang", Age: 22},
    }
    for _, stu := range stus {
        m[stu.Name] = &stu
    }
    return m
}
func main() {
    students := pase_student()
    for k, v := range students {
        fmt.Printf("key=%s,value=%v \n", k, v)
    }
}

```

### 运行结果

```
key=zhou,value=&{wang 22} 
key=li,value=&{wang 22} 
key=wang,value=&{wang 22} 

Process finished with exit code 0
```
## 修改

### 代码修改

```text
package main
import (
"fmt"
)
type student struct {
	Name string
	Age  int
}
func pase_student() map[string]*student {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}

	for _, stu := range stus {
		fmt.Printf("%v\t%p\n",stu,&stu)
		tmp := stu
		m[stu.Name] = &tmp
	}
	return m
}
func main() {
	students := pase_student()
	for k, v := range students {
		fmt.Printf("key=%s,value=%v \n", k, v)
	}
}

```

### 运行结果

```
{zhou 24}       0xc000090020
{li 23} 0xc000090020
{wang 22}       0xc000090020
key=zhou,value=&{zhou 24} 
key=li,value=&{li 23} 
key=wang,value=&{wang 22} 

```

## 总结

通过上面的案例，我们不难发现stu变量的地址始终保持不变，每次遍历仅进行struct值拷贝，
故m[stu.Name]=&stu实际上一直指向同一个地址，最终该地址的值为遍历的最后一个struct的值拷贝。

