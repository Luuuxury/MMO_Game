package core

import (
	"fmt"
	"github.com/aceld/zinx/ziface"
	"math/rand"
	"sync"
)

// Player 玩家对象
type Player struct {
	Pid  int32              //玩家ID
	Conn ziface.IConnection //当前玩家的连接
	X    float32            //平面x坐标
	Y    float32            //高度
	Z    float32            //平面y坐标 (注意不是Y)
	V    float32            //旋转0-360度
}

/*
	Player ID 生成器
*/
// 全局ID 计数器
var PidGen int32 = 1  //用来生成玩家ID的计数器
var IdLock sync.Mutex //保护PidGen的互斥机制

// NewPlayer 创建一个玩家对象
func NewPlayer(conn ziface.IConnection) *Player {
	//生成一个PID
	IdLock.Lock()
	id := PidGen
	PidGen++
	IdLock.Unlock()

	p := &Player{
		Pid:  id,
		Conn: conn,
		X:    float32(160 + rand.Intn(10)), //随机在160坐标点 基于X轴偏移若干坐标
		Y:    0,                            //高度为0
		Z:    float32(134 + rand.Intn(17)), //随机在134坐标点 基于Y轴偏移若干坐标
		V:    0,                            //角度为0，尚未实现
	}

	return p
}

/*
	发送消息给客户端，
	主要是将pb的protobuf数据序列化之后发送
	Plyaer类中有当前玩家的ID，和当前玩家与客户端绑定的conn，还有就是地图的坐标信,NewPlayer()提供初始化玩家方法。
	由于Player经常需要和客户端发送消息，那么我们可以给Player提供一个SendMsg()方法，供客户端发送消息
*/
func (p *Player) SendMsg(msgId uint32, data proto.Message) {
	fmt.Printf("before Marshal data = %+v\n", data)
	//将proto Message结构体序列化
	msg, err := proto.Marshal(data)
	if err != nil {
		fmt.Println("marshal msg err: ", err)
		return
	}
	fmt.Printf("after Marshal data = %+v\n", msg)

	if p.Conn == nil {
		fmt.Println("connection in player is nil")
		return
	}

	//调用Zinx框架的SendMsg发包
	if err := p.Conn.SendMsg(msgId, msg); err != nil {
		fmt.Println("Player SendMsg error !")
		return
	}

	return
}
