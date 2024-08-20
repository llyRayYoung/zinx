package main

import (
	"time"

	"github.com/llyRayYoung/zinx/ziface"
	"github.com/llyRayYoung/zinx/znet"
)

func main() {
	client := znet.NewClient("127.0.0.1", 8999)

	client.SetOnConnStart(func(connection ziface.IConnection) {
		_ = connection.SendMsg(1, []byte("hello zinx"))
	})

	client.Start()

	time.Sleep(time.Second)
}
