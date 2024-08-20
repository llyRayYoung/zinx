package main

import (
	"github.com/llyRayYoung/zinx/examples/zinx_server/s_router"
	"github.com/llyRayYoung/zinx/zconf"
	"github.com/llyRayYoung/zinx/znet"
)

func main() {
	// Set up as WebSocket before starting. (在启动之前设置为 websocket)
	zconf.GlobalObject.Mode = ""
	zconf.GlobalObject.LogFile = ""

	s := znet.NewServer()

	s.AddRouter(100, &s_router.PingRouter{})
	s.AddRouter(1, &s_router.HelloZinxRouter{})

	s.Serve()
}
