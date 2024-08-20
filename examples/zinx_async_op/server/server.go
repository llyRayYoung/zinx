package main

import (
	"github.com/llyRayYoung/zinx/examples/zinx_async_op/router"
	"github.com/llyRayYoung/zinx/ziface"
	"github.com/llyRayYoung/zinx/zlog"
	"github.com/llyRayYoung/zinx/znet"
)

func OnConnectionAdd(conn ziface.IConnection) {
	zlog.Debug("zinx_async_op OnConnectionAdd ===>")
}

func OnConnectionLost(conn ziface.IConnection) {
	zlog.Debug("zinx_async_op OnConnectionLost ===>")
}

func main() {
	s := znet.NewServer()

	s.SetOnConnStart(OnConnectionAdd)
	s.SetOnConnStop(OnConnectionLost)

	s.AddRouter(1, &router.LoginRouter{})

	s.Serve()
}
