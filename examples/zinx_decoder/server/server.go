package main

import (
	"github.com/llyRayYoung/zinx/examples/zinx_decoder/router"
	"github.com/llyRayYoung/zinx/zdecoder"
	"github.com/llyRayYoung/zinx/ziface"
	"github.com/llyRayYoung/zinx/zlog"
	"github.com/llyRayYoung/zinx/znet"
)

func DoConnectionBegin(conn ziface.IConnection) {
	zlog.Ins().InfoF("DoConnectionBegin is Called ...")
}

func DoConnectionLost(conn ziface.IConnection) {
	zlog.Ins().InfoF("Conn is Lost")
}

func main() {
	//zlog.SetLogFile("./logs", "app.log")
	s := znet.NewServer()

	s.SetOnConnStart(DoConnectionBegin)
	s.SetOnConnStop(DoConnectionLost)

	// TLV protocol corresponding to business function
	// TLV协议对应业务功能
	s.AddRouter(0x00000001, &router.TLVBusinessRouter{})

	// Process HTLVCRC protocol data
	// 处理HTLVCRC协议数据
	s.SetDecoder(zdecoder.NewHTLVCRCDecoder())

	// TLV protocol corresponding to business function, because the funcode field in client.go is 0x10
	// TLV协议对应业务功能，因为client.go中模拟数据funcode字段为0x10
	s.AddRouter(0x10, &router.HtlvCrcBusinessRouter{})

	// TLV protocol corresponding to business function, because the funcode field in client.go is 0x13
	// TLV协议对应业务功能，因为client.go中模拟数据funcode字段为0x13
	s.AddRouter(0x13, &router.HtlvCrcBusinessRouter{})

	//开启服务
	s.Serve()
}
