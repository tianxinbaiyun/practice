@[TOC](四数之和)

## 四数之和

给定一个包含 n 个整数的数组 nums 和一个目标值 target，
判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。

注意：

答案中不可以包含重复的四元组。


```

示例:

给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。

满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]




```



## 代码
```text
package main

import (
	"fmt"
	"time"
)

func main()  {
	t:=time.Now()
	res :=letterCombinations("45678")
	fmt.Println(time.Since(t))
	fmt.Println(res)
}
func letterCombinations(digits string) []string {
	var (
		chars map[int32][]string
		res []string
	)

	chars = map[int32][]string{
		'2':[]string{"a","b","c"},
		'3':[]string{"d","e","f"},
		'4':[]string{"g","h","i"},
		'5':[]string{"j","k","l"},
		'6':[]string{"m","n","o"},
		'7':[]string{"p","q","r","s"},
		'8':[]string{"t","u","v"},
		'9':[]string{"w","x","y","z"},
	}
	for k, v := range digits {
		if _,ok:=chars[v];!ok{
			return []string{}
		}
		if k==0{
			res =  chars[v]
		}else{
			tmp  := make([]string,0)
			for _, vv := range res {
				for _, vvv := range chars[v] {
					tmp = append(tmp,fmt.Sprintf("%s%s",vv,vvv))
				}
			}
			res = tmp
		}
	}
	return res
}

```


## 运行结果

```
[[-99524 -394 99918] [-97960 -1958 99918] [-97460 -2445 99905] [-93651 -6167 99818] [-91003 -8798 99801] [-85560 -14104 99664] [-85530 -14104 99634] [-84981 -14606 99587] [-83966 -15619 99585] [-83940 -15619 99559] [-76090 -23447 99537] [-72973 -26521 99494] [-72207 -27101 99308] [-63466 -35731 99197]]

```


## 总结

### 算法

分治

于是，“将整数转换为罗马数字”的过程，就是用上面这张表中右边的数字作为“加法因子”去分解一个整数，目的是“分解的整数个数”尽可能少，因此，对于这道问题，类似于用最少的纸币凑成一个整数，贪心算法的规则如下：

每一步都使用当前较大的罗马数字作为加法因子，最后得到罗马数字表示就是长度最少的。


### 时间复杂度

时间复杂度： O(3^N×4^M) ，其中 N 是输入数字中对应 3 个字母的数目（比方说 2，3，4，5，6，8）， 
M 是输入数字中对应 4 个字母的数目（比方说 7，9），N+M 是输入数字的总数。

空间复杂度：O(3^N×4^M) ，这是因为需要保存3^N×4^M个结果。













