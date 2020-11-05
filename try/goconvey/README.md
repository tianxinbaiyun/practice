[@top](GoConvey)

##GoConvey

在软件开发中，产品代码的正确性通过测试代码来保证，
而测试代码的正确性谁来保证？
答案是毫无争议的，肯定是程序员自己。
这就要求测试代码必须足够简单且表达力强，让错误无处藏身。
我们要有一个好鼻子，能够嗅出测试的坏味道，
及时的进行测试重构，从而让测试代码易于维护。
笔者从大量的编码实践中感悟道：
虽然能写出好的产品代码的程序员很牛，
但能写出好的测试代码的程序员更牛，尤其对于TDD实践。

要写出好的测试代码，必须精通相关的测试框架。
对于Golang的程序员来说，至少需要掌握下面四个测试框架：
```shell script
GoConvey
GoStub
GoMock
Monkey
```

笔者将通过多篇文章来阐述这四个测试框架，
同时对于GoStub框架还将进行二次开发实践，
以便高效的解决较复杂场景的打桩问题。

本文将主要介绍GoConvey框架的基本使用方法，
从而指导读者更好的进行测试实践，最终写出简单优雅的测试代码。

## GoConvey简介
GoConvey是一款针对Golang的测试框架，
可以管理和运行测试用例，同时提供了丰富的断言函数，
并支持很多 Web 界面特性。

Golang虽然自带了单元测试功能，
并且在GoConvey框架诞生之前也出现了许多第三方测试框架，
但没有一个测试框架像GoConvey一样能够让程序员如此简洁优雅的编写测试代码。

## 安装
在命令行运行下面的命令：
```shell script
go get github.com/smartystreets/goconvey
```

运行时间较长，运行完后你会发现：

在$GOPATH/src目录下新增了github.com子目录，
该子目录里包含了GoConvey框架的库代码
在$GOPATH/bin目录下新增了GoConvey框架的可执行程序goconvey
基本使用方法
我们通过一个案例来介绍GoConvey框架的基本使用方法，并对要点进行归纳。

## 产品代码
我们实现一个判断两个字符串切片是否相等的函数StringSliceEqual，主要逻辑包括：

两个字符串切片长度不相等时，返回false

两个字符串切片一个是nil，另一个不是nil时，返回false

遍历两个切片，比较对应索引的两个切片元素值，如果不相等，返回false

否则，返回true

根据上面的逻辑，代码实现如下所示：

```text
func StringSliceEqual(a, b []string) bool {
    if len(a) != len(b) {
        return false
    }

    if (a == nil) != (b == nil) {
        return false
    }

    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}
```
对于逻辑“两个字符串切片一个是nil，另一个不是nil时，
返回false”的实现代码有点不好理解：
```text
if (a == nil) != (b == nil) {
    return false
}
```

我们实例化一下a和b，即[]string{}和[]string(nil)，
这时两个字符串切片的长度都是0，但肯定不相等。

## 测试代码
先写一个正常情况的测试用例，如下所示：
```text
import (
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

func TestStringSliceEqual(t *testing.T) {
    Convey("TestStringSliceEqual should return true when a != nil  && b != nil", t, func() {
        a := []string{"hello", "goconvey"}
        b := []string{"hello", "goconvey"}
        So(StringSliceEqual(a, b), ShouldBeTrue)
    })
}
```

由于GoConvey框架兼容Golang原生的单元测试，
所以可以使用go test -v来运行测试。
打开命令行，进入$GOPATH/src/infra/alg目录下，
运行go test -v，则测试用例的执行结果日下：
```shell script
=== RUN   TestStringSliceEqual

  TestStringSliceEqual should return true when a != nil  && b != nil ✔


1 total assertion

--- PASS: TestStringSliceEqual (0.00s)
PASS
ok      infra/alg       0.006s
```


**上面的测试用例代码有如下几个要点：**

import goconvey包时，前面加点号"."，以减少冗余的代码。
凡是在测试代码中看到Convey和So两个方法，
肯定是convey包的，不要在产品代码中定义相同的函数名

测试函数的名字必须以Test开头，
而且参数类型必须为*testing.T

每个测试用例必须使用Convey函数包裹起来，
它的第一个参数为string类型的测试描述，
第二个参数为测试函数的入参（类型为*testing.T），
第三个参数为不接收任何参数也不返回任何值的函数（习惯使用闭包）

