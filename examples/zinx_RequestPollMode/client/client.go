package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/llyRayYoung/zinx/examples/zinx_client/c_router"
	"github.com/llyRayYoung/zinx/ziface"
	"github.com/llyRayYoung/zinx/zlog"
	"github.com/llyRayYoung/zinx/znet"
)

// Custom business logic of the client (客户端自定义业务)
func business(conn ziface.IConnection) {

	for i := 0; i < 100; i++ {
		err := conn.SendMsg(1, []byte("Ping...[FromClient]"))
		if err != nil {
			fmt.Println(err)
			zlog.Error(err)
			break
		}

		err = conn.SendMsg(2, []byte("Ping...[FromClient]"))
		if err != nil {
			fmt.Println(err)
			zlog.Error(err)
			break
		}

	}
}

// Function to execute when the connection is created (创建连接的时候执行)
func DoClientConnectedBegin(conn ziface.IConnection) {
	zlog.Debug("DoConnecionBegin is Called ... ")

	// Set two connection properties after the connection is created (设置两个链接属性，在连接创建之后)
	conn.SetProperty("Name", "刘丹冰Aceld")
	conn.SetProperty("Home", "https://yuque.com/aceld")

	go business(conn)
}

// Function to execute when the connection is lost (连接断开的时候执行)
func DoClientConnectedLost(conn ziface.IConnection) {
	// Get the Name and Home properties of the connection before it is destroyed
	// (在连接销毁之前，查询conn的Name，Home属性)
	if name, err := conn.GetProperty("Name"); err == nil {
		zlog.Debug("Conn Property Name = ", name)
	}

	if home, err := conn.GetProperty("Home"); err == nil {
		zlog.Debug("Conn Property Home = ", home)
	}

	zlog.Debug("DoClientConnectedLost is Called ... ")
}

func main() {
	// Create a client handle using Zinx's Method (创建一个Client句柄，使用Zinx的方法)
	client := znet.NewClient("127.0.0.1", 8999)

	// Set the business logic to execute when the connection is created or lost
	// (添加首次建立链接时的业务)
	client.SetOnConnStart(DoClientConnectedBegin)
	client.SetOnConnStop(DoClientConnectedLost)

	// Register routers for the messages received from the server
	// (注册收到服务器消息业务路由)
	client.AddRouter(2, &c_router.PingRouter{})
	client.AddRouter(3, &c_router.HelloRouter{})

	// Start the client
	client.Start()

	// close
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	fmt.Println("===exit===", sig)
	client.Stop()
	time.Sleep(time.Second * 2)
}
