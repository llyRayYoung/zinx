package router

import (
	"bytes"
	"encoding/hex"
	"github.com/llyRayYoung/zinx/zdecoder"
	"github.com/llyRayYoung/zinx/ziface"
	"github.com/llyRayYoung/zinx/zlog"
	"github.com/llyRayYoung/zinx/znet"
)

type Data0x10Router struct {
	znet.BaseRouter
}

func (this *Data0x10Router) Handle(request ziface.IRequest) {
	zlog.Ins().DebugF("Data0x10Router Handle %s \n", hex.EncodeToString(request.GetMessage().GetData()))
	_response := request.GetResponse()
	if _response != nil {
		switch _response.(type) {
		case zdecoder.HtlvCrcDecoder:
			_data := _response.(zdecoder.HtlvCrcDecoder)
			//zlog.Ins().DebugF("Data0x10Router %v \n", _data)
			buffer := pack10(_data)
			request.GetConnection().Send(buffer)
		}
	}
}

// 头码   功能码 数据长度      Body                         CRC
// A2      10     0E        0102030405060708091011121314 050B
// HeadCode FuncCode DataLen   Body                         CRC
// A2       10       0E        0102030405060708091011121314 050B
func pack10(_data zdecoder.HtlvCrcDecoder) []byte {
	buffer := bytes.NewBuffer([]byte{})
	buffer.WriteByte(0xA1)
	buffer.WriteByte(_data.Funcode)
	buffer.WriteByte(0x1E)
	//3~9:唯一设备码	将IMEI码转换为16进制
	//3~9: Unique device code: Convert the IMEI code to hexadecimal
	buffer.Write(_data.Body[:7])
	//10~14：园区代码	后台根据幼儿园生成的唯一代码
	//10~14: Park code: A unique code generated by the background according to the kindergarten
	buffer.Write([]byte{10, 11, 12, 13, 14})
	//15~18：时间戳	实际当前北京时间的时间戳，转换为16进制
	//15~18: Timestamp: The actual current Beijing time stamp, converted to hexadecimal
	buffer.Write([]byte{15, 16, 17, 18})
	//19：RFID模块工作模式	0x01-离线工作模式（默认工作模式）0x02-在线工作模式
	//19: RFID module working mode: 0x01-offline working mode (default working mode) 0x02-online working mode
	buffer.WriteByte(0x02)
	//20~27：通讯密匙	预留，全填0x00
	//20~27: Communication key: Reserved, fill with 0x00
	buffer.Write([]byte{20, 21, 22, 23, 24, 25, 26, 27})
	//28：出水方式	0x00-放杯出水，取杯停止出水 0x01-刷一下卡出水，再刷停止出水【数联默认】
	//28: Water outlet mode: 0x00-put cup and stop water, 0x01-swipe card for water, swipe card again to stop water [Suliandu default]
	buffer.WriteByte(0x01)
	//29~32：预留	全填0x00
	//29~32: Reserved: Fill with 0x00
	buffer.Write([]byte{29, 30, 31, 32})
	crc := zdecoder.GetCrC(buffer.Bytes())
	buffer.Write(crc)
	return buffer.Bytes()

}
