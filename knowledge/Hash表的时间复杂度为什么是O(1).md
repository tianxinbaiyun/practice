@[TOC](Hash表的时间复杂度为什么是O(1)?)


今天在面试的时候说到HashMap，面试官问了这么一个问题：你说HashMap的get迭代了一个链表，那怎么保证HashMap的时间复杂度O(1)?链表的查找的时间复杂度又是多少？ 
在这之前我是阅读过HashMap的源码的：Java7源码浅析——对HashMap的理解 
由上一个博客可知我们对HashMap的查询，如源代码所示：

```shell script
public V get(Object key) {  
        if (key == null)  
            return getForNullKey();  
        int hash = hash(key.hashCode());  
        for (Entry<K,V> e = table[indexFor(hash, table.length)];  
             e != null;  
             e = e.next) {  
            Object k;  
            if (e.hash == hash && ((k = e.key) == key || key.equals(k)))  
                return e.value;  
        }  
        return null;  
    }  

```

分四步： 

1.判断key，根据key算出索引。 

2.根据索引获得索引位置所对应的键值对链表。 

3.遍历键值对链表，根据key找到对应的Entry键值对。 

4.拿到value。 

分析： 

以上四步要保证HashMap的时间复杂度O(1)，
需要保证每一步都是O(1)，
现在看起来就第三步对链表的循环的时间复杂度影响最大，
链表查找的时间复杂度为O(n)，与链表长度有关。
我们要保证那个链表长度为1，才可以说时间复杂度能满足O(1)。

但这么说来只有那个hash算法尽量减少冲突，才能使链表长度尽可能短，理想状态为1。
因此可以得出结论：HashMap的查找时间复杂度只有在最理想的情况下才会为O(1)，
而要保证这个理想状态不是我们开发者控制的。

————————————————

版权声明：本文为CSDN博主「GrayHJX」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/donggua3694857/article/details/64127131