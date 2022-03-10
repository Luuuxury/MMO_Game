package src

import "github.com/aceld/zinx/znet"

func main() {
	//创建服务器句柄
	s := znet.NewServer()

	//启动服务
	s.Serve()
}
