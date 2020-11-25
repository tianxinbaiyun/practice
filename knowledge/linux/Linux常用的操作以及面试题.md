@[TOC](Linux常用的操作以及面试题)

## 1.文件操作命令

0)列出当前目录下所有文件和目录？
```shell script
ls dir
```

1）查看当前目录下的详细列表信息,包括权限属性
```shell script
ls -l
```

2）查看当前目录下所有以cron开头命名的文件
```shell script
ll cron*
```


3)获取当前目录地址？
```shell script
pwd
```


4)退出到上一级目录？
```shell script
cd ..
```

5)回到上次所在的目录？
```shell script
cd -
```


6)进入用户默认的目录，进入用户家目录？
```shell script
cd ~
```
```shell script
cd
```

7)进入系统的根目录？
```shell script
cd /
```


8)创建一个目录？
```shell script
mkdir test
```

9)连续创建多层目录/a/b/c？
```shell script
mkdir -p /a/b/c #根目录下创建
```

```shell script
mkdir -p a/b/c   #当前目录下
```


10)如何创建一个文件？
```shell script
touch tt.txt
```


11)如何将一段内容写入文件？
```shell script
echo hello world>test.txt
```

如果当前目录下没有test.txt，创建该文件
```shell script
cat>test.txt #ctrl+d保存退出
```


12)重定向>和>>有什么区别

>覆盖旧内容

>>文件后追加

13)如何为文件重新命名？
```shell script
mv test.txt newtest.txt
```


14)移动tt.txt文件到上级目录？
```shell script
mv tt.txt ../
```


15)将当前目录下所有的文件全部一次性移到/root下
```shell script
mv * /root
```


16)如何拷贝tt.txt到/root目录下
```shell script
cp tt.txt /root
```


17）复制tomcat8080目录及该目录下所有的目录和文件到tomcat8081目录
```shell script
cp -r|R tomcat8080 tomcat8081
```


18)如何删除一个文件？
```shell script
rm test.txt #询问
```

```shell script
rm -f test.txt 不询问
```


19)如何一次性删除多级非空目录?(该目录下结构为a/b/c/test.txt)
```shell script
rm -rf a/b/c/test.txt
```


20)如何删除一个空目录？
```shell script
rmdir testdir
```


21)如何一次性删除多级空目录?a/b/c
```shell script
rmdir -p a/b/c
```


22)如何打包tar文件和解包?
```shell script
tar -zcvf mytar.tar test.txt
```

```shell script
tar -zxvf mytar.tar
```


23)如何压缩.zip文件和解压缩？
```shell script
zip myzip.zip test.txt
```

```shell script
unzip -o myzip.zip #不询问，直接覆盖原来文件
```


24)如何查看小文件内容？如何查看文件内容的同时显示行号？
```shell script
cat test.txt
```

```shell script
cat -n test.txt
```


25)如何只查看test.txt文件的前2行？
```shell script
head -2 test.txt
```


26)如何只查看test.txt文件的末尾2行？
```shell script
tail -2 test.txt
```


27)动态监控catalina.out文件的内容？
```shell script
tail -f Catalina.out
```


28)分页查看文件的命令是什么？
```shell script
more filename
```

```shell script
less filename
```


29)统计文本总行数?

```shell script
wc -l filename
```


30）编辑文本内容的命令是什么
```shell script
vi vim
```


31）vi的三种模式如何切换

命令模式  I-》输入模式   ：->底行模式

输入模式  esc->命令模式

底行模式   w保存 q退出  wq保存退出 q!强制退出

32)linux里如何查看帮助

```shell script
help ls
```
```shell script
ls --help
```
     
```shell script
man ls
```


33)如何在指定的文件中查找某个含有关键字的行？
```shell script
grep this filename
```


34)查找系统中所有txt后缀文件
```shell script
find -type f -name *.txt
```


-type:文件类型

-name:根据文件名称查找

35)查看磁盘剩余的磁盘空间？
```shell script
df
```
  

36)切换用户  
```shell script
su
```


————————————————


版权声明：本文为CSDN博主「mengchuan6666」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/mengchuan6666/article/details/85860802