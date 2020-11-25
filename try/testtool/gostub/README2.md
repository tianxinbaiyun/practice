@[TOC](GoStub应用)
## 1、GoStub应用场景
GoStub框架的使用场景如下：
A、为一个全局变量打桩
B、为一个函数打桩
C、为一个过程打桩
D、由任意相同或不同的基本场景组合而成

## 2、为全局变量打桩
假设counter为被测函数中使用的一个全局整型变量，
当前测试用例中假定counter的值为200，则打桩的代码如下：
```
package main

import (
   "fmt"

   "github.com/prashantv/gostub"
)

var counter = 100

func stubGlobalVariable() {
   stubs := gostub.Stub(&counter, 200)
   defer stubs.Reset()
   fmt.Println("Counter:", counter)
}

func main() {
   stubGlobalVariable()
}
```

```shell script
// output:
// Counter: 200
```

stubs是GoStub框架的函数接口Stub返回的对象，
Reset方法将全局变量的值恢复为原值。
```
package main

import (
   "io/ioutil"

   "fmt"

   "github.com/prashantv/gostub"
)

var configFile = "config.json"

func GetConfig() ([]byte, error) {
   return ioutil.ReadFile(configFile)
}

func stubGlobalVariable() {
   stubs := gostub.Stub(&configFile, "/tmp/test.config")
   defer stubs.Reset()
   /// 返回tmp/test.config文件的内容
   data, err := GetConfig()
   if err != nil {
      fmt.Println(err)
   }
   fmt.Println(data)
}

func main() {
   stubGlobalVariable()
}
```

## 3、为函数打桩
通常函数分为工程自定义函数与库函数。
假设工程中自定义函数如下：
```shell script
func Exec(cmd string, args ...string) (string, error) {
   ...
}
```

Exec函数是不能通过GoStub框架打桩的。
如果想要通过GoStub框架对Exec函数进行打桩，
则仅需对自定义函数进行简单的重构，即将Exec函数定义为匿名函数，同时将其赋值给Exec变量，重构后的代码如下：
```
var Exec = func(cmd string, args ...string) (string, error) {
   ...
}
```

当Exec函数重构成Exec变量后，并不影响既有代码中对Exec函数的调用。由于Exec变量是函数变量，因此一般函数变量也叫做函数。对Exec函数变量进行打桩的代码如下：
```shell script
stubs := Stub(&Exec, func(cmd string, args ...string) (string, error) {
   return "test", nil
})
defer stubs.Reset()
```

GoStub框架专门提供了StubFunc函数用于函数打桩，
对于函数的打桩代码如下：
```
stubs := StubFunc(&Exec,"test", nil)
defer stubs.Reset()
```

工程代码中会调用Golang库函数或第三方库函数，
由于不能重构库函数，因此需要在工程代码中增加一层适配层，
在适配层中定义库函数的变量，然后在工程代码中使用函数变量。
```
package Adapter

import (
   "time"

   "fmt"

   "os"

   "github.com/prashantv/gostub"
)

var timeNow = time.Now
var osHostname = os.Hostname

func getDate() int {
   return timeNow().Day()
}
func getHostName() (string, error) {
   return osHostname()
}

func StubTimeNowFunction() {
   stubs := gostub.Stub(&timeNow, func() time.Time {
      return time.Date(2015, 6, 1, 0, 0, 0, 0, time.UTC)
   })
   fmt.Println(getDate())
   defer stubs.Reset()
}

func StubHostNameFunction() {
   stubs := gostub.StubFunc(&osHostname, "LocalHost", nil)
   defer stubs.Reset()
   fmt.Println(getHostName())
}
```

使用示例：
```
package main

import "GoExample/GoStub/StubFunction"

func main() {
   Adapter.StubTimeNowFunction()
   Adapter.StubHostNameFunction()
}
```

## 4、为过程打桩
没有返回值的函数称为过程。通常将资源清理类函数定义为过程。
```
package main

import (
   "fmt"

   "github.com/prashantv/gostub"
)

var CleanUp = cleanUp

func cleanUp(val string) {
   fmt.Println(val)
}

func main() {
   stubs := gostub.StubFunc(&CleanUp)
   CleanUp("Hello go")
   defer stubs.Reset()
}
```

## 5、复杂场景
不论是调用Stub函数还是StubFunc函数，
都会生成一个Stubs对象，
Stubs对象仍然有Stub方法和StubFunc方法，
所以在一个测试用例中可以同时对多个全局变量、函数或过程打桩。
全局变量、函数或过程会将初始值存在一个map中，
并在延迟语句中通过Reset方法统一做回滚处理。
多次打桩代码如下：
```shell script
stubs := gostub.Stub(&v1, 1)
defer stubs.Reset()

// Do some testing
stubs.Stub(&v1, 5)

// More testing
stubs.Stub(&b2, 6)
```

多次打桩的级联表达式代码如下：
```shell script
defer gostub.Stub(&v1, 1).Stub(&v2, 2).Reset()
```

使用GoConvey测试框架和GoStub测试框架编写的测试用例如下：
```
package main

import (
   "fmt"
   "testing"

   "GoExample/GoStub/StubFunction"

   "time"

   "github.com/prashantv/gostub"
   . "github.com/smartystreets/goconvey/convey"
)

var counter = 100
var CleanUp = cleanUp

func cleanUp(val string) {
   fmt.Println(val)
}

func TestFuncDemo(t *testing.T) {
   Convey("TestFuncDemo", t, func() {
      Convey("for succ", func() {
         stubs := gostub.Stub(&counter, 200)
         defer stubs.Reset()
         stubs.Stub(&Adapter.TimeNow, func() time.Time {
            return time.Date(2015, 6, 1, 0, 0, 0, 0, time.UTC)
         })
         stubs.StubFunc(&CleanUp)
         fmt.Println(counter)
         fmt.Println(Adapter.TimeNow().Day())
         CleanUp("Hello go")
      })
   })
}
```

## 6、不适用场景
GoStub框架可以解决很多场景的函数打桩问题，
但下列复杂场景除外：

A、被测函数中多次调用了数据库读操作函数接口，
并且数据库为key-value型。

B、被测函数中有一个循环，用于一个批量操作，
当某一次操作失败，则返回失败，并进行错误处理。

C、被测函数中多次调用了同一底层操作函数。