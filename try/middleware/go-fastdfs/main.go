package main

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
)

func main() {
	var obj interface{}
	req := httplib.Post("http://192.168.0.110:8081/group1/upload")
	req.PostFile("file", "./111.txt") //注意不是全路径
	req.Param("output", "json")
	req.Param("scene", "")
	req.Param("path", "")
	req.ToJSON(&obj)
	fmt.Print(obj)
}