Convey函数的第三个参数闭包的实现中通过So函数完成断言判断，
它的第一个参数为实际值，第二个参数为断言函数变量，
第三个参数或者没有（当第二个参数为类ShouldBeTrue形式的函数变量）或者有（当第二个函数为类ShouldEqual形式的函数变量）

### 我们故意将该测试用例改为不过：
```text
import (
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

func TestStringSliceEqual(t *testing.T) {
    Convey("TestStringSliceEqual should return true when a != nil  && b != nil", t, func() {
        a := []string{"hello", "goconvey"}
        b := []string{"hello", "goconvey"}
        So(StringSliceEqual(a, b), ShouldBeFalse)
    })
}
```

测试用例的执行结果日下：
```shell script
=== RUN   TestStringSliceEqual

  TestStringSliceEqual should return true when a != nil  && b != nil ✘


Failures:

  * /Users/zhangxiaolong/Desktop/D/go-workspace/src/infra/alg/slice_test.go 
  Line 45:
  Expected: false
  Actual:   true


1 total assertion

--- FAIL: TestStringSliceEqual (0.00s)
FAIL
exit status 1
FAIL    infra/alg       0.006s

```


我们再补充3个测试用例：
```text
import (
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

func TestStringSliceEqual(t *testing.T) {
    Convey("TestStringSliceEqual should return true when a != nil  && b != nil", t, func() {
        a := []string{"hello", "goconvey"}
        b := []string{"hello", "goconvey"}
        So(StringSliceEqual(a, b), ShouldBeTrue)
    })

    Convey("TestStringSliceEqual should return true when a ＝= nil  && b ＝= nil", t, func() {
        So(StringSliceEqual(nil, nil), ShouldBeTrue)
    })

    Convey("TestStringSliceEqual should return false when a ＝= nil  && b != nil", t, func() {
        a := []string(nil)
        b := []string{}
        So(StringSliceEqual(a, b), ShouldBeFalse)
    })

    Convey("TestStringSliceEqual should return false when a != nil  && b != nil", t, func() {
        a := []string{"hello", "world"}
        b := []string{"hello", "goconvey"}
        So(StringSliceEqual(a, b), ShouldBeFalse)
    })
}
```

从上面的测试代码可以看出，
每一个Convey语句对应一个测试用例，
那么一个函数的多个测试用例可以通过一个测试函数的多个Convey语句来呈现。

测试用例的执行结果如下：
```shell script
=== RUN   TestStringSliceEqual

  TestStringSliceEqual should return true when a != nil  && b != nil ✔


1 total assertion


  TestStringSliceEqual should return true when a ＝= nil  && b ＝= nil ✔


2 total assertions


  TestStringSliceEqual should return false when a ＝= nil  && b != nil ✔


3 total assertions


  TestStringSliceEqual should return false when a != nil  && b != nil ✔


4 total assertions

--- PASS: TestStringSliceEqual (0.00s)
PASS
ok      infra/alg       0.006s
```


## Convey语句的嵌套
Convey语句可以无限嵌套，以体现测试用例之间的关系。
需要注意的是，只有最外层的Convey需要传入*testing.T类型的变量t。
我们将前面的测试用例通过嵌套的方式写另一个版本：
```text
import (
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

func TestStringSliceEqual(t *testing.T) {
    Convey("TestStringSliceEqual", t, func() {
        Convey("should return true when a != nil  && b != nil", func() {
            a := []string{"hello", "goconvey"}
            b := []string{"hello", "goconvey"}
            So(StringSliceEqual(a, b), ShouldBeTrue)
        })

        Convey("should return true when a ＝= nil  && b ＝= nil", func() {
            So(StringSliceEqual(nil, nil), ShouldBeTrue)
        })

        Convey("should return false when a ＝= nil  && b != nil", func() {
            a := []string(nil)
            b := []string{}
            So(StringSliceEqual(a, b), ShouldBeFalse)
        })

        Convey("should return false when a != nil  && b != nil", func() {
            a := []string{"hello", "world"}
            b := []string{"hello", "goconvey"}
            So(StringSliceEqual(a, b), ShouldBeFalse)
        })
    })
}
```

测试用例的执行结果如下：
```shell script
=== RUN   TestStringSliceEqual

  TestStringSliceEqual 
    should return true when a != nil  && b != nil ✔
    should return true when a ＝= nil  && b ＝= nil ✔
    should return false when a ＝= nil  && b != nil ✔
    should return false when a != nil  && b != nil ✔


4 total assertions

--- PASS: TestStringSliceEqual (0.00s)
PASS
ok      infra/alg       0.006s
```


