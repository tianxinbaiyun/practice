@[TOC](两数相加)

## 两数相加

两数相加
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807


## 代码
```text

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l *ListNode = &ListNode{}
	pre := l
	flag := 0
	for l1 != nil || l2 != nil {
		pre.Next = &ListNode{}
		p := pre.Next
		x := 0
		y := 0
		if l1 != nil {
			x = l1.Val
		}
		if l2 != nil {
			y = l2.Val
		}
		p.Val = (x + y + flag) % 10
		flag = (x + y + flag) / 10
		pre = p
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if flag != 0 {
		pre.Next = &ListNode{Val: flag}
	}

	return l.Next
}

```
## 运行结果
```
&{3 0xc0000a6090}
&{3 0xc0000a60d0}
&{6 <nil>}

```
## 总结

### 算法
我们使用变量来跟踪进位，并从包含最低有效位的表头开始模拟逐位相加的过程。

就像你在纸上计算两个数字的和那样，我们首先从最低有效位也就是列表 l1l1 和 l2l2 的表头开始相加。由于每位数字都应当处于 0 \ldots 90…9 的范围内，我们计算两个数字的和时可能会出现 “溢出”。例如，5 + 7 = 125+7=12。在这种情况下，我们会将当前位的数值设置为 22，并将进位 carry = 1carry=1 带入下一次迭代。进位 carrycarry 必定是 00 或 11，这是因为两个数字相加（考虑到进位）可能出现的最大和为 9 + 9 + 1 = 199+9+1=19。

### 时间复杂度

时间复杂度：O(\max(m, n))O(max(m,n))，假设 mm 和 nn 分别表示 l1l1 和 l2l2 的长度，上面的算法最多重复 \max(m, n)max(m,n) 次。

空间复杂度：O(\max(m, n))O(max(m,n))， 新列表的长度最多为 \max(m,n) + 1max(m,n)+1。
