package main

import (
	"context"
	"fmt"
	pd "github.com/tianxinbaiyun/practice/try/frame/go-micro/core/pb/helloworld"
	"google.golang.org/grpc"
	"net"
)

type server struct{}

//一个打招呼的函数

//rpc
//函数关键字（对象）函数名（客户端发送过来的内容 ， 返回给客户端的内容） 错误返回值

// SayHellogrpc
//函数关键字 （对象）函数名 （cotext，客户端发过来的参数 ）（发送给客户端的参数，错误）
func (c *server) SayHello(ctx context.Context, in *pd.HelloReq) (out *pd.HelloRsp, err error) {

	return &pd.HelloRsp{Msg: "hello" + in.Name}, nil
}

// SayName 一个说名字的服务
func (c *server) SayName(ctx context.Context, in *pd.NameReq) (out *pd.NameRsp, err error) {

	return &pd.NameRsp{Msg: in.Name + "早上好"}, nil
}

func main() {
	//创建网络
	ln, err := net.Listen("tcp", ":10086")
	if err != nil {
		fmt.Println("网络错误", err)
	}

	//创建grpc的服务
	srv := grpc.NewServer()

	//注册服务
	pd.RegisterHelloworldServiceServer(srv, &server{})

	//等待网络连接
	err = srv.Serve(ln)
	if err != nil {
		fmt.Println("网络错误", err)
	}

}