可见，Convey语句嵌套的测试日志和Convey语句不嵌套的测试日志的显示有差异，笔者更喜欢这种以测试函数为单位多个测试用例集中显示的形式。

## Web 界面
GoConvey不仅支持在命令行进行自动化编译测试，
而且还支持在 Web 界面进行自动化编译测试。
想要使用GoConvey的 Web 界面特性，
需要在测试文件所在目录下执行goconvey：

$GOPATH/bin/goconvey
这时弹出一个页面，如下图所示：

goconvey-web.png
在 Web 界面中:

可以设置界面主题
查看完整的测试结果
使用浏览器提醒等实用功能
自动检测代码变动并编译测试
半自动化书写测试用例
查看测试覆盖率
临时屏蔽某个包的编译测试
Skip
针对想忽略但又不想删掉或注释掉某些断言操作，
GoConvey提供了Convey/So的Skip方法：

SkipConvey函数表明相应的闭包函数将不被执行
SkipSo函数表明相应的断言将不被执行
当存在SkipConvey或SkipSo时，
测试日志中会显式打上"skipped"形式的标记：

当测试代码中存在SkipConvey时，
相应闭包函数中不管是否为SkipSo，
都将被忽略，测试日志中对应的符号仅为一个"⚠"
当测试代码Convey语句中存在SkipSo时，
测试日志中每个So对应一个"✔"或"✘"，
每个SkipSo对应一个"⚠"，按实际顺序排列
不管存在SkipConvey还是SkipSo时，
测试日志中都有字符串"{n} total assertions (one or more sections skipped)"，其中{n}表示测试中实际已运行的断言语句数
定制断言函数
我们先看一下So的函数原型：

func So(actual interface{}, assert assertion, expected ...interface{})
第二个参数为assertion，它的原型为：

type assertion func(actual interface{}, expected ...interface{}) string
当assertion的返回值为""时表示断言成功，否则表示失败，GoConvey框架中的相关代码为：

const (
    success                = ""
    needExactValues        = "This assertion requires exactly %d comparison values (you provided %d)."
    needNonEmptyCollection = "This assertion requires at least 1 comparison value (you provided 0)."
)
我们简单实现一个assertion函数：

func ShouldSummerBeComming(actual interface{}, expected ...interface{}) string {
    if actual == "summer" && expected[0] == "comming" {
        return ""
    } else {
        return "summer is not comming!"
    }
}
我们仍然在slice_test文件中写一个简单测试：

func TestSummer(t *testing.T) {
    Convey("TestSummer", t, func() {
        So("summer", ShouldSummerBeComming, "comming")
        So("winter", ShouldSummerBeComming, "comming")
    })
}
根据ShouldSummerBeComming的实现，
Convey语句中第一个So将断言成功，第二个So将断言失败。
我们运行测试，查看执行结果，符合期望：

=== RUN   TestSummer

  TestSummer ✔✘


Failures:

  * /Users/zhangxiaolong/Desktop/D/go-workspace/src/infra/alg/slice_test.go 
  Line 52:
  summer is not comming!


2 total assertions

--- FAIL: TestSummer (0.00s)
FAIL
exit status 1
FAIL    infra/alg       0.006s
小结
Golang虽然自带了单元测试功能，但笔者建议大家使用已经成熟的第三方测试框架。本文主要介绍了GoConvey框架，通过文字结合代码示例讲解基本的使用方法，要点归纳如下：

import goconvey包时，前面加点号"."，以减少冗余的代码
测试函数的名字必须以Test开头，而且参数类型必须为*testing.T
每个测试用例必须使用Convey函数包裹起来，推荐使用Convey语句的嵌套，即一个函数有一个测试函数，测试函数中嵌套两级Convey语句，第一级Convey语句对应测试函数，第二级Convey语句对应测试用例
Convey语句的第三个参数习惯以闭包的形式实现，在闭包中通过So语句完成断言
使用GoConvey框架的 Web 界面特性，作为命令行的补充
在适当的场景下使用SkipConvey函数或SkipSo函数
当测试中有需要时，可以定制断言函数
至此，希望读者已经掌握了GoConvey框架的基本用法，
从而可以写出简单优雅的测试代码。

然而，事情并没有这么简单！
试想，如果在被测函数中调用了底层os包的stat函数，
我们该如何写测试代码？
其实答案也并不复杂，我们将在下一篇文章中揭晓。

作者：_张晓龙_
链接：https://www.jianshu.com/p/e3b2b1194830
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。