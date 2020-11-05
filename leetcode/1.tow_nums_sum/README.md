@[TOC](两数之和)

## 两数之和

给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

## 思路 

### 方法一，暴力法

暴力法很简单，遍历每个元素 xx，并查找是否存在一个值与 target - xtarget−x 相等的目标元素。

复杂度分析：

时间复杂度：O(n^2)
对于每个元素，我们试图通过遍历数组的其余部分来寻找它所对应的目标元素，这将耗费 O(n)O(n) 的时间。因此时间复杂度为 O(n^2)O(n 2)。

空间复杂度：O(1)。

```text

func twoSum0(nums []int,target int) []int {
	if len(nums)<2{
		return nil
	}
	numsLen:= len(nums)
	for i:=0;i<numsLen ;i++  {
		for j:=i+1; j<numsLen;j++  {
			if nums[i]+nums[j]==target{
				return []int{i,j}
			}
		}
	}
	return nil
}

```

### 方法二：一遍哈希表

在进行迭代并将元素插入到表中的同时，我们还会回过头来检查表中是否已经存在当前元素所对应的目标元素。如果它存在，那我们已经找到了对应解，并立即将其返回。

复杂度分析：

时间复杂度：O(n)，
我们把包含有 nn 个元素的列表遍历两次。由于哈希表将查找时间缩短到 O(1) ，所以时间复杂度为 O(n)O(n)。

空间复杂度：O(n)O(n)，
所需的额外空间取决于哈希表中存储的元素数量，该表中存储了 n 个元素。


```textlang

func twoSum(nums []int, target int) []int {
	if len(nums)<2{
		return nil
	}
	var tmp map
	for i, i2 := range nums {
		if _,ok:=tmp[target-i2];ok{
			return []int{tmp[target-i2],i}
		}else {
			tmp[i2]=i
		}
	}
	return nil
}

```

## 细节注意
1.判断前对输入数值进行长度判断，程序的执行时间不变，但可以降低程序执行的内存消耗

2.对map类型数据，只对其进行内存分配（make），不对函数进行类型定义（var），会增加程序执行时间

3.对应map类型数据，对其进行类型定义（var），再对其进行内存分配（make），会比使用简写定义':='，所执行的内存消耗会小一点