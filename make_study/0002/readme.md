## 前言
Go语言作为编程语言中的后起之秀，在博采众长的同时又不失个性，在注重运行效率的同时又重视开发效率，不失为一种好的开发语言。在go语言中，没有类的概念，但是仍然可以用struct+interface来实现类的功能，下面的这个简单的例子演示了如何用go来模拟c++中的多态的行为。
示例代码
```text


package main
 
 
import "os"
import "fmt"
 
 
type Human interface {
  sayHello()
}
 
 
type Chinese struct {
  name string
}
 
 
type English struct {
  name string
}
 
 
func (c *Chinese) sayHello() {
  fmt.Println(c.name,"说：你好，世界")
}
 
 
func (e *English) sayHello() {
  fmt.Println(e.name,"says: hello,world")
}
 
 
func main() {
  fmt.Println(len(os.Args))
   
  c := Chinese{"汪星人"}
  e := English{"jorn"}
  m := map[int]Human{}
   
  m[0] = &c
  m[1] = &e
   
  for i:=0;i<2;i++ {
    m[i].sayHello()
  }
}
```
总结
从上面的例子来看，在go中实现类似C++中的多态可谓是非常的简单，只要实现相同的接口即可。