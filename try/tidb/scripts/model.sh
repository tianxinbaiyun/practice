#!/bin/bash

# 根目录
root=`pwd`

# 数据库配置
host=192.168.15.131
port=3306
dbname=study
username=root
passwd=123456
# 数据库配置

sxorm=${GOPATH}/bin/xorm
table=$1

if ! [[ -x ${sxorm} ]]; then
    echo "未安装xorm工具，请先安装!"
    exit 1
fi

if [[ ${table} == "" ]]; then
    echo "请输入需要导出的表名称(支持正则表达式)"
    exit 1
fi

# 数据库连接字符串
conn="${username}:${passwd}@tcp(${host}:${port})/${dbname}?charset=utf8&parseTime=True&loc=Local"

${sxorm} reverse "mysql" ${conn} ${root}/config ${root}/models ${table}

#xorm reverse "mysql" "root:123456@tcp(192.168.15.131:3306)/study?charset=utf8&parseTime=True&loc=Local" ./config ./models users