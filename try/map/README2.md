@[TOC](golang之map的使用声明)
## 1.map的基本介绍

map是key-value数据结构，又称为字段或者关联数组。
类似其它编程语言的集合，在编程中是经常使用到的

## map的操作
### 2.map的声明

#### 1）基本语法
```
var map 变量名 map[keytype] valuetype
```
　　　　　
注意：声明是不会分配内存的，初始化需要make，分配内存后才能赋值和使用。
```text
func main() {
    var a map[string]string
    a = make(map[string]string, 10)
    a["no1"] = "allin"
    a["no2"] = "alex"
    a["no1"] = "egon"
    a["no3"] = "sdfegon"
    fmt.Println(a)
}
```

对上面代码的说明

　　a.map在使用前一定要make

　　b.map的key是不能重复，如果重复了，则以最后这个key-value为准

　　c.map的value是可以相同的

　　d.map的key-value是无序

　　e.make内置函数数目

 

#### map的三种声明方式：
```text
    var a map[string]string
    a = make(map[string]string, 10)
    a["no1"] = "allin"
    a["no2"] = "alex"
    a["no1"] = "egon"
    a["no3"] = "sdfegon"
    fmt.Println(a)

    var b map[string]string
    b = make(map[string]string, 10)
    b["no1"] = "宋江"
    b["no2"] = "吴用"
    b["no1"] = "武松"
    b["no3"] = "吴用"
    fmt.Println(b)

    heros := map[string]string {
        "hero1": "宋江",
        "hero2": "lujunyi",
        "hero3": "吴用",
    }
    heros["hero4"] = "林冲"
    fmt.Println("heros=", heros)
```

### map删除：

delete(map, "key")，delete是一个内置函数，如果key存在，就删除该key-value,如果key不存在，不操作，但是也不会报错。

如果我们要删除map的所有key,没有一个专门的方法一次删除，可以遍历一下key，逐个删除或者map = kake(...),make一个新的，让原来的成为垃圾，被gc回收。

### map查找：

　　
```text
    val, ok := studentMap["stu02"]["name"]
    if ok {
        fmt.Println("aaa", val)
    } else {
        fmt.Println("bbb")
    }
```

说明：如果student这个mapk中存在“nol", 那么返回true, 否则返回false

### map遍历：

　　

for _, v := range studentMap {
        //fmt.Println(k, v)
        for k1, v1 := range v {
            fmt.Println(k1, v1)
        }
    }
### map切片：

　　切片的数据类型如果是map，则我们称为slice of map,map切片，这样使用则map个数就可以动态变化了。

### map排序：

　　golang中没有一个专门的方法针对map的key进行排序

　　golang中map的排序，是先将key进行排序，然后根据key值遍历输出即可
```text
    map1 := make(map[int]int, 10)
    map1[10] = 100
    map1[1] = 13
    map1[4] = 56
    map1[8] = 90
    fmt.Println(map1)

    var keys []int
    for k, _ := range map1 {
        keys = append(keys, k)
    }
    sort.Ints(keys)
    for _, k := range keys {
        fmt.Printf("map1[%v]=%v \n", k, map1[k])
    }
```

## map使用细节：

　　1）map费用类型，遵守引用类型传递的机制，在一个函数接收map，修改后，会直接修改原来的map

　　2）map的容量达到后，再想map增加元素，会自动扩容，并不会发生panic，也就是说map能动态的增长键值对

　　3）map的value也经常使用struct类型，更适合管理复杂的数据






————————————————

版权声明：本文为博客园博主「顽强的allin」的原创文章，
遵循 CC 4.0 BY-SA 版权协议，转载请附上原文出处链接及本声明。
原文链接：https://www.cnblogs.com/xiangxiaolin/p/11839162.html




