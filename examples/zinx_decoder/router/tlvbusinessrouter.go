package router

import (
	"github.com/llyRayYoung/zinx/zdecoder"
	"github.com/llyRayYoung/zinx/ziface"
	"github.com/llyRayYoung/zinx/zlog"
	"github.com/llyRayYoung/zinx/znet"
)

type TLVBusinessRouter struct {
	znet.BaseRouter
}

func (this *TLVBusinessRouter) Handle(request ziface.IRequest) {

	msgID := request.GetMessage().GetMsgID()
	zlog.Ins().DebugF("Call TLVRouter Handle %d %+v\n", msgID, request.GetMessage().GetData())

	resp := request.GetResponse()
	if resp == nil {
		return
	}

	tlvData := resp.(zdecoder.TLVDecoder)
	zlog.Ins().DebugF("do msgid=0x00000001 data business %+v\n", tlvData)
}
