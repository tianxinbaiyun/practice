@[TOC](电话号码的字母组合)

## 电话号码的字母组合

给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。


```

示例:

输入："23"
输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
说明:
尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。


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













