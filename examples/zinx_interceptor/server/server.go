package main

import (
	"github.com/llyRayYoung/zinx/examples/zinx_interceptor/interceptors"
	"github.com/llyRayYoung/zinx/examples/zinx_interceptor/router"
	"github.com/llyRayYoung/zinx/znet"
)

func main() {
	server := znet.NewServer()

	server.AddRouter(1, &router.HelloRouter{})

	// Add Custom Interceptor
	server.AddInterceptor(&interceptors.MyInterceptor{})

	server.Serve()
}
