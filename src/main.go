package src

import (
	"MM0_Game/src/core"
	"fmt"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
)

/*
我们先在Server的main入口，给链接绑定一个创建之后的hook方法，因为上线的时候是服务器自动回复客户端玩家ID和坐标，
那么需要我们在连接创建完毕之后，自动触发，正好我们可以利用Zinx框架的SetOnConnStart方法.
*/

// OnConnecionAdd 当客户端建立连接的时候的hook函数
func OnConnecionAdd(conn ziface.IConnection) {
	//创建一个玩家
	player := core.NewPlayer(conn)
	//同步当前的PlayerID给客户端， 走MsgID:1 消息
	player.SyncPid()
	//同步当前玩家的初始化坐标信息给客户端，走MsgID:200消息
	player.BroadCastStartPosition()

	fmt.Println("=====> Player pidId = ", player.Pid, " arrived ====")
}

func main() {
	//创建服务器句柄
	s := znet.NewServer()

	//注册客户端连接建立和丢失函数
	s.SetOnConnStart(OnConnecionAdd)

	//启动服务
	s.Serve()
}
